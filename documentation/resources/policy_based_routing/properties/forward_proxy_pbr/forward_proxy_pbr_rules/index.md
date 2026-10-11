---
page_title: "forward_proxy_pbr.forward_proxy_pbr_rules"
subcategory: ""
description: "Network(L3/L4) routing policy rules."
xcsh_docs: {"aliases": ["forward proxy pbr forward proxy pbr rules"], "body_bytes": 4385, "body_sha256": "sha256:96c1ab68910d43a0b4b613102e3a0cffb9ff39fc6fb9aa4fefb925aa06bd568e", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:all_destinations", "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:all_sources", "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:forwarding_class_list", "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:http_list", "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:ip_prefix_set", "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:label_selector", "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:metadata", "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:prefix_list", "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:tls_list"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:policy_based_routing:collection", "completeness": "complete", "id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules", "parent_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr", "path": "documentation/resources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/index.md", "product": "distributed-cloud", "provider_name": "policy_based_routing", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-1023022331001313-2101300112322121-1321113113010332-1211211323300300-1300220233000200-1322111330001221-2023000212123212-1222001323211312", "registry_path": "docs/guides/resources--policy_based_routing--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:ConflictingListObjectAttributes:all_destinations,http_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:all_destinations", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:ConflictingListObjectAttributes:all_destinations,tls_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:all_destinations", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:ConflictingListObjectAttributes:all_sources,ip_prefix_set", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:all_sources", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:ConflictingListObjectAttributes:all_sources,label_selector", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:all_sources", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:ConflictingListObjectAttributes:all_sources,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:all_sources", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:ConflictingListObjectAttributes:all_destinations,http_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:http_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:ConflictingListObjectAttributes:http_list,tls_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:http_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:ConflictingListObjectAttributes:all_sources,ip_prefix_set", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:ip_prefix_set", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:ConflictingListObjectAttributes:ip_prefix_set,label_selector", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:ip_prefix_set", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:ConflictingListObjectAttributes:ip_prefix_set,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:ip_prefix_set", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:ConflictingListObjectAttributes:all_sources,label_selector", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:label_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:ConflictingListObjectAttributes:ip_prefix_set,label_selector", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:label_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:ConflictingListObjectAttributes:label_selector,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:label_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:ConflictingListObjectAttributes:all_sources,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:ConflictingListObjectAttributes:ip_prefix_set,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:ConflictingListObjectAttributes:label_selector,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:ConflictingListObjectAttributes:all_destinations,tls_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:tls_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:ConflictingListObjectAttributes:http_list,tls_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:tls_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:RequiredListObjectAttributes:forwarding_class_list", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:forwarding_class_list", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["forward_proxy_pbr", "forward_proxy_pbr_rules"], "schema_version": 1, "sections": [{"aliases": ["forward proxy pbr forward proxy pbr rules all destinations"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:all_destinations", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["forward_proxy_pbr", "forward_proxy_pbr_rules", "all_destinations"], "syntax": "attribute", "type": "object"}, {"aliases": ["forward proxy pbr forward proxy pbr rules all sources"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:all_sources", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["forward_proxy_pbr", "forward_proxy_pbr_rules", "all_sources"], "syntax": "attribute", "type": "object"}, {"aliases": ["forward proxy pbr forward proxy pbr rules forwarding class list"], "anchor": "section", "description": "Ordered list of forwarding Class to be used if no rule match.", "document_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:forwarding_class_list", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-forward_proxy_pbr--forward_proxy_pbr_rules--forwarding_class_list--name", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:forwarding_class_list", "type": "requires"}], "schema_path": ["forward_proxy_pbr", "forward_proxy_pbr_rules", "forwarding_class_list"], "syntax": "block", "type": "object"}, {"aliases": ["forward proxy pbr forward proxy pbr rules http list"], "anchor": "section", "description": "URLListType.", "document_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:http_list", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["forward_proxy_pbr", "forward_proxy_pbr_rules", "http_list"], "syntax": "block", "type": "object"}, {"aliases": ["forward proxy pbr forward proxy pbr rules ip prefix set"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:ip_prefix_set", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-forward_proxy_pbr--forward_proxy_pbr_rules--ip_prefix_set--name", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:ip_prefix_set", "type": "requires"}], "schema_path": ["forward_proxy_pbr", "forward_proxy_pbr_rules", "ip_prefix_set"], "syntax": "block", "type": "object"}, {"aliases": ["forward proxy pbr forward proxy pbr rules label selector"], "anchor": "section", "description": "This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expressions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects.", "document_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:label_selector", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-forward_proxy_pbr--forward_proxy_pbr_rules--label_selector--expressions", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules.label_selector:RequiredObjectAttributes:expressions", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:label_selector", "type": "requires"}], "schema_path": ["forward_proxy_pbr", "forward_proxy_pbr_rules", "label_selector"], "syntax": "block", "type": "object"}, {"aliases": ["forward proxy pbr forward proxy pbr rules metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-forward_proxy_pbr--forward_proxy_pbr_rules--metadata--name", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules.metadata:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:metadata", "type": "requires"}], "schema_path": ["forward_proxy_pbr", "forward_proxy_pbr_rules", "metadata"], "syntax": "block", "type": "object"}, {"aliases": ["forward proxy pbr forward proxy pbr rules prefix list"], "anchor": "section", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:prefix_list", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["forward_proxy_pbr", "forward_proxy_pbr_rules", "prefix_list"], "syntax": "block", "type": "object"}, {"aliases": ["forward proxy pbr forward proxy pbr rules tls list"], "anchor": "section", "description": "DomainListType.", "document_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:tls_list", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["forward_proxy_pbr", "forward_proxy_pbr_rules", "tls_list"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Network(L3/L4) routing policy rules.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["policy_based_routingCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# forward_proxy_pbr.forward_proxy_pbr_rules

Breadcrumbs:

- [xcsh_policy_based_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/)
- [forward_proxy_pbr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/forward_proxy_pbr/)
- forward_proxy_pbr.forward_proxy_pbr_rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

L3/L4 routing rules. Network(L3/L4) routing policy rules.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("forwarding_class_list"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "http_list"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "tls_list"),
  validators.ConflictingListObjectAttributes("all_sources",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("all_sources",
    "label_selector"),
  validators.ConflictingListObjectAttributes("all_sources",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("http_list",
    "tls_list"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
    "label_selector"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("label_selector",
    "prefix_list")}
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
forward_proxy_pbr_rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [all_destinations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/all_destinations/): complete subsection reference.

- [all_sources](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/all_sources/): complete subsection reference.

- [forwarding_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/forwarding_class_list/): complete subsection reference.

- [http_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/http_list/): complete subsection reference.

- [ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/ip_prefix_set/): complete subsection reference.

- [label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/label_selector/): complete subsection reference.

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/metadata/): complete subsection reference.

- [prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/prefix_list/): complete subsection reference.

- [tls_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/tls_list/): complete subsection reference.
