---
page_title: "rules.action.dynamic.elastic_ips"
subcategory: ""
description: "rules.action.dynamic.elastic_ips for xcsh_nat_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1809, "body_sha256": "sha256:dc1228a3e0a534e9a8d06f0dd598094ce96ddec67a05d267a0e53e002e545e8e", "child_ids": ["xcsh-docs:resources:nat_policy:properties:rules:action:dynamic:elastic_ips:refs"], "collection_id": "xcsh-docs:resources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:nat_policy:properties:rules:action:dynamic:elastic_ips", "parent_id": "xcsh-docs:resources:nat_policy:properties:rules:action:dynamic", "path": "documentation/resources/nat_policy/properties/rules/action/dynamic/elastic_ips/index.md", "provider_name": "nat_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["rules", "action", "dynamic", "elastic_ips"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nat_policy/properties/rules/action/dynamic/elastic_ips/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.action.dynamic.elastic_ips for xcsh_nat_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nat_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rules.action.dynamic.elastic_ips

Breadcrumbs:

- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/)
- [rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/)
- [rules.action.dynamic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/dynamic/)
- rules.action.dynamic.elastic_ips

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

List of references to Cloud Elastic IP Object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("refs")}
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
elastic_ips {
  # Configure direct properties listed below.
}
```

## Direct properties

- [refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/dynamic/elastic_ips/refs/): complete subsection reference.

## Next pages

- [rules.action.dynamic.elastic_ips.refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/dynamic/elastic_ips/refs/)
- [rules.action.dynamic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/dynamic/)
- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/)
