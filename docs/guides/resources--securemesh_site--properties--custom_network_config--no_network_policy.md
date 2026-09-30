---
page_title: "custom_network_config.no_network_policy"
subcategory: ""
description: "custom_network_config.no_network_policy for xcsh_securemesh_site."
xcsh_docs: {"aliases": [], "body_bytes": 988, "body_sha256": "sha256:4e299f9098ff24930c5211bc20abce06e78534ada206c3dd4d975646e488a901", "canonical_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:no_network_policy", "child_ids": [], "collection_id": "xcsh-docs:resources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:no_network_policy", "parent_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config", "path": "docs/guides/resources--securemesh_site--properties--custom_network_config--no_network_policy.md", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_network_config", "no_network_policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site/properties/custom_network_config/no_network_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_network_config.no_network_policy for xcsh_securemesh_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# custom_network_config.no_network_policy

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md)
- [Property reference](resources--securemesh_site--reference.md)
- [custom_network_config](resources--securemesh_site--properties--custom_network_config.md)
- custom_network_config.no_network_policy

<a id="section"></a>

Type: `["object", {}]`. Optional.

Policy configuration for this feature.

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
no_network_policy = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [custom_network_config](resources--securemesh_site--properties--custom_network_config.md)
- [xcsh_securemesh_site](../resources/securemesh_site.md)
