---
page_title: "ring_hash"
subcategory: "Load Balancing"
description: "List of hash policy rules."
xcsh_docs: {"aliases": ["ring hash"], "body_bytes": 1131, "body_sha256": "sha256:e52c521b56d998e75d68f531cd7d20245bf8a2c083dfc07102bb763a28f485a8", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "documentation/resources/http_loadbalancer/properties/ring_hash/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-024.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ring_hash:RequiredObjectAttributes:hash_policy", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ring_hash"], "schema_version": 1, "sections": [{"aliases": ["ring hash hash policy"], "anchor": "section", "description": "Specifies a list of hash policies to use for ring hash load balancing. Each hash policy is evaluated individually and the combined result is used to route the request.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-ring_hash--hash_policy--header_name", "enforcement": "provider-schema", "group": "ring_hash.hash_policy:ConflictingListObjectAttributes:cookie,header_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy", "type": "conflicts"}, {"anchor": "schema-ring_hash--hash_policy--header_name", "enforcement": "provider-schema", "group": "ring_hash.hash_policy:ConflictingListObjectAttributes:header_name,source_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy", "type": "conflicts"}, {"anchor": "schema-ring_hash--hash_policy--source_ip", "enforcement": "provider-schema", "group": "ring_hash.hash_policy:ConflictingListObjectAttributes:cookie,source_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy", "type": "conflicts"}, {"anchor": "schema-ring_hash--hash_policy--source_ip", "enforcement": "provider-schema", "group": "ring_hash.hash_policy:ConflictingListObjectAttributes:header_name,source_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ring_hash.hash_policy:ConflictingListObjectAttributes:cookie,header_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ring_hash.hash_policy:ConflictingListObjectAttributes:cookie,source_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie", "type": "conflicts"}], "schema_path": ["ring_hash", "hash_policy"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/ring_hash/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "List of hash policy rules.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ring_hash

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- ring_hash

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Hash Policy List. List of hash policy rules.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("hash_policy")}
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
ring_hash {
  # Configure direct properties listed below.
}
```

## Direct properties

- [hash_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/ring_hash/hash_policy/): complete subsection reference.
