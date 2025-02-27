---

page_title: "Volterra: tcp_loadbalancer"

description: "The tcp_loadbalancer allows CRUD of Tcp Loadbalancer resource on Volterra SaaS"
---------------------------------------------------------------------------------------------

Resource volterra_tcp_loadbalancer
==================================

The Tcp Loadbalancer allows CRUD of Tcp Loadbalancer resource on Volterra SaaS

~> **Note:** Please refer to [Tcp Loadbalancer API docs](https://volterra.io/docs/api/views-tcp-loadbalancer) to learn more

Example Usage
-------------

```hcl
resource "volterra_tcp_loadbalancer" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "advertise_on_public advertise_custom do_not_advertise advertise_on_public_default_vip" must be set

  advertise_custom {
    advertise_where {
      // One of the arguments from this list "site virtual_site vk8s_service virtual_network" must be set

      site {
        ip      = "8.8.8.8"
        network = "network"

        site {
          name      = "test1"
          namespace = "staging"
          tenant    = "acmecorp"
        }
      }

      // One of the arguments from this list "use_default_port port" must be set
      use_default_port = true
    }
  }
  // One of the arguments from this list "retract_cluster do_not_retract_cluster" must be set
  retract_cluster = true
  // One of the arguments from this list "hash_policy_choice_round_robin hash_policy_choice_least_active hash_policy_choice_random hash_policy_choice_source_ip_stickiness" must be set
  hash_policy_choice_round_robin = true
  // One of the arguments from this list "tls_tcp tcp tls_tcp_auto_cert" must be set
  tcp = true
  // One of the arguments from this list "no_sni sni default_lb_with_sni" must be set
  no_sni = true
}

```

Argument Reference
------------------

### Metadata Argument Reference

`annotations` - (Optional) queryable and should be preserved when modifying objects. (`String`).

`description` - (Optional) Human readable description for the object (`String`).

`disable` - (Optional) A value of true will administratively disable the object (`Bool`).

`labels` - (Optional) by selector expression (`String`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

`namespace` - (Optional) Must be a DNS_LABEL format. For a namespace object itself, namespace value will be "" (`String`).

### Spec Argument Reference

`advertise_custom` - (Optional) Advertise this VIP on specific sites. See [Advertise Custom ](#advertise-custom) below for details.

`advertise_on_public` - (Optional) Advertise this load balancer on public network. See [Advertise On Public ](#advertise-on-public) below for details.

`advertise_on_public_default_vip` - (Optional) Advertise this load balancer on public network with default VIP (bool).

`do_not_advertise` - (Optional) Do not advertise this load balancer (bool).

`do_not_retract_cluster` - (Optional) configuration. (bool).

`retract_cluster` - (Optional) for route (bool).

`dns_volterra_managed` - (Optional) This requires the domain to be delegated to Volterra using the Delegated Domain feature. (`Bool`).

`domains` - (Optional) Domains also indicate the list of names for which DNS resolution will be done by VER (`List of String`).

`hash_policy_choice_least_active` - (Optional) Connections are sent to origin server that has least active connections (bool).

`hash_policy_choice_random` - (Optional) Connections are sent to all eligible origin servers in random fashion (bool).

`hash_policy_choice_round_robin` - (Optional) Connections are sent to all eligible origin servers in round robin fashion (bool).

`hash_policy_choice_source_ip_stickiness` - (Optional) Connections are sent to all eligible origin servers using hash of source ip. Consistent hashing algorithm, ring hash, is used to select origin server (bool).

`idle_timeout` - (Optional) The amount of time that a stream can exist without upstream or downstream activity, in milliseconds. (`Int`).

`listen_port` - (Optional) Listen Port for this load balancer (`Int`).

`tcp` - (Optional) TCP Load Balancer. (bool).

`tls_tcp` - (Optional) User is responsible for managing DNS to this load balancer.. See [Tls Tcp ](#tls-tcp) below for details.

`tls_tcp_auto_cert` - (Optional) or a DNS CNAME record should be created in your DNS provider's portal.. See [Tls Tcp Auto Cert ](#tls-tcp-auto-cert) below for details.

`origin_pools_weights` - (Optional) Origin pools and weights used for this load balancer.. See [Origin Pools Weights ](#origin-pools-weights) below for details.

`default_lb_with_sni` - (Optional) Also enables usage as Default LB for Non SNI Clients (bool).

`no_sni` - (Optional) Loadbalancer without Server Name Indication support (bool).

`sni` - (Optional) Enables Server Name Indication for Loadbalancer (bool).

### Advertise Custom

Advertise this VIP on specific sites.

`advertise_where` - (Required) Where should this load balancer be available. See [Advertise Where ](#advertise-where) below for details.

### Advertise On Public

Advertise this load balancer on public network.

`public_ip` - (Required) Dedicated Public IP, which is allocated by F5 Distributed Cloud on request, is used as a VIP.. See [ref](#ref) below for details.

### Advertise Where

Where should this load balancer be available.

`site` - (Optional) Advertise on a customer site and a given network.. See [Site ](#site) below for details.

`virtual_network` - (Optional) Advertise on a virtual network. See [Virtual Network ](#virtual-network) below for details.

`virtual_site` - (Optional) Advertise on a customer virtual site and a given network.. See [Virtual Site ](#virtual-site) below for details.

`vk8s_service` - (Optional) Advertise on vK8s Service Network on RE.. See [Vk8s Service ](#vk8s-service) below for details.

`port` - (Optional) TCP port to Listen. (`Int`).

`use_default_port` - (Optional) For HTTP, default is 80. For HTTPS/SNI, default is 443. (bool).

### Blindfold Secret Info

Blindfold Secret is used for the secrets managed by Volterra Secret Management Service.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Clear Secret Info

Clear Secret is used for the secrets that are not encrypted.

`provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

`url` - (Required) When asked for this secret, caller will get Secret bytes after Base64 decoding. (`String`).

### Custom Hash Algorithms

Use hash algorithms in the custom order. Volterra will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set..

`hash_algorithms` - (Required) Ordered list of hash algorithms to be used. (`List of Strings`).

### Custom Security

Custom selection of TLS versions and cipher suites.

`cipher_suites` - (Required) The TLS listener will only support the specified cipher list. (`String`).

`max_version` - (Optional) Maximum TLS protocol version. (`String`).

`min_version` - (Optional) Minimum TLS protocol version. (`String`).

### Default Security

TLS v1.2+ with PFS ciphers and strong crypto algorithms..

### Default Vip

Use the default VIP, system allocated or configured in the virtual network.

### Disable Ocsp Stapling

This is the default behavior if no choice is selected..

### Low Security

TLS v1.0+ including non-PFS ciphers and weak crypto algorithms..

### Medium Security

TLS v1.0+ with PFS ciphers and medium strength crypto algorithms..

### No Crl

Client certificate revocation status is not verified.

### No Mtls

x-displayName: "Disable".

### Origin Pools Weights

Origin pools and weights used for this load balancer..

`endpoint_subsets` - (Optional) upstream origin pool which match this metadata will be selected for load balancing (`String`).

`cluster` - (Required) More flexible, advanced feature control with cluster. See [ref](#ref) below for details.

`pool` - (Required) Simple, commonly used pool parameters with origin pool. See [ref](#ref) below for details.

`priority` - (Optional) made active as per the increasing priority. (`Int`).

`weight` - (Optional) Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool (`Int`).

### Private Key

TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by Volterra Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in Volterra Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Site

Advertise on a customer site and a given network..

`ip` - (Optional) Use given IP address as VIP on the site (`String`).

`network` - (Required) By default VIP chosen as ip address of primary network interface in the network (`String`).

`site` - (Required) Reference to site object. See [ref](#ref) below for details.

### Tls Certificates

Set of TLS certificates.

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).

`custom_hash_algorithms` - (Optional) Use hash algorithms in the custom order. Volterra will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set.. See [Custom Hash Algorithms ](#custom-hash-algorithms) below for details.

`disable_ocsp_stapling` - (Optional) This is the default behavior if no choice is selected.. See [Disable Ocsp Stapling ](#disable-ocsp-stapling) below for details.

`use_system_defaults` - (Optional) Volterra will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order.. See [Use System Defaults ](#use-system-defaults) below for details.

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Private Key ](#private-key) below for details.

### Tls Config

Configuration of TLS settings such as min/max TLS version and ciphersuites.

`custom_security` - (Optional) Custom selection of TLS versions and cipher suites. See [Custom Security ](#custom-security) below for details.

`default_security` - (Optional) TLS v1.2+ with PFS ciphers and strong crypto algorithms. (bool).

`low_security` - (Optional) TLS v1.0+ including non-PFS ciphers and weak crypto algorithms. (bool).

`medium_security` - (Optional) TLS v1.0+ with PFS ciphers and medium strength crypto algorithms. (bool).

### Tls Parameters

TLS parameters for downstream connections..

`no_mtls` - (Optional) x-displayName: "Disable" (bool).

`use_mtls` - (Optional) x-displayName: "Enable". See [Use Mtls ](#use-mtls) below for details.

`tls_certificates` - (Required) Set of TLS certificates. See [Tls Certificates ](#tls-certificates) below for details.

`tls_config` - (Optional) Configuration of TLS settings such as min/max TLS version and ciphersuites. See [Tls Config ](#tls-config) below for details.

### Tls Tcp

User is responsible for managing DNS to this load balancer..

`tls_parameters` - (Optional) TLS parameters for downstream connections.. See [Tls Parameters ](#tls-parameters) below for details.

### Tls Tcp Auto Cert

or a DNS CNAME record should be created in your DNS provider's portal..

`no_mtls` - (Optional) x-displayName: "Disable" (bool).

`use_mtls` - (Optional) x-displayName: "Enable". See [Use Mtls ](#use-mtls) below for details.

`tls_config` - (Optional) Configuration of TLS settings such as min/max TLS version and ciphersuites. See [Tls Config ](#tls-config) below for details.

### Use Default Port

For HTTP, default is 80. For HTTPS/SNI, default is 443..

### Use Mtls

x-displayName: "Enable".

`crl` - (Optional) Specify the CRL server information to download the certificate revocation list. See [ref](#ref) below for details.

`no_crl` - (Optional) Client certificate revocation status is not verified (bool).

`trusted_ca_url` - (Required) The URL for a trust store (`String`).

### Use System Defaults

Volterra will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order..

### Vault Secret Info

Vault Secret is used for the secrets managed by Hashicorp Vault.

`key` - (Optional) If not provided entire secret will be returned. (`String`).

`location` - (Required) Path to secret in Vault. (`String`).

`provider` - (Required) Name of the Secret Management Access object that contains information about the backend Vault. (`String`).

`secret_encoding` - (Optional) This field defines the encoding type of the secret BEFORE the secret is put into Hashicorp Vault. (`String`).

`version` - (Optional) If not provided latest version will be returned. (`Int`).

### Virtual Network

Advertise on a virtual network.

`default_vip` - (Optional) Use the default VIP, system allocated or configured in the virtual network (bool).

`specific_vip` - (Optional) Use given IP address as VIP on VoltADN private Network (`String`).

`virtual_network` - (Required) Select virtual network reference. See [ref](#ref) below for details.

### Virtual Site

Advertise on a customer virtual site and a given network..

`network` - (Required) IP address of primary network interface in the network (`String`).

`virtual_site` - (Required) Reference to virtual site object. See [ref](#ref) below for details.

### Vk8s Service

Advertise on vK8s Service Network on RE..

`site` - (Optional) Reference to site object. See [ref](#ref) below for details.

`virtual_site` - (Optional) Reference to virtual site object. See [ref](#ref) below for details.

### Wingman Secret Info

Secret is given as bootstrap secret in Volterra Security Sidecar.

`name` - (Required) Name of the secret. (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured tcp_loadbalancer.
