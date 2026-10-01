---
page_title: "Property reference"
subcategory: "Load Balancing"
description: "Property reference for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 19983, "body_sha256": "sha256:e008a336057f018ff53aa2354363de3f775a31470f6ebade4e866f713e9afd4f", "canonical_id": "xcsh-docs:resources:http_loadbalancer:reference", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:active_service_policies", "xcsh-docs:resources:http_loadbalancer:properties:advertise_custom", "xcsh-docs:resources:http_loadbalancer:properties:advertise_dualstack_on_public", "xcsh-docs:resources:http_loadbalancer:properties:advertise_on_public", "xcsh-docs:resources:http_loadbalancer:properties:advertise_on_public_default_vip", "xcsh-docs:resources:http_loadbalancer:properties:advertise_v6_on_public", "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules", "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit", "xcsh-docs:resources:http_loadbalancer:properties:api_specification", "xcsh-docs:resources:http_loadbalancer:properties:api_testing", "xcsh-docs:resources:http_loadbalancer:properties:app_firewall", "xcsh-docs:resources:http_loadbalancer:properties:blocked_clients", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection", "xcsh-docs:resources:http_loadbalancer:properties:caching_policy", "xcsh-docs:resources:http_loadbalancer:properties:captcha_challenge", "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense", "xcsh-docs:resources:http_loadbalancer:properties:cookie_stickiness", "xcsh-docs:resources:http_loadbalancer:properties:cors_policy", "xcsh-docs:resources:http_loadbalancer:properties:csrf_policy", "xcsh-docs:resources:http_loadbalancer:properties:data_guard_rules", "xcsh-docs:resources:http_loadbalancer:properties:ddos_mitigation_rules", "xcsh-docs:resources:http_loadbalancer:properties:default_pool", "xcsh-docs:resources:http_loadbalancer:properties:default_pool_list", "xcsh-docs:resources:http_loadbalancer:properties:default_route_pools", "xcsh-docs:resources:http_loadbalancer:properties:default_sensitive_data_policy", "xcsh-docs:resources:http_loadbalancer:properties:disable_api_definition", "xcsh-docs:resources:http_loadbalancer:properties:disable_api_discovery", "xcsh-docs:resources:http_loadbalancer:properties:disable_api_testing", "xcsh-docs:resources:http_loadbalancer:properties:disable_bot_defense", "xcsh-docs:resources:http_loadbalancer:properties:disable_caching", "xcsh-docs:resources:http_loadbalancer:properties:disable_client_side_defense", "xcsh-docs:resources:http_loadbalancer:properties:disable_ip_reputation", "xcsh-docs:resources:http_loadbalancer:properties:disable_malicious_user_detection", "xcsh-docs:resources:http_loadbalancer:properties:disable_malware_protection", "xcsh-docs:resources:http_loadbalancer:properties:disable_rate_limit", "xcsh-docs:resources:http_loadbalancer:properties:disable_threat_mesh", "xcsh-docs:resources:http_loadbalancer:properties:disable_trust_client_ip_headers", "xcsh-docs:resources:http_loadbalancer:properties:disable_waf", "xcsh-docs:resources:http_loadbalancer:properties:do_not_advertise", "xcsh-docs:resources:http_loadbalancer:properties:enable_api_discovery", "xcsh-docs:resources:http_loadbalancer:properties:enable_challenge", "xcsh-docs:resources:http_loadbalancer:properties:enable_ip_reputation", "xcsh-docs:resources:http_loadbalancer:properties:enable_malicious_user_detection", "xcsh-docs:resources:http_loadbalancer:properties:enable_threat_mesh", "xcsh-docs:resources:http_loadbalancer:properties:enable_trust_client_ip_headers", "xcsh-docs:resources:http_loadbalancer:properties:graphql_rules", "xcsh-docs:resources:http_loadbalancer:properties:http", "xcsh-docs:resources:http_loadbalancer:properties:https", "xcsh-docs:resources:http_loadbalancer:properties:https_auto_cert", "xcsh-docs:resources:http_loadbalancer:properties:js_challenge", "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation", "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_action_block", "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_action_default", "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_action_js_challenge", "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection", "xcsh-docs:resources:http_loadbalancer:properties:least_active", "xcsh-docs:resources:http_loadbalancer:properties:malware_protection_settings", "xcsh-docs:resources:http_loadbalancer:properties:more_option", "xcsh-docs:resources:http_loadbalancer:properties:multi_lb_app", "xcsh-docs:resources:http_loadbalancer:properties:no_challenge", "xcsh-docs:resources:http_loadbalancer:properties:no_service_policies", "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge", "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies", "xcsh-docs:resources:http_loadbalancer:properties:random", "xcsh-docs:resources:http_loadbalancer:properties:rate_limit", "xcsh-docs:resources:http_loadbalancer:properties:ring_hash", "xcsh-docs:resources:http_loadbalancer:properties:round_robin", "xcsh-docs:resources:http_loadbalancer:properties:routes", "xcsh-docs:resources:http_loadbalancer:properties:sensitive_data_disclosure_rules", "xcsh-docs:resources:http_loadbalancer:properties:sensitive_data_policy", "xcsh-docs:resources:http_loadbalancer:properties:service_policies_from_namespace", "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app", "xcsh-docs:resources:http_loadbalancer:properties:slow_ddos_mitigation", "xcsh-docs:resources:http_loadbalancer:properties:source_ip_stickiness", "xcsh-docs:resources:http_loadbalancer:properties:system_default_timeouts", "xcsh-docs:resources:http_loadbalancer:properties:timeouts", "xcsh-docs:resources:http_loadbalancer:properties:trusted_clients", "xcsh-docs:resources:http_loadbalancer:properties:user_id_client_ip", "xcsh-docs:resources:http_loadbalancer:properties:user_identification", "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:reference", "parent_id": "xcsh-docs:resources:http_loadbalancer:fundamentals", "path": "docs/guides/resources--http_loadbalancer--reference--part-1.md", "projection_part": 1, "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- Property reference

## Direct properties

- [active_service_policies](resources--http_loadbalancer--properties--active_service_policies.md): complete subsection reference.

<a id="schema-add_location"></a>

### add_location property

Type: `"bool"`. Optional, Computed.

Add Location. X-example: true Appends header x-F5 Distributed Cloud-location = &lt;RE-site-name&gt;
in responses. This configuration is ignored on CE sites. Defaults to \`false\`. Server applies
default when omitted.

Upstream description:

X-example: true Appends header x-F5 Distributed Cloud-location = &lt;RE-site-name&gt; in responses.
This configuration is ignored on CE sites.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [advertise_custom](resources--http_loadbalancer--properties--advertise_custom.md): complete subsection reference.

- [advertise_dualstack_on_public](resources--http_loadbalancer--properties--advertise_dualstack_on_public.md): complete subsection reference.

- [advertise_on_public](resources--http_loadbalancer--properties--advertise_on_public.md): complete subsection reference.

- [advertise_on_public_default_vip](resources--http_loadbalancer--properties--advertise_on_public_default_vip.md): complete subsection reference.

- [advertise_v6_on_public](resources--http_loadbalancer--properties--advertise_v6_on_public.md): complete subsection reference.

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

Upstream description:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [api_protection_rules](resources--http_loadbalancer--properties--api_protection_rules.md): complete subsection reference.

- [api_rate_limit](resources--http_loadbalancer--properties--api_rate_limit.md): complete subsection reference.

- [api_specification](resources--http_loadbalancer--properties--api_specification.md): complete subsection reference.

- [api_testing](resources--http_loadbalancer--properties--api_testing.md): complete subsection reference.

- [app_firewall](resources--http_loadbalancer--properties--app_firewall.md): complete subsection reference.

- [blocked_clients](resources--http_loadbalancer--properties--blocked_clients.md): complete subsection reference.

- [bot_defense](resources--http_loadbalancer--properties--bot_defense.md): complete subsection reference.

- [bot_defense_advanced_protection](resources--http_loadbalancer--properties--bot_defense_advanced_protection.md): complete subsection reference.

- [caching_policy](resources--http_loadbalancer--properties--caching_policy.md): complete subsection reference.

- [captcha_challenge](resources--http_loadbalancer--properties--captcha_challenge.md): complete subsection reference.

- [client_side_defense](resources--http_loadbalancer--properties--client_side_defense.md): complete subsection reference.

- [cookie_stickiness](resources--http_loadbalancer--properties--cookie_stickiness.md): complete subsection reference.

- [cors_policy](resources--http_loadbalancer--properties--cors_policy.md): complete subsection reference.

- [csrf_policy](resources--http_loadbalancer--properties--csrf_policy.md): complete subsection reference.

- [data_guard_rules](resources--http_loadbalancer--properties--data_guard_rules.md): complete subsection reference.

- [ddos_mitigation_rules](resources--http_loadbalancer--properties--ddos_mitigation_rules.md): complete subsection reference.

- [default_pool](resources--http_loadbalancer--properties--default_pool.md): complete subsection reference.

- [default_pool_list](resources--http_loadbalancer--properties--default_pool_list.md): complete subsection reference.

- [default_route_pools](resources--http_loadbalancer--properties--default_route_pools.md): complete subsection reference.

- [default_sensitive_data_policy](resources--http_loadbalancer--properties--default_sensitive_data_policy.md): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Optional.

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="schema-disable"></a>

### disable property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

A value of true will administratively disable the object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [disable_api_definition](resources--http_loadbalancer--properties--disable_api_definition.md): complete subsection reference.

- [disable_api_discovery](resources--http_loadbalancer--properties--disable_api_discovery.md): complete subsection reference.

- [disable_api_testing](resources--http_loadbalancer--properties--disable_api_testing.md): complete subsection reference.

- [disable_bot_defense](resources--http_loadbalancer--properties--disable_bot_defense.md): complete subsection reference.

- [disable_caching](resources--http_loadbalancer--properties--disable_caching.md): complete subsection reference.

- [disable_client_side_defense](resources--http_loadbalancer--properties--disable_client_side_defense.md): complete subsection reference.

- [disable_ip_reputation](resources--http_loadbalancer--properties--disable_ip_reputation.md): complete subsection reference.

- [disable_malicious_user_detection](resources--http_loadbalancer--properties--disable_malicious_user_detection.md): complete subsection reference.

- [disable_malware_protection](resources--http_loadbalancer--properties--disable_malware_protection.md): complete subsection reference.

- [disable_rate_limit](resources--http_loadbalancer--properties--disable_rate_limit.md): complete subsection reference.

- [disable_threat_mesh](resources--http_loadbalancer--properties--disable_threat_mesh.md): complete subsection reference.

- [disable_trust_client_ip_headers](resources--http_loadbalancer--properties--disable_trust_client_ip_headers.md): complete subsection reference.

- [disable_waf](resources--http_loadbalancer--properties--disable_waf.md): complete subsection reference.

- [do_not_advertise](resources--http_loadbalancer--properties--do_not_advertise.md): complete subsection reference.

<a id="schema-domains"></a>

### domains property

Type: `["list", "string"]`. Required.

List of Domains (host/authority header) that will be matched to load balancer. Supported Domains and
search order: 1. Exact Domain names: www&#46;example.com. 2.

Upstream description:

A list of Domains (host/authority header) that will be matched to load balancer.

Supported Domains and search order: &#8203;1. Exact Domain names: www&#46;example.com. &#8203;2.
Domains starting with a Wildcard: \*.example.com.

Not supported Domains: &#8203;- Just a Wildcard: \* &#8203;- A Wildcard and TLD with no root Domain:
\*.com. &#8203;- A Wildcard not matching a whole DNS label. E.g. \*.example.com and
\*.bar.example.com are valid Wildcards however \*bar.example.com, \*-bar.example.com, and
bar\*.example.com are all invalid.

Additional notes: A Wildcard will not match empty string. E.g. \*.example.com will match
bar.example.com and baz-bar.example.com but not .example.com. The longest Wildcards match first.
Only a single virtual host in the entire route configuration can match on \*. Also a Domain must be
unique across all virtual hosts within an advertise policy.

Domains are also used for SNI matching if the Loadbalancer type is HTTPS. Domains also indicate the
list of names for which DNS resolution will be automatically resolved to IP addresses by the system.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [enable_api_discovery](resources--http_loadbalancer--properties--enable_api_discovery.md): complete subsection reference.

- [enable_challenge](resources--http_loadbalancer--properties--enable_challenge.md): complete subsection reference.

- [enable_ip_reputation](resources--http_loadbalancer--properties--enable_ip_reputation.md): complete subsection reference.

- [enable_malicious_user_detection](resources--http_loadbalancer--properties--enable_malicious_user_detection.md): complete subsection reference.

- [enable_threat_mesh](resources--http_loadbalancer--properties--enable_threat_mesh.md): complete subsection reference.

- [enable_trust_client_ip_headers](resources--http_loadbalancer--properties--enable_trust_client_ip_headers.md): complete subsection reference.

- [graphql_rules](resources--http_loadbalancer--properties--graphql_rules.md): complete subsection reference.

- [http](resources--http_loadbalancer--properties--http.md): complete subsection reference.

- [https](resources--http_loadbalancer--properties--https.md): complete subsection reference.

- [https_auto_cert](resources--http_loadbalancer--properties--https_auto_cert.md): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [js_challenge](resources--http_loadbalancer--properties--js_challenge.md): complete subsection reference.

- [jwt_validation](resources--http_loadbalancer--properties--jwt_validation.md): complete subsection reference.

- [l7_ddos_action_block](resources--http_loadbalancer--properties--l7_ddos_action_block.md): complete subsection reference.

- [l7_ddos_action_default](resources--http_loadbalancer--properties--l7_ddos_action_default.md): complete subsection reference.

- [l7_ddos_action_js_challenge](resources--http_loadbalancer--properties--l7_ddos_action_js_challenge.md): complete subsection reference.

- [l7_ddos_protection](resources--http_loadbalancer--properties--l7_ddos_protection.md): complete subsection reference.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [least_active](resources--http_loadbalancer--properties--least_active.md): complete subsection reference.

- [malware_protection_settings](resources--http_loadbalancer--properties--malware_protection_settings.md): complete subsection reference.

- [more_option](resources--http_loadbalancer--properties--more_option.md): complete subsection reference.

- [multi_lb_app](resources--http_loadbalancer--properties--multi_lb_app.md): complete subsection reference.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the HTTP Load Balancer. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace where the HTTP Load Balancer is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [no_challenge](resources--http_loadbalancer--properties--no_challenge.md): complete subsection reference.

- [no_service_policies](resources--http_loadbalancer--properties--no_service_policies.md): complete subsection reference.

- [origin_server_subset_rule_list](resources--http_loadbalancer--properties--origin_server_subset_rule_list.md): complete subsection reference.

- [policy_based_challenge](resources--http_loadbalancer--properties--policy_based_challenge.md): complete subsection reference.

- [protected_cookies](resources--http_loadbalancer--properties--protected_cookies.md): complete subsection reference.

- [random](resources--http_loadbalancer--properties--random.md): complete subsection reference.

- [rate_limit](resources--http_loadbalancer--properties--rate_limit.md): complete subsection reference.

- [ring_hash](resources--http_loadbalancer--properties--ring_hash.md): complete subsection reference.

- [round_robin](resources--http_loadbalancer--properties--round_robin.md): complete subsection reference.

- [routes](resources--http_loadbalancer--properties--routes.md): complete subsection reference.

- [sensitive_data_disclosure_rules](resources--http_loadbalancer--properties--sensitive_data_disclosure_rules.md): complete subsection reference.

- [sensitive_data_policy](resources--http_loadbalancer--properties--sensitive_data_policy.md): complete subsection reference.

- [service_policies_from_namespace](resources--http_loadbalancer--properties--service_policies_from_namespace.md): complete subsection reference.

- [single_lb_app](resources--http_loadbalancer--properties--single_lb_app.md): complete subsection reference.

- [slow_ddos_mitigation](resources--http_loadbalancer--properties--slow_ddos_mitigation.md): complete subsection reference.

- [source_ip_stickiness](resources--http_loadbalancer--properties--source_ip_stickiness.md): complete subsection reference.

- [system_default_timeouts](resources--http_loadbalancer--properties--system_default_timeouts.md): complete subsection reference.

- [timeouts](resources--http_loadbalancer--properties--timeouts.md): complete subsection reference.

- [trusted_clients](resources--http_loadbalancer--properties--trusted_clients.md): complete subsection reference.

- [user_id_client_ip](resources--http_loadbalancer--properties--user_id_client_ip.md): complete subsection reference.

- [user_identification](resources--http_loadbalancer--properties--user_identification.md): complete subsection reference.

- [waf_exclusion](resources--http_loadbalancer--properties--waf_exclusion.md): complete subsection reference.
