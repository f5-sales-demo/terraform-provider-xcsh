---
page_title: "Property reference"
subcategory: "Load Balancing"
description: "Property reference for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 19164, "body_sha256": "sha256:5c25036401f7b76de4fce1c6067f988c54031a305d11cfc6675db21a48d71545", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:reference", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:active_service_policies", "xcsh-docs:data-sources:http_loadbalancer:properties:advertise_custom", "xcsh-docs:data-sources:http_loadbalancer:properties:advertise_dualstack_on_public", "xcsh-docs:data-sources:http_loadbalancer:properties:advertise_on_public", "xcsh-docs:data-sources:http_loadbalancer:properties:advertise_on_public_default_vip", "xcsh-docs:data-sources:http_loadbalancer:properties:advertise_v6_on_public", "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules", "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit", "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification", "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing", "xcsh-docs:data-sources:http_loadbalancer:properties:app_firewall", "xcsh-docs:data-sources:http_loadbalancer:properties:blocked_clients", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection", "xcsh-docs:data-sources:http_loadbalancer:properties:caching_policy", "xcsh-docs:data-sources:http_loadbalancer:properties:captcha_challenge", "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense", "xcsh-docs:data-sources:http_loadbalancer:properties:cookie_stickiness", "xcsh-docs:data-sources:http_loadbalancer:properties:cors_policy", "xcsh-docs:data-sources:http_loadbalancer:properties:csrf_policy", "xcsh-docs:data-sources:http_loadbalancer:properties:data_guard_rules", "xcsh-docs:data-sources:http_loadbalancer:properties:ddos_mitigation_rules", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool_list", "xcsh-docs:data-sources:http_loadbalancer:properties:default_route_pools", "xcsh-docs:data-sources:http_loadbalancer:properties:default_sensitive_data_policy", "xcsh-docs:data-sources:http_loadbalancer:properties:disable_api_definition", "xcsh-docs:data-sources:http_loadbalancer:properties:disable_api_discovery", "xcsh-docs:data-sources:http_loadbalancer:properties:disable_api_testing", "xcsh-docs:data-sources:http_loadbalancer:properties:disable_bot_defense", "xcsh-docs:data-sources:http_loadbalancer:properties:disable_caching", "xcsh-docs:data-sources:http_loadbalancer:properties:disable_client_side_defense", "xcsh-docs:data-sources:http_loadbalancer:properties:disable_ip_reputation", "xcsh-docs:data-sources:http_loadbalancer:properties:disable_malicious_user_detection", "xcsh-docs:data-sources:http_loadbalancer:properties:disable_malware_protection", "xcsh-docs:data-sources:http_loadbalancer:properties:disable_rate_limit", "xcsh-docs:data-sources:http_loadbalancer:properties:disable_threat_mesh", "xcsh-docs:data-sources:http_loadbalancer:properties:disable_trust_client_ip_headers", "xcsh-docs:data-sources:http_loadbalancer:properties:disable_waf", "xcsh-docs:data-sources:http_loadbalancer:properties:do_not_advertise", "xcsh-docs:data-sources:http_loadbalancer:properties:enable_api_discovery", "xcsh-docs:data-sources:http_loadbalancer:properties:enable_challenge", "xcsh-docs:data-sources:http_loadbalancer:properties:enable_ip_reputation", "xcsh-docs:data-sources:http_loadbalancer:properties:enable_malicious_user_detection", "xcsh-docs:data-sources:http_loadbalancer:properties:enable_threat_mesh", "xcsh-docs:data-sources:http_loadbalancer:properties:enable_trust_client_ip_headers", "xcsh-docs:data-sources:http_loadbalancer:properties:graphql_rules", "xcsh-docs:data-sources:http_loadbalancer:properties:http", "xcsh-docs:data-sources:http_loadbalancer:properties:https", "xcsh-docs:data-sources:http_loadbalancer:properties:https_auto_cert", "xcsh-docs:data-sources:http_loadbalancer:properties:js_challenge", "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation", "xcsh-docs:data-sources:http_loadbalancer:properties:l7_ddos_action_block", "xcsh-docs:data-sources:http_loadbalancer:properties:l7_ddos_action_default", "xcsh-docs:data-sources:http_loadbalancer:properties:l7_ddos_action_js_challenge", "xcsh-docs:data-sources:http_loadbalancer:properties:l7_ddos_protection", "xcsh-docs:data-sources:http_loadbalancer:properties:least_active", "xcsh-docs:data-sources:http_loadbalancer:properties:malware_protection_settings", "xcsh-docs:data-sources:http_loadbalancer:properties:more_option", "xcsh-docs:data-sources:http_loadbalancer:properties:multi_lb_app", "xcsh-docs:data-sources:http_loadbalancer:properties:no_challenge", "xcsh-docs:data-sources:http_loadbalancer:properties:no_service_policies", "xcsh-docs:data-sources:http_loadbalancer:properties:origin_server_subset_rule_list", "xcsh-docs:data-sources:http_loadbalancer:properties:policy_based_challenge", "xcsh-docs:data-sources:http_loadbalancer:properties:protected_cookies", "xcsh-docs:data-sources:http_loadbalancer:properties:random", "xcsh-docs:data-sources:http_loadbalancer:properties:rate_limit", "xcsh-docs:data-sources:http_loadbalancer:properties:ring_hash", "xcsh-docs:data-sources:http_loadbalancer:properties:round_robin", "xcsh-docs:data-sources:http_loadbalancer:properties:routes", "xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_disclosure_rules", "xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_policy", "xcsh-docs:data-sources:http_loadbalancer:properties:service_policies_from_namespace", "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app", "xcsh-docs:data-sources:http_loadbalancer:properties:slow_ddos_mitigation", "xcsh-docs:data-sources:http_loadbalancer:properties:source_ip_stickiness", "xcsh-docs:data-sources:http_loadbalancer:properties:system_default_timeouts", "xcsh-docs:data-sources:http_loadbalancer:properties:trusted_clients", "xcsh-docs:data-sources:http_loadbalancer:properties:user_id_client_ip", "xcsh-docs:data-sources:http_loadbalancer:properties:user_identification", "xcsh-docs:data-sources:http_loadbalancer:properties:waf_exclusion"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:reference", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:fundamentals", "path": "docs/guides/data-sources--http_loadbalancer--reference--part-1.md", "projection_part": 1, "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- Property reference

## Direct properties

- [active_service_policies](data-sources--http_loadbalancer--properties--active_service_policies.md): complete subsection reference.

<a id="schema-add_location"></a>

### add_location property

Type: `"bool"`. Computed.

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

- [advertise_custom](data-sources--http_loadbalancer--properties--advertise_custom.md): complete subsection reference.

- [advertise_dualstack_on_public](data-sources--http_loadbalancer--properties--advertise_dualstack_on_public.md): complete subsection reference.

- [advertise_on_public](data-sources--http_loadbalancer--properties--advertise_on_public.md): complete subsection reference.

- [advertise_on_public_default_vip](data-sources--http_loadbalancer--properties--advertise_on_public_default_vip.md): complete subsection reference.

- [advertise_v6_on_public](data-sources--http_loadbalancer--properties--advertise_v6_on_public.md): complete subsection reference.

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

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

- [api_protection_rules](data-sources--http_loadbalancer--properties--api_protection_rules.md): complete subsection reference.

- [api_rate_limit](data-sources--http_loadbalancer--properties--api_rate_limit.md): complete subsection reference.

- [api_specification](data-sources--http_loadbalancer--properties--api_specification.md): complete subsection reference.

- [api_testing](data-sources--http_loadbalancer--properties--api_testing.md): complete subsection reference.

- [app_firewall](data-sources--http_loadbalancer--properties--app_firewall.md): complete subsection reference.

- [blocked_clients](data-sources--http_loadbalancer--properties--blocked_clients.md): complete subsection reference.

- [bot_defense](data-sources--http_loadbalancer--properties--bot_defense.md): complete subsection reference.

- [bot_defense_advanced_protection](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection.md): complete subsection reference.

- [caching_policy](data-sources--http_loadbalancer--properties--caching_policy.md): complete subsection reference.

- [captcha_challenge](data-sources--http_loadbalancer--properties--captcha_challenge.md): complete subsection reference.

- [client_side_defense](data-sources--http_loadbalancer--properties--client_side_defense.md): complete subsection reference.

- [cookie_stickiness](data-sources--http_loadbalancer--properties--cookie_stickiness.md): complete subsection reference.

- [cors_policy](data-sources--http_loadbalancer--properties--cors_policy.md): complete subsection reference.

- [csrf_policy](data-sources--http_loadbalancer--properties--csrf_policy.md): complete subsection reference.

- [data_guard_rules](data-sources--http_loadbalancer--properties--data_guard_rules.md): complete subsection reference.

- [ddos_mitigation_rules](data-sources--http_loadbalancer--properties--ddos_mitigation_rules.md): complete subsection reference.

- [default_pool](data-sources--http_loadbalancer--properties--default_pool.md): complete subsection reference.

- [default_pool_list](data-sources--http_loadbalancer--properties--default_pool_list.md): complete subsection reference.

- [default_route_pools](data-sources--http_loadbalancer--properties--default_route_pools.md): complete subsection reference.

- [default_sensitive_data_policy](data-sources--http_loadbalancer--properties--default_sensitive_data_policy.md): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the HTTPLoadBalancer.

Upstream description:

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

- [disable_api_definition](data-sources--http_loadbalancer--properties--disable_api_definition.md): complete subsection reference.

- [disable_api_discovery](data-sources--http_loadbalancer--properties--disable_api_discovery.md): complete subsection reference.

- [disable_api_testing](data-sources--http_loadbalancer--properties--disable_api_testing.md): complete subsection reference.

- [disable_bot_defense](data-sources--http_loadbalancer--properties--disable_bot_defense.md): complete subsection reference.

- [disable_caching](data-sources--http_loadbalancer--properties--disable_caching.md): complete subsection reference.

- [disable_client_side_defense](data-sources--http_loadbalancer--properties--disable_client_side_defense.md): complete subsection reference.

- [disable_ip_reputation](data-sources--http_loadbalancer--properties--disable_ip_reputation.md): complete subsection reference.

- [disable_malicious_user_detection](data-sources--http_loadbalancer--properties--disable_malicious_user_detection.md): complete subsection reference.

- [disable_malware_protection](data-sources--http_loadbalancer--properties--disable_malware_protection.md): complete subsection reference.

- [disable_rate_limit](data-sources--http_loadbalancer--properties--disable_rate_limit.md): complete subsection reference.

- [disable_threat_mesh](data-sources--http_loadbalancer--properties--disable_threat_mesh.md): complete subsection reference.

- [disable_trust_client_ip_headers](data-sources--http_loadbalancer--properties--disable_trust_client_ip_headers.md): complete subsection reference.

- [disable_waf](data-sources--http_loadbalancer--properties--disable_waf.md): complete subsection reference.

- [do_not_advertise](data-sources--http_loadbalancer--properties--do_not_advertise.md): complete subsection reference.

<a id="schema-domains"></a>

### domains property

Type: `["list", "string"]`. Computed.

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

- [enable_api_discovery](data-sources--http_loadbalancer--properties--enable_api_discovery.md): complete subsection reference.

- [enable_challenge](data-sources--http_loadbalancer--properties--enable_challenge.md): complete subsection reference.

- [enable_ip_reputation](data-sources--http_loadbalancer--properties--enable_ip_reputation.md): complete subsection reference.

- [enable_malicious_user_detection](data-sources--http_loadbalancer--properties--enable_malicious_user_detection.md): complete subsection reference.

- [enable_threat_mesh](data-sources--http_loadbalancer--properties--enable_threat_mesh.md): complete subsection reference.

- [enable_trust_client_ip_headers](data-sources--http_loadbalancer--properties--enable_trust_client_ip_headers.md): complete subsection reference.

- [graphql_rules](data-sources--http_loadbalancer--properties--graphql_rules.md): complete subsection reference.

- [http](data-sources--http_loadbalancer--properties--http.md): complete subsection reference.

- [https](data-sources--http_loadbalancer--properties--https.md): complete subsection reference.

- [https_auto_cert](data-sources--http_loadbalancer--properties--https_auto_cert.md): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [js_challenge](data-sources--http_loadbalancer--properties--js_challenge.md): complete subsection reference.

- [jwt_validation](data-sources--http_loadbalancer--properties--jwt_validation.md): complete subsection reference.

- [l7_ddos_action_block](data-sources--http_loadbalancer--properties--l7_ddos_action_block.md): complete subsection reference.

- [l7_ddos_action_default](data-sources--http_loadbalancer--properties--l7_ddos_action_default.md): complete subsection reference.

- [l7_ddos_action_js_challenge](data-sources--http_loadbalancer--properties--l7_ddos_action_js_challenge.md): complete subsection reference.

- [l7_ddos_protection](data-sources--http_loadbalancer--properties--l7_ddos_protection.md): complete subsection reference.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

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

- [least_active](data-sources--http_loadbalancer--properties--least_active.md): complete subsection reference.

- [malware_protection_settings](data-sources--http_loadbalancer--properties--malware_protection_settings.md): complete subsection reference.

- [more_option](data-sources--http_loadbalancer--properties--more_option.md): complete subsection reference.

- [multi_lb_app](data-sources--http_loadbalancer--properties--multi_lb_app.md): complete subsection reference.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the HTTPLoadBalancer.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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

Namespace where the HTTPLoadBalancer exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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

- [no_challenge](data-sources--http_loadbalancer--properties--no_challenge.md): complete subsection reference.

- [no_service_policies](data-sources--http_loadbalancer--properties--no_service_policies.md): complete subsection reference.

- [origin_server_subset_rule_list](data-sources--http_loadbalancer--properties--origin_server_subset_rule_list.md): complete subsection reference.

- [policy_based_challenge](data-sources--http_loadbalancer--properties--policy_based_challenge.md): complete subsection reference.

- [protected_cookies](data-sources--http_loadbalancer--properties--protected_cookies.md): complete subsection reference.

- [random](data-sources--http_loadbalancer--properties--random.md): complete subsection reference.

- [rate_limit](data-sources--http_loadbalancer--properties--rate_limit.md): complete subsection reference.

- [ring_hash](data-sources--http_loadbalancer--properties--ring_hash.md): complete subsection reference.

- [round_robin](data-sources--http_loadbalancer--properties--round_robin.md): complete subsection reference.

- [routes](data-sources--http_loadbalancer--properties--routes.md): complete subsection reference.

- [sensitive_data_disclosure_rules](data-sources--http_loadbalancer--properties--sensitive_data_disclosure_rules.md): complete subsection reference.

- [sensitive_data_policy](data-sources--http_loadbalancer--properties--sensitive_data_policy.md): complete subsection reference.

- [service_policies_from_namespace](data-sources--http_loadbalancer--properties--service_policies_from_namespace.md): complete subsection reference.

- [single_lb_app](data-sources--http_loadbalancer--properties--single_lb_app.md): complete subsection reference.

- [slow_ddos_mitigation](data-sources--http_loadbalancer--properties--slow_ddos_mitigation.md): complete subsection reference.

- [source_ip_stickiness](data-sources--http_loadbalancer--properties--source_ip_stickiness.md): complete subsection reference.

- [system_default_timeouts](data-sources--http_loadbalancer--properties--system_default_timeouts.md): complete subsection reference.

- [trusted_clients](data-sources--http_loadbalancer--properties--trusted_clients.md): complete subsection reference.

- [user_id_client_ip](data-sources--http_loadbalancer--properties--user_id_client_ip.md): complete subsection reference.

- [user_identification](data-sources--http_loadbalancer--properties--user_identification.md): complete subsection reference.

- [waf_exclusion](data-sources--http_loadbalancer--properties--waf_exclusion.md): complete subsection reference.
