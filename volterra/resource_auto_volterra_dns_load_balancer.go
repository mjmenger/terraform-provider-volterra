//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-tf-provider. DO NOT EDIT.
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
	ves_io_schema_dns_load_balancer "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/dns_load_balancer"
	ves_io_schema_views "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
)

// resourceVolterraDnsLoadBalancer is implementation of Volterra's DnsLoadBalancer resources
func resourceVolterraDnsLoadBalancer() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraDnsLoadBalancerCreate,
		Read:   resourceVolterraDnsLoadBalancerRead,
		Update: resourceVolterraDnsLoadBalancerUpdate,
		Delete: resourceVolterraDnsLoadBalancerDelete,

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
				ForceNew: true,
			},

			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"record_type": {
				Type:     schema.TypeString,
				Required: true,
			},

			"response_cache": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"default_response_cache_parameters": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"disable": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"response_cache_parameters": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"cache_cidr_ipv4": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"cache_cidr_ipv6": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"cache_ttl": {
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"rule_list": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"rules": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"nxdomain": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"pool": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

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

									"geo_location_label_selector": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"expressions": {

													Type: schema.TypeList,

													Required: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},

									"geo_location_set": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

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

									"score": {
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

// resourceVolterraDnsLoadBalancerCreate creates DnsLoadBalancer resource
func resourceVolterraDnsLoadBalancerCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_dns_load_balancer.CreateSpecType{}
	createReq := &ves_io_schema_dns_load_balancer.CreateRequest{
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

	//record_type
	if v, ok := d.GetOk("record_type"); ok && !isIntfNil(v) {

		createSpec.RecordType = ves_io_schema_dns_load_balancer.ResourceRecordType(ves_io_schema_dns_load_balancer.ResourceRecordType_value[v.(string)])

	}

	//response_cache
	if v, ok := d.GetOk("response_cache"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		responseCache := &ves_io_schema_dns_load_balancer.ResponseCache{}
		createSpec.ResponseCache = responseCache
		for _, set := range sl {
			responseCacheMapStrToI := set.(map[string]interface{})

			responseCacheParametersChoiceTypeFound := false

			if v, ok := responseCacheMapStrToI["default_response_cache_parameters"]; ok && !isIntfNil(v) && !responseCacheParametersChoiceTypeFound {

				responseCacheParametersChoiceTypeFound = true

				if v.(bool) {
					responseCacheParametersChoiceInt := &ves_io_schema_dns_load_balancer.ResponseCache_DefaultResponseCacheParameters{}
					responseCacheParametersChoiceInt.DefaultResponseCacheParameters = &ves_io_schema.Empty{}
					responseCache.ResponseCacheParametersChoice = responseCacheParametersChoiceInt
				}

			}

			if v, ok := responseCacheMapStrToI["disable"]; ok && !isIntfNil(v) && !responseCacheParametersChoiceTypeFound {

				responseCacheParametersChoiceTypeFound = true

				if v.(bool) {
					responseCacheParametersChoiceInt := &ves_io_schema_dns_load_balancer.ResponseCache_Disable{}
					responseCacheParametersChoiceInt.Disable = &ves_io_schema.Empty{}
					responseCache.ResponseCacheParametersChoice = responseCacheParametersChoiceInt
				}

			}

			if v, ok := responseCacheMapStrToI["response_cache_parameters"]; ok && !isIntfNil(v) && !responseCacheParametersChoiceTypeFound {

				responseCacheParametersChoiceTypeFound = true
				responseCacheParametersChoiceInt := &ves_io_schema_dns_load_balancer.ResponseCache_ResponseCacheParameters{}
				responseCacheParametersChoiceInt.ResponseCacheParameters = &ves_io_schema_dns_load_balancer.ResponseCacheParameters{}
				responseCache.ResponseCacheParametersChoice = responseCacheParametersChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["cache_cidr_ipv4"]; ok && !isIntfNil(v) {

						responseCacheParametersChoiceInt.ResponseCacheParameters.CacheCidrIpv4 = uint32(v.(int))

					}

					if v, ok := cs["cache_cidr_ipv6"]; ok && !isIntfNil(v) {

						responseCacheParametersChoiceInt.ResponseCacheParameters.CacheCidrIpv6 = uint32(v.(int))

					}

					if v, ok := cs["cache_ttl"]; ok && !isIntfNil(v) {

						responseCacheParametersChoiceInt.ResponseCacheParameters.CacheTtl = uint32(v.(int))

					}

				}

			}

		}

	}

	//rule_list
	if v, ok := d.GetOk("rule_list"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		ruleList := &ves_io_schema_dns_load_balancer.LoadBalancingRuleList{}
		createSpec.RuleList = ruleList
		for _, set := range sl {
			ruleListMapStrToI := set.(map[string]interface{})

			if v, ok := ruleListMapStrToI["rules"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				rules := make([]*ves_io_schema_dns_load_balancer.LoadBalancingRule, len(sl))
				ruleList.Rules = rules
				for i, set := range sl {
					rules[i] = &ves_io_schema_dns_load_balancer.LoadBalancingRule{}
					rulesMapStrToI := set.(map[string]interface{})

					actionChoiceTypeFound := false

					if v, ok := rulesMapStrToI["nxdomain"]; ok && !isIntfNil(v) && !actionChoiceTypeFound {

						actionChoiceTypeFound = true

						if v.(bool) {
							actionChoiceInt := &ves_io_schema_dns_load_balancer.LoadBalancingRule_Nxdomain{}
							actionChoiceInt.Nxdomain = &ves_io_schema.Empty{}
							rules[i].ActionChoice = actionChoiceInt
						}

					}

					if v, ok := rulesMapStrToI["pool"]; ok && !isIntfNil(v) && !actionChoiceTypeFound {

						actionChoiceTypeFound = true
						actionChoiceInt := &ves_io_schema_dns_load_balancer.LoadBalancingRule_Pool{}
						actionChoiceInt.Pool = &ves_io_schema_views.ObjectRefType{}
						rules[i].ActionChoice = actionChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["name"]; ok && !isIntfNil(v) {

								actionChoiceInt.Pool.Name = v.(string)

							}

							if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

								actionChoiceInt.Pool.Namespace = v.(string)

							}

							if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

								actionChoiceInt.Pool.Tenant = v.(string)

							}

						}

					}

					geoLocationChoiceTypeFound := false

					if v, ok := rulesMapStrToI["geo_location_label_selector"]; ok && !isIntfNil(v) && !geoLocationChoiceTypeFound {

						geoLocationChoiceTypeFound = true
						geoLocationChoiceInt := &ves_io_schema_dns_load_balancer.LoadBalancingRule_GeoLocationLabelSelector{}
						geoLocationChoiceInt.GeoLocationLabelSelector = &ves_io_schema.LabelSelectorType{}
						rules[i].GeoLocationChoice = geoLocationChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["expressions"]; ok && !isIntfNil(v) {

								ls := make([]string, len(v.([]interface{})))
								for i, v := range v.([]interface{}) {
									ls[i] = v.(string)
								}
								geoLocationChoiceInt.GeoLocationLabelSelector.Expressions = ls

							}

						}

					}

					if v, ok := rulesMapStrToI["geo_location_set"]; ok && !isIntfNil(v) && !geoLocationChoiceTypeFound {

						geoLocationChoiceTypeFound = true
						geoLocationChoiceInt := &ves_io_schema_dns_load_balancer.LoadBalancingRule_GeoLocationSet{}
						geoLocationChoiceInt.GeoLocationSet = &ves_io_schema_views.ObjectRefType{}
						rules[i].GeoLocationChoice = geoLocationChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["name"]; ok && !isIntfNil(v) {

								geoLocationChoiceInt.GeoLocationSet.Name = v.(string)

							}

							if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

								geoLocationChoiceInt.GeoLocationSet.Namespace = v.(string)

							}

							if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

								geoLocationChoiceInt.GeoLocationSet.Tenant = v.(string)

							}

						}

					}

					if w, ok := rulesMapStrToI["score"]; ok && !isIntfNil(w) {
						rules[i].Score = uint32(w.(int))
					}

				}

			}

		}

	}

	log.Printf("[DEBUG] Creating Volterra DnsLoadBalancer object with struct: %+v", createReq)

	createDnsLoadBalancerResp, err := client.CreateObject(context.Background(), ves_io_schema_dns_load_balancer.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating DnsLoadBalancer: %s", err)
	}
	d.SetId(createDnsLoadBalancerResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraDnsLoadBalancerRead(d, meta)
}

func resourceVolterraDnsLoadBalancerRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_dns_load_balancer.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] DnsLoadBalancer %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra DnsLoadBalancer %q: %s", d.Id(), err)
	}
	return setDnsLoadBalancerFields(client, d, resp)
}

func setDnsLoadBalancerFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraDnsLoadBalancerUpdate updates DnsLoadBalancer resource
func resourceVolterraDnsLoadBalancerUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_dns_load_balancer.ReplaceSpecType{}
	updateReq := &ves_io_schema_dns_load_balancer.ReplaceRequest{
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

	if v, ok := d.GetOk("record_type"); ok && !isIntfNil(v) {

		updateSpec.RecordType = ves_io_schema_dns_load_balancer.ResourceRecordType(ves_io_schema_dns_load_balancer.ResourceRecordType_value[v.(string)])

	}

	if v, ok := d.GetOk("response_cache"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		responseCache := &ves_io_schema_dns_load_balancer.ResponseCache{}
		updateSpec.ResponseCache = responseCache
		for _, set := range sl {
			responseCacheMapStrToI := set.(map[string]interface{})

			responseCacheParametersChoiceTypeFound := false

			if v, ok := responseCacheMapStrToI["default_response_cache_parameters"]; ok && !isIntfNil(v) && !responseCacheParametersChoiceTypeFound {

				responseCacheParametersChoiceTypeFound = true

				if v.(bool) {
					responseCacheParametersChoiceInt := &ves_io_schema_dns_load_balancer.ResponseCache_DefaultResponseCacheParameters{}
					responseCacheParametersChoiceInt.DefaultResponseCacheParameters = &ves_io_schema.Empty{}
					responseCache.ResponseCacheParametersChoice = responseCacheParametersChoiceInt
				}

			}

			if v, ok := responseCacheMapStrToI["disable"]; ok && !isIntfNil(v) && !responseCacheParametersChoiceTypeFound {

				responseCacheParametersChoiceTypeFound = true

				if v.(bool) {
					responseCacheParametersChoiceInt := &ves_io_schema_dns_load_balancer.ResponseCache_Disable{}
					responseCacheParametersChoiceInt.Disable = &ves_io_schema.Empty{}
					responseCache.ResponseCacheParametersChoice = responseCacheParametersChoiceInt
				}

			}

			if v, ok := responseCacheMapStrToI["response_cache_parameters"]; ok && !isIntfNil(v) && !responseCacheParametersChoiceTypeFound {

				responseCacheParametersChoiceTypeFound = true
				responseCacheParametersChoiceInt := &ves_io_schema_dns_load_balancer.ResponseCache_ResponseCacheParameters{}
				responseCacheParametersChoiceInt.ResponseCacheParameters = &ves_io_schema_dns_load_balancer.ResponseCacheParameters{}
				responseCache.ResponseCacheParametersChoice = responseCacheParametersChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["cache_cidr_ipv4"]; ok && !isIntfNil(v) {

						responseCacheParametersChoiceInt.ResponseCacheParameters.CacheCidrIpv4 = uint32(v.(int))

					}

					if v, ok := cs["cache_cidr_ipv6"]; ok && !isIntfNil(v) {

						responseCacheParametersChoiceInt.ResponseCacheParameters.CacheCidrIpv6 = uint32(v.(int))

					}

					if v, ok := cs["cache_ttl"]; ok && !isIntfNil(v) {

						responseCacheParametersChoiceInt.ResponseCacheParameters.CacheTtl = uint32(v.(int))

					}

				}

			}

		}

	}

	if v, ok := d.GetOk("rule_list"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		ruleList := &ves_io_schema_dns_load_balancer.LoadBalancingRuleList{}
		updateSpec.RuleList = ruleList
		for _, set := range sl {
			ruleListMapStrToI := set.(map[string]interface{})

			if v, ok := ruleListMapStrToI["rules"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				rules := make([]*ves_io_schema_dns_load_balancer.LoadBalancingRule, len(sl))
				ruleList.Rules = rules
				for i, set := range sl {
					rules[i] = &ves_io_schema_dns_load_balancer.LoadBalancingRule{}
					rulesMapStrToI := set.(map[string]interface{})

					actionChoiceTypeFound := false

					if v, ok := rulesMapStrToI["nxdomain"]; ok && !isIntfNil(v) && !actionChoiceTypeFound {

						actionChoiceTypeFound = true

						if v.(bool) {
							actionChoiceInt := &ves_io_schema_dns_load_balancer.LoadBalancingRule_Nxdomain{}
							actionChoiceInt.Nxdomain = &ves_io_schema.Empty{}
							rules[i].ActionChoice = actionChoiceInt
						}

					}

					if v, ok := rulesMapStrToI["pool"]; ok && !isIntfNil(v) && !actionChoiceTypeFound {

						actionChoiceTypeFound = true
						actionChoiceInt := &ves_io_schema_dns_load_balancer.LoadBalancingRule_Pool{}
						actionChoiceInt.Pool = &ves_io_schema_views.ObjectRefType{}
						rules[i].ActionChoice = actionChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["name"]; ok && !isIntfNil(v) {

								actionChoiceInt.Pool.Name = v.(string)

							}

							if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

								actionChoiceInt.Pool.Namespace = v.(string)

							}

							if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

								actionChoiceInt.Pool.Tenant = v.(string)

							}

						}

					}

					geoLocationChoiceTypeFound := false

					if v, ok := rulesMapStrToI["geo_location_label_selector"]; ok && !isIntfNil(v) && !geoLocationChoiceTypeFound {

						geoLocationChoiceTypeFound = true
						geoLocationChoiceInt := &ves_io_schema_dns_load_balancer.LoadBalancingRule_GeoLocationLabelSelector{}
						geoLocationChoiceInt.GeoLocationLabelSelector = &ves_io_schema.LabelSelectorType{}
						rules[i].GeoLocationChoice = geoLocationChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["expressions"]; ok && !isIntfNil(v) {

								ls := make([]string, len(v.([]interface{})))
								for i, v := range v.([]interface{}) {
									ls[i] = v.(string)
								}
								geoLocationChoiceInt.GeoLocationLabelSelector.Expressions = ls

							}

						}

					}

					if v, ok := rulesMapStrToI["geo_location_set"]; ok && !isIntfNil(v) && !geoLocationChoiceTypeFound {

						geoLocationChoiceTypeFound = true
						geoLocationChoiceInt := &ves_io_schema_dns_load_balancer.LoadBalancingRule_GeoLocationSet{}
						geoLocationChoiceInt.GeoLocationSet = &ves_io_schema_views.ObjectRefType{}
						rules[i].GeoLocationChoice = geoLocationChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["name"]; ok && !isIntfNil(v) {

								geoLocationChoiceInt.GeoLocationSet.Name = v.(string)

							}

							if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

								geoLocationChoiceInt.GeoLocationSet.Namespace = v.(string)

							}

							if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

								geoLocationChoiceInt.GeoLocationSet.Tenant = v.(string)

							}

						}

					}

					if w, ok := rulesMapStrToI["score"]; ok && !isIntfNil(w) {
						rules[i].Score = uint32(w.(int))
					}

				}

			}

		}

	}

	log.Printf("[DEBUG] Updating Volterra DnsLoadBalancer obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_dns_load_balancer.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating DnsLoadBalancer: %s", err)
	}

	return resourceVolterraDnsLoadBalancerRead(d, meta)
}

func resourceVolterraDnsLoadBalancerDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_dns_load_balancer.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] DnsLoadBalancer %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra DnsLoadBalancer before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra DnsLoadBalancer obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_dns_load_balancer.ObjectType, namespace, name)
}
