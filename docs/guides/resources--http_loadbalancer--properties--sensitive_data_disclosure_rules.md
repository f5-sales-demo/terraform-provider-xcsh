---
page_title: "sensitive_data_disclosure_rules"
subcategory: "Load Balancing"
description: "sensitive_data_disclosure_rules for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1310, "body_sha256": "sha256:dc38e1ff95a2d1b4f3af90d0b3c3e9e132abdf24337f36f9c02f02be1db670d1", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:sensitive_data_disclosure_rules", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:sensitive_data_disclosure_rules", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "docs/guides/resources--http_loadbalancer--properties--sensitive_data_disclosure_rules.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["sensitive_data_disclosure_rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/sensitive_data_disclosure_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "sensitive_data_disclosure_rules for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# sensitive_data_disclosure_rules

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- sensitive_data_disclosure_rules

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Sensitive Data Exposure Rules allows specifying rules to mask sensitive data fields in API
responses.

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
sensitive_data_disclosure_rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [sensitive_data_types_in_response](resources--http_loadbalancer--properties--sensitive_data_disclosure_rules--sensitive_data_types_in_response.md): complete subsection reference.

## Next pages

- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](resources--http_loadbalancer--properties--sensitive_data_disclosure_rules--sensitive_data_types_in_response.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
