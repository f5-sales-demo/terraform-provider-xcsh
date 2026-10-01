---
page_title: "direct_connect_enabled"
subcategory: ""
description: "direct_connect_enabled for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 2989, "body_sha256": "sha256:85ebeec7089e64a2f0613aa7d37294366310239494e2ee5700d14aa663ed75d8", "canonical_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:auto_asn", "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs", "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:standard_vifs"], "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled", "parent_id": "xcsh-docs:resources:aws_tgw_site:reference", "path": "docs/guides/resources--aws_tgw_site--properties--direct_connect_enabled.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["direct_connect_enabled"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/direct_connect_enabled/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "direct_connect_enabled for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# direct_connect_enabled

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
- [Property reference](resources--aws_tgw_site--reference.md)
- direct_connect_enabled

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Direct Connect Configuration. Direct Connect Configuration.

Upstream description:

Direct Connect Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auto_asn",
    "custom_asn"),
  validators.ConflictingObjectAttributes("hosted_vifs",
    "standard_vifs")}
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
  "x-ves-oneof-field-asn_choice": "[\"auto_asn\",\"custom_asn\"]",
  "x-ves-oneof-field-vif_choice": "[\"hosted_vifs\",\"standard_vifs\"]"
}
```

Terraform syntax:

```terraform
direct_connect_enabled {
  # Configure direct properties listed below.
}
```

## Direct properties

- [auto_asn](resources--aws_tgw_site--properties--direct_connect_enabled--auto_asn.md): complete subsection reference.

<a id="schema-direct_connect_enabled--custom_asn"></a>

### custom_asn property

Type: `"number"`. Optional.

Exclusive with \[auto\_asn\] Custom Autonomous System Number.

Upstream description:

Exclusive with \[auto\_asn\] Custom Autonomous System Number.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

- [hosted_vifs](resources--aws_tgw_site--properties--direct_connect_enabled--hosted_vifs.md): complete subsection reference.

- [standard_vifs](resources--aws_tgw_site--properties--direct_connect_enabled--standard_vifs.md): complete subsection reference.

## Next pages

- [direct_connect_enabled.auto_asn](resources--aws_tgw_site--properties--direct_connect_enabled--auto_asn.md)
- [direct_connect_enabled.hosted_vifs](resources--aws_tgw_site--properties--direct_connect_enabled--hosted_vifs.md)
- [direct_connect_enabled.standard_vifs](resources--aws_tgw_site--properties--direct_connect_enabled--standard_vifs.md)
- [Property reference](resources--aws_tgw_site--reference.md)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
