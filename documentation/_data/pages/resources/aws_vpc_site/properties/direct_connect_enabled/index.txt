---
page_title: "direct_connect_enabled"
subcategory: "Infrastructure"
description: "Direct Connect Configuration."
xcsh_docs: {"aliases": ["direct connect enabled"], "body_bytes": 3497, "body_sha256": "sha256:3ff0c4e95c88af712a9625ab48998a2056daa9a0c414434a637e12181d2db90f", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_vpc_site:properties:direct_connect_enabled:auto_asn", "xcsh-docs:resources:aws_vpc_site:properties:direct_connect_enabled:hosted_vifs", "xcsh-docs:resources:aws_vpc_site:properties:direct_connect_enabled:standard_vifs"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:direct_connect_enabled", "parent_id": "xcsh-docs:resources:aws_vpc_site:reference", "path": "documentation/resources/aws_vpc_site/properties/direct_connect_enabled/index.md", "product": "distributed-cloud", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0130012221000232-1311333022303301-0020012113300213-0322011110120320-0313333031011211-1310113213303030-0201300310003031-3011331001021112", "registry_path": "docs/guides/resources--aws_vpc_site--reference--group-002.md", "relationships": [{"anchor": "schema-direct_connect_enabled--custom_asn", "enforcement": "provider-schema", "group": "direct_connect_enabled:ConflictingObjectAttributes:auto_asn,custom_asn", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:direct_connect_enabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "direct_connect_enabled:ConflictingObjectAttributes:auto_asn,custom_asn", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:direct_connect_enabled:auto_asn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "direct_connect_enabled:ConflictingObjectAttributes:hosted_vifs,standard_vifs", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:direct_connect_enabled:hosted_vifs", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "direct_connect_enabled:ConflictingObjectAttributes:hosted_vifs,standard_vifs", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:direct_connect_enabled:standard_vifs", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["direct_connect_enabled"], "schema_version": 1, "sections": [{"aliases": ["auto asn"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_vpc_site:properties:direct_connect_enabled:auto_asn", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["direct_connect_enabled", "auto_asn"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom asn"], "anchor": "schema-direct_connect_enabled--custom_asn", "description": "Exclusive with Custom Autonomous System Number.", "document_id": "xcsh-docs:resources:aws_vpc_site:properties:direct_connect_enabled", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["direct_connect_enabled", "custom_asn"], "syntax": "attribute", "type": "number"}, {"aliases": ["hosted vifs"], "anchor": "section", "description": "AWS Direct Connect Hosted VIF Configuration.", "document_id": "xcsh-docs:resources:aws_vpc_site:properties:direct_connect_enabled:hosted_vifs", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "direct_connect_enabled.hosted_vifs:ConflictingObjectAttributes:site_registration_over_direct_connect,site_registration_over_internet", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:direct_connect_enabled:hosted_vifs:site_registration_over_direct_connect", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "direct_connect_enabled.hosted_vifs:ConflictingObjectAttributes:site_registration_over_direct_connect,site_registration_over_internet", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:direct_connect_enabled:hosted_vifs:site_registration_over_internet", "type": "conflicts"}], "schema_path": ["direct_connect_enabled", "hosted_vifs"], "syntax": "block", "type": "object"}, {"aliases": ["standard vifs"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_vpc_site:properties:direct_connect_enabled:standard_vifs", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["direct_connect_enabled", "standard_vifs"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/direct_connect_enabled/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Direct Connect Configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# direct_connect_enabled

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/)
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

- [auto_asn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/direct_connect_enabled/auto_asn/): complete subsection reference.

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

- [hosted_vifs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/direct_connect_enabled/hosted_vifs/): complete subsection reference.

- [standard_vifs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/direct_connect_enabled/standard_vifs/): complete subsection reference.

## Next pages

- [direct_connect_enabled.auto_asn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/direct_connect_enabled/auto_asn/)
- [direct_connect_enabled.hosted_vifs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/direct_connect_enabled/hosted_vifs/)
- [direct_connect_enabled.standard_vifs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/direct_connect_enabled/standard_vifs/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
