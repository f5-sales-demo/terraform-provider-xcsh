---
page_title: "sensitive_data_disclosure_rules.sensitive_data_types_in_response.mask"
subcategory: "Load Balancing"
description: "sensitive_data_disclosure_rules.sensitive_data_types_in_response.mask for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1416, "body_sha256": "sha256:c319efdc69c8fd1ad33911bd7f7269a41d5670485a5f75ba2808c8fbdcc049ee", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:mask", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:mask", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response", "path": "docs/guides/resources--http_loadbalancer--properties--sensitive_data_disclosure_rules--sensitive_data_types_in_response--mask.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["sensitive_data_disclosure_rules", "sensitive_data_types_in_response", "mask"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/mask/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "sensitive_data_disclosure_rules.sensitive_data_types_in_response.mask for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# sensitive_data_disclosure_rules.sensitive_data_types_in_response.mask

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [sensitive_data_disclosure_rules](resources--http_loadbalancer--properties--sensitive_data_disclosure_rules.md)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](resources--http_loadbalancer--properties--sensitive_data_disclosure_rules--sensitive_data_types_in_response.md)
- sensitive_data_disclosure_rules.sensitive_data_types_in_response.mask

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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
mask = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](resources--http_loadbalancer--properties--sensitive_data_disclosure_rules--sensitive_data_types_in_response.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
