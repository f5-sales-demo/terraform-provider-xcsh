---
page_title: "site_subnet_params"
subcategory: ""
description: "site_subnet_params for xcsh_subnet."
xcsh_docs: {"aliases": [], "body_bytes": 2579, "body_sha256": "sha256:100e06034d91fa9af9fce896483f035298d3949a64189c7c49a9cedd04c03cca", "canonical_id": "xcsh-docs:resources:subnet:properties:site_subnet_params", "child_ids": ["xcsh-docs:resources:subnet:properties:site_subnet_params:dhcp", "xcsh-docs:resources:subnet:properties:site_subnet_params:site", "xcsh-docs:resources:subnet:properties:site_subnet_params:static_ip", "xcsh-docs:resources:subnet:properties:site_subnet_params:subnet_dhcp_server_params"], "collection_id": "xcsh-docs:resources:subnet:collection", "completeness": "complete", "id": "xcsh-docs:resources:subnet:properties:site_subnet_params", "parent_id": "xcsh-docs:resources:subnet:reference", "path": "docs/guides/resources--subnet--properties--site_subnet_params.md", "provider_name": "subnet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["site_subnet_params"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/subnet/properties/site_subnet_params/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "site_subnet_params for xcsh_subnet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["subnetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_subnet_params

Breadcrumbs:

- [xcsh_subnet](../resources/subnet.md)
- [Property reference](resources--subnet--reference.md)
- site_subnet_params

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Site Subnet Parameters. Configure subnet parameters per site.

Upstream description:

Configure subnet parameters per site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("dhcp",
    "static_ip")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
site_subnet_params {
  # Configure direct properties listed below.
}
```

## Direct properties

- [dhcp](resources--subnet--properties--site_subnet_params--dhcp.md): complete subsection reference.

- [site](resources--subnet--properties--site_subnet_params--site.md): complete subsection reference.

- [static_ip](resources--subnet--properties--site_subnet_params--static_ip.md): complete subsection reference.

- [subnet_dhcp_server_params](resources--subnet--properties--site_subnet_params--subnet_dhcp_server_params.md): complete subsection reference.

## Next pages

- [site_subnet_params.dhcp](resources--subnet--properties--site_subnet_params--dhcp.md)
- [site_subnet_params.site](resources--subnet--properties--site_subnet_params--site.md)
- [site_subnet_params.static_ip](resources--subnet--properties--site_subnet_params--static_ip.md)
- [site_subnet_params.subnet_dhcp_server_params](resources--subnet--properties--site_subnet_params--subnet_dhcp_server_params.md)
- [Property reference](resources--subnet--reference.md)
- [xcsh_subnet](../resources/subnet.md)
