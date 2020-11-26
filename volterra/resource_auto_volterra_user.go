//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//

package volterra

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gopkg.volterra.us/stdlib/client/vesapi"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_user "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/user"
)

// resourceVolterraUser is implementation of Volterra's User resources
func resourceVolterraUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraUserCreate,
		Read:   resourceVolterraUserRead,
		Update: resourceVolterraUserUpdate,
		Delete: resourceVolterraUserDelete,

		Schema: map[string]*schema.Schema{

			"annotations": {
				Type:     schema.TypeMap,
				Optional: true,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"disable": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},

			"contacts": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kind": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"namespace": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"tenant": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"email": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"first_name": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"last_name": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"locale": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

// resourceVolterraUserCreate creates User resource
func resourceVolterraUserCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_user.CreateSpecType{}
	createReq := &ves_io_schema_user.CreateRequest{
		Metadata: createMeta,
		Spec:     createSpec,
	}

	if v, ok := d.GetOk("annotations"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		createMeta.Annotations = ms
	}

	if v, ok := d.GetOk("description"); ok && !isIntfNil(v) {
		createMeta.Description =
			v.(string)
	}

	if v, ok := d.GetOk("disable"); ok && !isIntfNil(v) {
		createMeta.Disable =
			v.(bool)
	}

	if v, ok := d.GetOk("labels"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		createMeta.Labels = ms
	}

	if v, ok := d.GetOk("name"); ok && !isIntfNil(v) {
		createMeta.Name =
			v.(string)
	}

	if v, ok := d.GetOk("namespace"); ok && !isIntfNil(v) {
		createMeta.Namespace =
			v.(string)
	}

	if v, ok := d.GetOk("contacts"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		contactsInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		createSpec.Contacts = contactsInt
		for i, ps := range sl {

			cMapToStrVal := ps.(map[string]interface{})
			contactsInt[i] = &ves_io_schema.ObjectRefType{}

			contactsInt[i].Kind = "contact"

			if v, ok := cMapToStrVal["name"]; ok && !isIntfNil(v) {
				contactsInt[i].Name = v.(string)
			}

			if v, ok := cMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				contactsInt[i].Namespace = v.(string)
			}

			if v, ok := cMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				contactsInt[i].Tenant = v.(string)
			}

			if v, ok := cMapToStrVal["uid"]; ok && !isIntfNil(v) {
				contactsInt[i].Uid = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("email"); ok && !isIntfNil(v) {

		createSpec.Email =
			v.(string)
	}

	if v, ok := d.GetOk("first_name"); ok && !isIntfNil(v) {

		createSpec.FirstName =
			v.(string)
	}

	if v, ok := d.GetOk("last_name"); ok && !isIntfNil(v) {

		createSpec.LastName =
			v.(string)
	}

	if v, ok := d.GetOk("locale"); ok && !isIntfNil(v) {

		createSpec.Locale =
			v.(string)
	}

	if v, ok := d.GetOk("type"); ok && !isIntfNil(v) {

		createSpec.Type = ves_io_schema_user.UserType(ves_io_schema_user.UserType_value[v.(string)])

	}

	log.Printf("[DEBUG] Creating Volterra User object with struct: %+v", createReq)

	createUserResp, err := client.CreateObject(context.Background(), ves_io_schema_user.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating User: %s", err)
	}
	d.SetId(createUserResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraUserRead(d, meta)
}

func resourceVolterraUserRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_user.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] User %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra User %q: %s", d.Id(), err)
	}
	return setUserFields(client, d, resp)
}

func setUserFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraUserUpdate updates User resource
func resourceVolterraUserUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_user.ReplaceSpecType{}
	updateReq := &ves_io_schema_user.ReplaceRequest{
		Metadata: updateMeta,
		Spec:     updateSpec,
	}
	if v, ok := d.GetOk("annotations"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		updateMeta.Annotations = ms
	}

	if v, ok := d.GetOk("description"); ok && !isIntfNil(v) {
		updateMeta.Description =
			v.(string)
	}

	if v, ok := d.GetOk("disable"); ok && !isIntfNil(v) {
		updateMeta.Disable =
			v.(bool)
	}

	if v, ok := d.GetOk("labels"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		updateMeta.Labels = ms
	}

	if v, ok := d.GetOk("name"); ok && !isIntfNil(v) {
		updateMeta.Name =
			v.(string)
	}

	if v, ok := d.GetOk("namespace"); ok && !isIntfNil(v) {
		updateMeta.Namespace =
			v.(string)
	}

	if v, ok := d.GetOk("contacts"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		contactsInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		updateSpec.Contacts = contactsInt
		for i, ps := range sl {

			cMapToStrVal := ps.(map[string]interface{})
			contactsInt[i] = &ves_io_schema.ObjectRefType{}

			contactsInt[i].Kind = "contact"

			if v, ok := cMapToStrVal["name"]; ok && !isIntfNil(v) {
				contactsInt[i].Name = v.(string)
			}

			if v, ok := cMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				contactsInt[i].Namespace = v.(string)
			}

			if v, ok := cMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				contactsInt[i].Tenant = v.(string)
			}

			if v, ok := cMapToStrVal["uid"]; ok && !isIntfNil(v) {
				contactsInt[i].Uid = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("first_name"); ok && !isIntfNil(v) {

		updateSpec.FirstName =
			v.(string)
	}

	if v, ok := d.GetOk("last_name"); ok && !isIntfNil(v) {

		updateSpec.LastName =
			v.(string)
	}

	if v, ok := d.GetOk("locale"); ok && !isIntfNil(v) {

		updateSpec.Locale =
			v.(string)
	}

	log.Printf("[DEBUG] Updating Volterra User obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_user.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating User: %s", err)
	}

	return resourceVolterraUserRead(d, meta)
}

func resourceVolterraUserDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_user.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] User %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra User before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra User obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_user.ObjectType, namespace, name)
}
