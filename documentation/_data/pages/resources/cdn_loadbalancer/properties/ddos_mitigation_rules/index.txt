---
page_title: "ddos_mitigation_rules"
subcategory: "Load Balancing"
description: "Define manual mitigation rules to block L7 DDoS attacks."
xcsh_docs: {"aliases": ["ddos mitigation rules"], "body_bytes": 4433, "body_sha256": "sha256:d6e0ceda039e47bb8b69a0137b53b5bb66f2fd8fac3ccdd8df84d05affd4857a", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:ddos_mitigation_rules:block", "xcsh-docs:resources:cdn_loadbalancer:properties:ddos_mitigation_rules:ddos_client_source", "xcsh-docs:resources:cdn_loadbalancer:properties:ddos_mitigation_rules:ip_prefix_list", "xcsh-docs:resources:cdn_loadbalancer:properties:ddos_mitigation_rules:metadata"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:ddos_mitigation_rules", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:reference", "path": "documentation/resources/cdn_loadbalancer/properties/ddos_mitigation_rules/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1113101222111100-0231333020320020-2111302302220332-3302220301331200-2213011011211300-1222011002133311-2332301203001000-3323001221330013", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-010.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ddos_mitigation_rules:ConflictingListObjectAttributes:ddos_client_source,ip_prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:ddos_mitigation_rules:ddos_client_source", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ddos_mitigation_rules:ConflictingListObjectAttributes:ddos_client_source,ip_prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:ddos_mitigation_rules:ip_prefix_list", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ddos_mitigation_rules"], "schema_version": 1, "sections": [{"aliases": ["block"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:ddos_mitigation_rules:block", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ddos_mitigation_rules", "block"], "syntax": "block", "type": "object"}, {"aliases": ["ddos client source"], "anchor": "section", "description": "DDoS Mitigation sources to be blocked.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:ddos_mitigation_rules:ddos_client_source", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ddos_mitigation_rules", "ddos_client_source"], "syntax": "block", "type": "object"}, {"aliases": ["expiration timestamp"], "anchor": "schema-ddos_mitigation_rules--expiration_timestamp", "description": "The expiration_timestamp is the RFC 3339 format timestamp at which the containing rule is considered to be logically expired. The rule continues to exist in the configuration but is not applied anymore.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:ddos_mitigation_rules", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ddos_mitigation_rules", "expiration_timestamp"], "syntax": "attribute", "type": "string"}, {"aliases": ["ip prefix list"], "anchor": "section", "description": "List of IP Prefix strings to match against.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:ddos_mitigation_rules:ip_prefix_list", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ddos_mitigation_rules", "ip_prefix_list"], "syntax": "block", "type": "object"}, {"aliases": ["metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:ddos_mitigation_rules:metadata", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ddos_mitigation_rules--metadata--name", "enforcement": "provider-schema", "group": "ddos_mitigation_rules.metadata:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:ddos_mitigation_rules:metadata", "type": "requires"}], "schema_path": ["ddos_mitigation_rules", "metadata"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/ddos_mitigation_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Define manual mitigation rules to block L7 DDoS attacks.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ddos_mitigation_rules

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- ddos_mitigation_rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Define manual mitigation rules to block L7 DDoS attacks.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("ddos_client_source",
    "ip_prefix_list")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
ddos_mitigation_rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [block](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/ddos_mitigation_rules/block/): complete subsection reference.

- [ddos_client_source](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/ddos_mitigation_rules/ddos_client_source/): complete subsection reference.

<a id="schema-ddos_mitigation_rules--expiration_timestamp"></a>

### expiration_timestamp property

Type: `"string"`. Optional.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Upstream description:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  }
}
```

- [ip_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/ddos_mitigation_rules/ip_prefix_list/): complete subsection reference.

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/ddos_mitigation_rules/metadata/): complete subsection reference.

## Next pages

- [ddos_mitigation_rules.block](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/ddos_mitigation_rules/block/)
- [ddos_mitigation_rules.ddos_client_source](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/ddos_mitigation_rules/ddos_client_source/)
- [ddos_mitigation_rules.ip_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/ddos_mitigation_rules/ip_prefix_list/)
- [ddos_mitigation_rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/ddos_mitigation_rules/metadata/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
