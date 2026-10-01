---
page_title: "api_protection_rules.api_endpoint_rules.client_matcher"
subcategory: "Load Balancing"
description: "api_protection_rules.api_endpoint_rules.client_matcher for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 5842, "body_sha256": "sha256:00f047fde6824a3b54fbea1b777ae7dd4a43d29df805d3bdd869c561ab5e384e", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:client_matcher", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:client_matcher:any_client", "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:client_matcher:any_ip", "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:client_matcher:asn_list", "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:client_matcher:asn_matcher", "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:client_matcher:client_selector", "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:client_matcher:ip_matcher", "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:client_matcher:ip_prefix_list", "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:client_matcher:ip_threat_category_list", "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:client_matcher:tls_fingerprint_matcher"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:client_matcher", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules", "path": "docs/guides/resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules--client_matcher.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_protection_rules", "api_endpoint_rules", "client_matcher"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/client_matcher/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_protection_rules.api_endpoint_rules.client_matcher for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_protection_rules.api_endpoint_rules.client_matcher

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [api_protection_rules](resources--http_loadbalancer--properties--api_protection_rules.md)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules.md)
- api_protection_rules.api_endpoint_rules.client_matcher

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

- [any_client](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules--client_matcher--any_client.md): complete subsection reference.

- [any_ip](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules--client_matcher--any_ip.md): complete subsection reference.

- [asn_list](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules--client_matcher--asn_list.md): complete subsection reference.

- [asn_matcher](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules--client_matcher--asn_matcher.md): complete subsection reference.

- [client_selector](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules--client_matcher--client_selector.md): complete subsection reference.

- [ip_matcher](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules--client_matcher--ip_matcher.md): complete subsection reference.

- [ip_prefix_list](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules--client_matcher--ip_prefix_list.md): complete subsection reference.

- [ip_threat_category_list](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules--client_matcher--ip_threat_category_list.md): complete subsection reference.

- [tls_fingerprint_matcher](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules--client_matcher--tls_fingerprint_matcher.md): complete subsection reference.

## Next pages

- [api_protection_rules.api_endpoint_rules.client_matcher.any_client](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules--client_matcher--any_client.md)
- [api_protection_rules.api_endpoint_rules.client_matcher.any_ip](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules--client_matcher--any_ip.md)
- [api_protection_rules.api_endpoint_rules.client_matcher.asn_list](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules--client_matcher--asn_list.md)
- [api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules--client_matcher--asn_matcher.md)
- [api_protection_rules.api_endpoint_rules.client_matcher.client_selector](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules--client_matcher--client_selector.md)
- [api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules--client_matcher--ip_matcher.md)
- [api_protection_rules.api_endpoint_rules.client_matcher.ip_prefix_list](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules--client_matcher--ip_prefix_list.md)
- [api_protection_rules.api_endpoint_rules.client_matcher.ip_threat_category_list](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules--client_matcher--ip_threat_category_list.md)
- [api_protection_rules.api_endpoint_rules.client_matcher.tls_fingerprint_matcher](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules--client_matcher--tls_fingerprint_matcher.md)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
