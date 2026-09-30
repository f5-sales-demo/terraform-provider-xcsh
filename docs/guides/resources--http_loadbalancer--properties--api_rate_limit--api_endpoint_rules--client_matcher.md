---
page_title: "api_rate_limit.api_endpoint_rules.client_matcher"
subcategory: "Load Balancing"
description: "api_rate_limit.api_endpoint_rules.client_matcher for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 5533, "body_sha256": "sha256:b49ace9a14f71531e176039d1ce2383a48aa19721d52362d96a77b0e6316e1fe", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher:any_client", "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher:any_ip", "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher:asn_list", "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher:asn_matcher", "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher:client_selector", "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher:ip_matcher", "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher:ip_prefix_list", "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher:ip_threat_category_list", "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher:tls_fingerprint_matcher"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules", "path": "docs/guides/resources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--client_matcher.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_rate_limit", "api_endpoint_rules", "client_matcher"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_rate_limit/api_endpoint_rules/client_matcher/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_rate_limit.api_endpoint_rules.client_matcher for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# api_rate_limit.api_endpoint_rules.client_matcher

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [api_rate_limit](resources--http_loadbalancer--properties--api_rate_limit.md)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules.md)
- api_rate_limit.api_endpoint_rules.client_matcher

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Client Matcher. Client conditions for matching a rule.

Upstream description:

Client conditions for matching a rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("any_client",
    "client_selector"),
  validators.ConflictingObjectAttributes("any_client",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("any_ip",
    "asn_list"),
  validators.ConflictingObjectAttributes("any_ip",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("asn_list",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("asn_list",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("asn_list",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("asn_matcher",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("asn_matcher",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("client_selector",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("ip_matcher",
    "ip_prefix_list")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-client_choice": "[\"any_client\",\"client_selector\",\"ip_threat_category_list\"]",
  "x-ves-oneof-field-ip_asn_choice": "[\"any_ip\",\"asn_list\",\"asn_matcher\",\"ip_matcher\",\"ip_prefix_list\"]"
}
```

Terraform syntax:

```terraform
client_matcher {
  # Configure direct properties listed below.
}
```

## Direct properties

- [any_client](resources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--client_matcher--any_client.md): complete subsection reference.

- [any_ip](resources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--client_matcher--any_ip.md): complete subsection reference.

- [asn_list](resources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--client_matcher--asn_list.md): complete subsection reference.

- [asn_matcher](resources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--client_matcher--asn_matcher.md): complete subsection reference.

- [client_selector](resources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--client_matcher--client_selector.md): complete subsection reference.

- [ip_matcher](resources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--client_matcher--ip_matcher.md): complete subsection reference.

- [ip_prefix_list](resources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--client_matcher--ip_prefix_list.md): complete subsection reference.

- [ip_threat_category_list](resources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--client_matcher--ip_threat_category_list.md): complete subsection reference.

- [tls_fingerprint_matcher](resources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--client_matcher--tls_fingerprint_matcher.md): complete subsection reference.

## Next pages

- [api_rate_limit.api_endpoint_rules.client_matcher.any_client](resources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--client_matcher--any_client.md)
- [api_rate_limit.api_endpoint_rules.client_matcher.any_ip](resources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--client_matcher--any_ip.md)
- [api_rate_limit.api_endpoint_rules.client_matcher.asn_list](resources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--client_matcher--asn_list.md)
- [api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher](resources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--client_matcher--asn_matcher.md)
- [api_rate_limit.api_endpoint_rules.client_matcher.client_selector](resources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--client_matcher--client_selector.md)
- [api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher](resources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--client_matcher--ip_matcher.md)
- [api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list](resources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--client_matcher--ip_prefix_list.md)
- [api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list](resources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--client_matcher--ip_threat_category_list.md)
- [api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher](resources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--client_matcher--tls_fingerprint_matcher.md)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
