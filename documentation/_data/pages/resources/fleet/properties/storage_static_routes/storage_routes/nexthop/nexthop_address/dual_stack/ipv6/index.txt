---
page_title: "storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6"
subcategory: ""
description: "IPv6 Address specified as hexadecimal numbers separated by ':'"
xcsh_docs: {"aliases": ["storage static routes storage routes nexthop nexthop address dual stack ipv6"], "body_bytes": 3052, "body_sha256": "sha256:2caa6d825134a99610d6a8b90c5bb6090b513f672b4cafc5b709851d6d6859db", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:dual_stack:ipv6", "parent_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:dual_stack", "path": "documentation/resources/fleet/properties/storage_static_routes/storage_routes/nexthop/nexthop_address/dual_stack/ipv6/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-2321102211320030-0223002220011332-1230213310102021-1131031000000012-3001331312132000-2000103313112030-3020012022321232-1101133321231212", "registry_path": "docs/guides/resources--fleet--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["storage_static_routes", "storage_routes", "nexthop", "nexthop_address", "dual_stack", "ipv6"], "schema_version": 1, "sections": [{"aliases": ["storage static routes storage routes nexthop nexthop address dual stack ipv6 addr"], "anchor": "schema-storage_static_routes--storage_routes--nexthop--nexthop_address--dual_stack--ipv6--addr", "description": "IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by ':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes '2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'", "document_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:dual_stack:ipv6", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_static_routes", "storage_routes", "nexthop", "nexthop_address", "dual_stack", "ipv6", "addr"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_static_routes/storage_routes/nexthop/nexthop_address/dual_stack/ipv6/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "IPv6 Address specified as hexadecimal numbers separated by ':'", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["fleetCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [storage_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_static_routes/)
- [storage_static_routes.storage_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_static_routes/storage_routes/)
- [storage_static_routes.storage_routes.nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_static_routes/storage_routes/nexthop/)
- [storage_static_routes.storage_routes.nexthop.nexthop_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_static_routes/storage_routes/nexthop/nexthop_address/)
- [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_static_routes/storage_routes/nexthop/nexthop_address/dual_stack/)
- storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

IPv6 Address specified as hexadecimal numbers separated by ':'.

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
ipv6 {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-storage_static_routes--storage_routes--nexthop--nexthop_address--dual_stack--ipv6--addr"></a>

### addr property

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```
