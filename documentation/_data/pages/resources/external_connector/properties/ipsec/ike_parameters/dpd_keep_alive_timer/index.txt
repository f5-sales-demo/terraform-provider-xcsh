---
page_title: "ipsec.ike_parameters.dpd_keep_alive_timer"
subcategory: ""
description: "Configuration parameter for dpd keep alive timer."
xcsh_docs: {"aliases": ["ipsec ike parameters dpd keep alive timer"], "body_bytes": 2346, "body_sha256": "sha256:c6ff6518bda9639a8194932d30f5cd0505175f17ba416d6fb3e0d439a3db517d", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:dpd_keep_alive_timer", "parent_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters", "path": "documentation/resources/external_connector/properties/ipsec/ike_parameters/dpd_keep_alive_timer/index.md", "product": "distributed-cloud", "provider_name": "external_connector", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0201200323101322-3010031113332333-0210103123123123-1012200001010033-3021311222223022-1213032313113101-3130323023010031-0220302232322010", "registry_path": "docs/guides/resources--external_connector--reference--group-001.md", "relationships": [{"anchor": "schema-ipsec--ike_parameters--dpd_keep_alive_timer--timeout", "enforcement": "provider-schema", "group": "ipsec.ike_parameters.dpd_keep_alive_timer:RequiredObjectAttributes:timeout", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:dpd_keep_alive_timer", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ipsec", "ike_parameters", "dpd_keep_alive_timer"], "schema_version": 1, "sections": [{"aliases": ["duration", "ipsec ike parameters dpd keep alive timer timeout"], "anchor": "schema-ipsec--ike_parameters--dpd_keep_alive_timer--timeout", "description": "Operation timeout duration", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:dpd_keep_alive_timer", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ike_parameters", "dpd_keep_alive_timer", "timeout"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/external_connector/properties/ipsec/ike_parameters/dpd_keep_alive_timer/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Configuration parameter for dpd keep alive timer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["external_connectorCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ipsec.ike_parameters.dpd_keep_alive_timer

Breadcrumbs:

- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/)
- [ipsec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/)
- [ipsec.ike_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/)
- ipsec.ike_parameters.dpd_keep_alive_timer

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for dpd keep alive timer.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("timeout")}
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
dpd_keep_alive_timer {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-ipsec--ike_parameters--dpd_keep_alive_timer--timeout"></a>

### timeout property

Type: `"number"`. Optional.

Keepalive Timer. Operation timeout duration

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 5),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  }
}
```
