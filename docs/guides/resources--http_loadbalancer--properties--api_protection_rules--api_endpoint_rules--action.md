---
page_title: "api_protection_rules.api_endpoint_rules.action"
subcategory: "Load Balancing"
description: "api_protection_rules.api_endpoint_rules.action for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2012, "body_sha256": "sha256:04c141a1e7982fa58eac349ed38045b222a5ac5ecebd18a55033bafea38eb296", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:action", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:action:allow", "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:action:deny"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:action", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules", "path": "docs/guides/resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules--action.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_protection_rules", "api_endpoint_rules", "action"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_protection_rules.api_endpoint_rules.action for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_protection_rules.api_endpoint_rules.action

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [api_protection_rules](resources--http_loadbalancer--properties--api_protection_rules.md)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules.md)
- api_protection_rules.api_endpoint_rules.action

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

The action to take if the input request matches the rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("allow",
    "deny")}
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
  "x-ves-oneof-field-action": "[\"allow\",\"deny\"]"
}
```

Terraform syntax:

```terraform
action {
  # Configure direct properties listed below.
}
```

## Direct properties

- [allow](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules--action--allow.md): complete subsection reference.

- [deny](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules--action--deny.md): complete subsection reference.

## Next pages

- [api_protection_rules.api_endpoint_rules.action.allow](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules--action--allow.md)
- [api_protection_rules.api_endpoint_rules.action.deny](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules--action--deny.md)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
