---
page_title: "api_protection_rules.api_groups_rules.client_matcher.asn_matcher"
subcategory: "Load Balancing"
description: "api_protection_rules.api_groups_rules.client_matcher.asn_matcher for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1939, "body_sha256": "sha256:240c83e5687e01713ae05e32a5bb47b9e6fcb4ff9409671fe3c8150bba172022", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:asn_matcher", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:asn_matcher:asn_sets"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:asn_matcher", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher", "path": "docs/guides/resources--http_loadbalancer--properties--api_protection_rules--api_groups_rules--client_matcher--asn_matcher.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_protection_rules", "api_groups_rules", "client_matcher", "asn_matcher"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/client_matcher/asn_matcher/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_protection_rules.api_groups_rules.client_matcher.asn_matcher for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_protection_rules.api_groups_rules.client_matcher.asn_matcher

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [api_protection_rules](resources--http_loadbalancer--properties--api_protection_rules.md)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--properties--api_protection_rules--api_groups_rules.md)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--properties--api_protection_rules--api_groups_rules--client_matcher.md)
- api_protection_rules.api_groups_rules.client_matcher.asn_matcher

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Match any AS number contained in the list of bgp\_asn\_sets.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("asn_sets")}
```

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

Terraform syntax:

```terraform
asn_matcher {
  # Configure direct properties listed below.
}
```

## Direct properties

- [asn_sets](resources--http_loadbalancer--properties--api_protection_rules--api_groups_rules--client_matcher--asn_matcher--asn_sets.md): complete subsection reference.

## Next pages

- [api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets](resources--http_loadbalancer--properties--api_protection_rules--api_groups_rules--client_matcher--asn_matcher--asn_sets.md)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--properties--api_protection_rules--api_groups_rules--client_matcher.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
