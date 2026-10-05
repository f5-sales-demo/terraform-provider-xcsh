---
page_title: "virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address"
subcategory: ""
description: "Destination Address Mask."
xcsh_docs: {"aliases": ["virtual server connection rate limit mode per virtual server destination address"], "body_bytes": 2734, "body_sha256": "sha256:fbb817281ee0cb3c14144d240f8cb7da8dce98589aa15950c300e1df6169bd73", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_virtual_server_destination_address", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:connection_rate_limit_mode", "path": "documentation/resources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_virtual_server_destination_address/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0301121332212300-0113023232303112-2333110211120123-2123112112111322-2120111311301310-3012013311220121-1203221022201203-1231323231200220", "registry_path": "docs/guides/resources--application_profiles--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["virtual_server", "connection_rate_limit_mode", "per_virtual_server_destination_address"], "schema_version": 1, "sections": [{"aliases": ["virtual server connection rate limit mode per virtual server destination address destination mask"], "anchor": "schema-virtual_server--connection_rate_limit_mode--per_virtual_server_destination_address--destination_mask", "description": "Configuration parameter for destination mask", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_virtual_server_destination_address", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "connection_rate_limit_mode", "per_virtual_server_destination_address", "destination_mask"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_virtual_server_destination_address/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Destination Address Mask.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/)
- [virtual_server.connection_rate_limit_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/connection_rate_limit_mode/)
- virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Destination Address Mask.

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
per_virtual_server_destination_address {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-virtual_server--connection_rate_limit_mode--per_virtual_server_destination_address--destination_mask"></a>

### destination_mask property

Type: `"number"`. Optional.

Configuration parameter for destination mask.

Upstream description:

Configuration parameter for destination mask

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 600000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

## Next pages

- [virtual_server.connection_rate_limit_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/connection_rate_limit_mode/)
- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
