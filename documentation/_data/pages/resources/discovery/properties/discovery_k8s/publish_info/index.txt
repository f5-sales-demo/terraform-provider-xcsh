---
page_title: "discovery_k8s.publish_info"
subcategory: ""
description: "K8s Configuration to publish VIPs."
xcsh_docs: {"aliases": ["discovery k8s publish info"], "body_bytes": 3269, "body_sha256": "sha256:7e761b144d2aeb5bd1dbaf9448570bdbbf42f5b9526432b3046ce91a301a4cad", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:disable_spec", "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:dns_delegation", "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:publish", "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:publish_fqdns"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info", "parent_id": "xcsh-docs:resources:discovery:properties:discovery_k8s", "path": "documentation/resources/discovery/properties/discovery_k8s/publish_info/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0203000021133233-3222203312230000-3212020101320013-0023011013001011-0302303200132012-2020233333033223-1131132020211023-1000310030022113", "registry_path": "docs/guides/resources--discovery--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.publish_info:ConflictingObjectAttributes:disable_spec,dns_delegation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:disable_spec", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.publish_info:ConflictingObjectAttributes:disable_spec,publish", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:disable_spec", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.publish_info:ConflictingObjectAttributes:disable_spec,publish_fqdns", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:disable_spec", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.publish_info:ConflictingObjectAttributes:disable_spec,dns_delegation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:dns_delegation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.publish_info:ConflictingObjectAttributes:dns_delegation,publish", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:dns_delegation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.publish_info:ConflictingObjectAttributes:dns_delegation,publish_fqdns", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:dns_delegation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.publish_info:ConflictingObjectAttributes:disable_spec,publish", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:publish", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.publish_info:ConflictingObjectAttributes:dns_delegation,publish", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:publish", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.publish_info:ConflictingObjectAttributes:publish,publish_fqdns", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:publish", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.publish_info:ConflictingObjectAttributes:disable_spec,publish_fqdns", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:publish_fqdns", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.publish_info:ConflictingObjectAttributes:dns_delegation,publish_fqdns", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:publish_fqdns", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.publish_info:ConflictingObjectAttributes:publish,publish_fqdns", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:publish_fqdns", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_k8s", "publish_info"], "schema_version": 1, "sections": [{"aliases": ["disable spec"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:disable_spec", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_k8s", "publish_info", "disable_spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["dns delegation"], "anchor": "section", "description": "Configuration parameter for dns delegation.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:dns_delegation", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-discovery_k8s--publish_info--dns_delegation--subdomain", "enforcement": "provider-schema", "group": "discovery_k8s.publish_info.dns_delegation:RequiredObjectAttributes:subdomain", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:dns_delegation", "type": "requires"}], "schema_path": ["discovery_k8s", "publish_info", "dns_delegation"], "syntax": "block", "type": "object"}, {"aliases": ["publish"], "anchor": "section", "description": "K8SPublishType.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:publish", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-discovery_k8s--publish_info--publish--namespace", "enforcement": "provider-schema", "group": "discovery_k8s.publish_info.publish:RequiredObjectAttributes:namespace", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:publish", "type": "requires"}], "schema_path": ["discovery_k8s", "publish_info", "publish"], "syntax": "block", "type": "object"}, {"aliases": ["publish fqdns"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:publish_fqdns", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_k8s", "publish_info", "publish_fqdns"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/properties/discovery_k8s/publish_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "K8s Configuration to publish VIPs.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_k8s.publish_info

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/)
- [discovery_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/)
- discovery_k8s.publish_info

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for publish info.

Upstream description:

K8s Configuration to publish VIPs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_spec",
    "dns_delegation"),
  validators.ConflictingObjectAttributes("disable_spec",
    "publish"),
  validators.ConflictingObjectAttributes("disable_spec",
    "publish_fqdns"),
  validators.ConflictingObjectAttributes("dns_delegation",
    "publish"),
  validators.ConflictingObjectAttributes("dns_delegation",
    "publish_fqdns"),
  validators.ConflictingObjectAttributes("publish",
    "publish_fqdns")}
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
  "x-ves-oneof-field-publish_choice": "[\"disable\",\"dns_delegation\",\"publish\",\"publish_fqdns\"]"
}
```

Terraform syntax:

```terraform
publish_info {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/publish_info/disable_spec/): complete subsection reference.

- [dns_delegation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/publish_info/dns_delegation/): complete subsection reference.

- [publish](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/publish_info/publish/): complete subsection reference.

- [publish_fqdns](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/publish_info/publish_fqdns/): complete subsection reference.

## Next pages

- [discovery_k8s.publish_info.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/publish_info/disable_spec/)
- [discovery_k8s.publish_info.dns_delegation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/publish_info/dns_delegation/)
- [discovery_k8s.publish_info.publish](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/publish_info/publish/)
- [discovery_k8s.publish_info.publish_fqdns](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/publish_info/publish_fqdns/)
- [discovery_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/)
- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
