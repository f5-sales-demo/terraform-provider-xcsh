---
page_title: "azure_vnet_site.vnet_attachments.vnet_list"
subcategory: ""
description: "azure_vnet_site.vnet_attachments.vnet_list for xcsh_cloud_connect."
xcsh_docs: {"aliases": [], "body_bytes": 5562, "body_sha256": "sha256:408bc308bbc1f9b1496c8bb3db28e25367fe25ed2ee21c62af90919f16c7310f", "canonical_id": "xcsh-docs:resources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list", "child_ids": ["xcsh-docs:resources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list:custom_routing", "xcsh-docs:resources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list:default_route", "xcsh-docs:resources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list:labels", "xcsh-docs:resources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list:manual_routing"], "collection_id": "xcsh-docs:resources:cloud_connect:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list", "parent_id": "xcsh-docs:resources:cloud_connect:properties:azure_vnet_site:vnet_attachments", "path": "docs/guides/resources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list.md", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["azure_vnet_site", "vnet_attachments", "vnet_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_connect/properties/azure_vnet_site/vnet_attachments/vnet_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "azure_vnet_site.vnet_attachments.vnet_list for xcsh_cloud_connect.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# azure_vnet_site.vnet_attachments.vnet_list

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md)
- [Property reference](resources--cloud_connect--reference.md)
- [azure_vnet_site](resources--cloud_connect--properties--azure_vnet_site.md)
- [azure_vnet_site.vnet_attachments](resources--cloud_connect--properties--azure_vnet_site--vnet_attachments.md)
- azure_vnet_site.vnet_attachments.vnet_list

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

VNet List. Collection of items or values

Upstream description:

Collection of items or values

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("subscription_id",
    "vnet_id"),
  validators.ConflictingListObjectAttributes("custom_routing",
    "default_route"),
  validators.ConflictingListObjectAttributes("custom_routing",
    "manual_routing"),
  validators.ConflictingListObjectAttributes("default_route",
    "manual_routing")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

Terraform syntax:

```terraform
vnet_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [custom_routing](resources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list--custom_routing.md): complete subsection reference.

- [default_route](resources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list--default_route.md): complete subsection reference.

- [labels](resources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list--labels.md): complete subsection reference.

- [manual_routing](resources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list--manual_routing.md): complete subsection reference.

<a id="schema-azure_vnet_site--vnet_attachments--vnet_list--subscription_id"></a>

### subscription_id property

Type: `"string"`. Optional.

Enter the Subscription ID of the VNet to be attached.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="schema-azure_vnet_site--vnet_attachments--vnet_list--vnet_id"></a>

### vnet_id property

Type: `"string"`. Optional.

Enter the VNet ID of the VNet to be attached in format
/&lt;resource-group-name&gt;/&lt;VNet-name&gt;.

Upstream description:

Enter the VNet ID of the VNet to be attached in format
/&lt;resource-group-name&gt;/&lt;VNet-name&gt;

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

## Next pages

- [azure_vnet_site.vnet_attachments.vnet_list.custom_routing](resources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list--custom_routing.md)
- [azure_vnet_site.vnet_attachments.vnet_list.default_route](resources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list--default_route.md)
- [azure_vnet_site.vnet_attachments.vnet_list.labels](resources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list--labels.md)
- [azure_vnet_site.vnet_attachments.vnet_list.manual_routing](resources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list--manual_routing.md)
- [azure_vnet_site.vnet_attachments](resources--cloud_connect--properties--azure_vnet_site--vnet_attachments.md)
- [xcsh_cloud_connect](../resources/cloud_connect.md)
