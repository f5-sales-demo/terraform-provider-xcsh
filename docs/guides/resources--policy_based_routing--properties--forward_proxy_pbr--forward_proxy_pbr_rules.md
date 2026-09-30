---
page_title: "forward_proxy_pbr.forward_proxy_pbr_rules"
subcategory: ""
description: "forward_proxy_pbr.forward_proxy_pbr_rules for xcsh_policy_based_routing."
xcsh_docs: {"aliases": [], "body_bytes": 5385, "body_sha256": "sha256:f6372cecfc333fa42bee4f4f066660485e847259a6ff758227647c8bfa56d5df", "canonical_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules", "child_ids": ["xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:all_destinations", "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:all_sources", "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:forwarding_class_list", "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:http_list", "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:ip_prefix_set", "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:label_selector", "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:metadata", "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:prefix_list", "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:tls_list"], "collection_id": "xcsh-docs:resources:policy_based_routing:collection", "completeness": "complete", "id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules", "parent_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr", "path": "docs/guides/resources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules.md", "provider_name": "policy_based_routing", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["forward_proxy_pbr", "forward_proxy_pbr_rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "forward_proxy_pbr.forward_proxy_pbr_rules for xcsh_policy_based_routing.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["policy_based_routingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# forward_proxy_pbr.forward_proxy_pbr_rules

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md)
- [Property reference](resources--policy_based_routing--reference.md)
- [forward_proxy_pbr](resources--policy_based_routing--properties--forward_proxy_pbr.md)
- forward_proxy_pbr.forward_proxy_pbr_rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

L3/L4 routing rules. Network(L3/L4) routing policy rules.

Upstream description:

Network(L3/L4) routing policy rules.

Provider validators and defaults (from schema source):

```go
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

- [all_destinations](resources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--all_destinations.md): complete subsection reference.

- [all_sources](resources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--all_sources.md): complete subsection reference.

- [forwarding_class_list](resources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--forwarding_class_list.md): complete subsection reference.

- [http_list](resources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--http_list.md): complete subsection reference.

- [ip_prefix_set](resources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--ip_prefix_set.md): complete subsection reference.

- [label_selector](resources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--label_selector.md): complete subsection reference.

- [metadata](resources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--metadata.md): complete subsection reference.

- [prefix_list](resources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--prefix_list.md): complete subsection reference.

- [tls_list](resources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--tls_list.md): complete subsection reference.

## Next pages

- [forward_proxy_pbr.forward_proxy_pbr_rules.all_destinations](resources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--all_destinations.md)
- [forward_proxy_pbr.forward_proxy_pbr_rules.all_sources](resources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--all_sources.md)
- [forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list](resources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--forwarding_class_list.md)
- [forward_proxy_pbr.forward_proxy_pbr_rules.http_list](resources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--http_list.md)
- [forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set](resources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--ip_prefix_set.md)
- [forward_proxy_pbr.forward_proxy_pbr_rules.label_selector](resources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--label_selector.md)
- [forward_proxy_pbr.forward_proxy_pbr_rules.metadata](resources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--metadata.md)
- [forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list](resources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--prefix_list.md)
- [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list](resources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--tls_list.md)
- [forward_proxy_pbr](resources--policy_based_routing--properties--forward_proxy_pbr.md)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md)
