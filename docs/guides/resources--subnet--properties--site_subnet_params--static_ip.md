---
page_title: "site_subnet_params.static_ip"
subcategory: ""
description: "site_subnet_params.static_ip for xcsh_subnet."
xcsh_docs: {"aliases": [], "body_bytes": 962, "body_sha256": "sha256:7132e0eef7a2bca289bee4a5246c284a93ba596905fad4ea578b6220cd14c75c", "canonical_id": "xcsh-docs:resources:subnet:properties:site_subnet_params:static_ip", "child_ids": [], "collection_id": "xcsh-docs:resources:subnet:collection", "completeness": "complete", "id": "xcsh-docs:resources:subnet:properties:site_subnet_params:static_ip", "parent_id": "xcsh-docs:resources:subnet:properties:site_subnet_params", "path": "docs/guides/resources--subnet--properties--site_subnet_params--static_ip.md", "provider_name": "subnet", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["site_subnet_params", "static_ip"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/subnet/properties/site_subnet_params/static_ip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "site_subnet_params.static_ip for xcsh_subnet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["subnetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_subnet_params.static_ip

Breadcrumbs:

- [xcsh_subnet](../resources/subnet.md)
- [Property reference](resources--subnet--reference.md)
- [site_subnet_params](resources--subnet--properties--site_subnet_params.md)
- site_subnet_params.static_ip

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
static_ip = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [site_subnet_params](resources--subnet--properties--site_subnet_params.md)
- [xcsh_subnet](../resources/subnet.md)
