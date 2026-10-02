---
page_title: "deny_list"
subcategory: "Security"
description: "URL(s) and domains policy for forward proxy for a connection type (TLS or HTTP)"
xcsh_docs: {"aliases": ["deny list"], "body_bytes": 3722, "body_sha256": "sha256:ae82fef62252e56b3a7c83a3a1a3a7ebb642431cb4ca87027c476afa33c97bb8", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:forward_proxy_policy:properties:deny_list:default_action_allow", "xcsh-docs:resources:forward_proxy_policy:properties:deny_list:default_action_deny", "xcsh-docs:resources:forward_proxy_policy:properties:deny_list:default_action_next_policy", "xcsh-docs:resources:forward_proxy_policy:properties:deny_list:dest_list", "xcsh-docs:resources:forward_proxy_policy:properties:deny_list:http_list", "xcsh-docs:resources:forward_proxy_policy:properties:deny_list:tls_list"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:forward_proxy_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:forward_proxy_policy:properties:deny_list", "parent_id": "xcsh-docs:resources:forward_proxy_policy:reference", "path": "documentation/resources/forward_proxy_policy/properties/deny_list/index.md", "product": "distributed-cloud", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3020120030321321-2120102321010031-3202220001101331-0300302332020322-0012012033030332-3200110122022021-0300012113131201-0010300210310110", "registry_path": "docs/guides/resources--forward_proxy_policy--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "deny_list:ConflictingObjectAttributes:default_action_allow,default_action_deny", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:deny_list:default_action_allow", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "deny_list:ConflictingObjectAttributes:default_action_allow,default_action_next_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:deny_list:default_action_allow", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "deny_list:ConflictingObjectAttributes:default_action_allow,default_action_deny", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:deny_list:default_action_deny", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "deny_list:ConflictingObjectAttributes:default_action_deny,default_action_next_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:deny_list:default_action_deny", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "deny_list:ConflictingObjectAttributes:default_action_allow,default_action_next_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:deny_list:default_action_next_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "deny_list:ConflictingObjectAttributes:default_action_deny,default_action_next_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:deny_list:default_action_next_policy", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["deny_list"], "schema_version": 1, "sections": [{"aliases": ["default action allow"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:forward_proxy_policy:properties:deny_list:default_action_allow", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["deny_list", "default_action_allow"], "syntax": "attribute", "type": "object"}, {"aliases": ["default action deny"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:forward_proxy_policy:properties:deny_list:default_action_deny", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["deny_list", "default_action_deny"], "syntax": "attribute", "type": "object"}, {"aliases": ["default action next policy"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:forward_proxy_policy:properties:deny_list:default_action_next_policy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["deny_list", "default_action_next_policy"], "syntax": "attribute", "type": "object"}, {"aliases": ["dest list"], "anchor": "section", "description": "L4 destinations for non-HTTP and non-TLS connections and TLS connections without SNI.", "document_id": "xcsh-docs:resources:forward_proxy_policy:properties:deny_list:dest_list", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["deny_list", "dest_list"], "syntax": "block", "type": "object"}, {"aliases": ["http list"], "anchor": "section", "description": "URLs for HTTP connections.", "document_id": "xcsh-docs:resources:forward_proxy_policy:properties:deny_list:http_list", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["deny_list", "http_list"], "syntax": "block", "type": "object"}, {"aliases": ["tls list"], "anchor": "section", "description": "Domains in SNI for TLS connections.", "document_id": "xcsh-docs:resources:forward_proxy_policy:properties:deny_list:tls_list", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["deny_list", "tls_list"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/forward_proxy_policy/properties/deny_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "URL(s) and domains policy for forward proxy for a connection type (TLS or HTTP)", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# deny_list

Breadcrumbs:

- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/)
- deny_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

URL(s) and domains policy for forward proxy for a connection type (TLS or HTTP).

Upstream description:

URL(s) and domains policy for forward proxy for a connection type (TLS or HTTP)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_action_allow",
    "default_action_deny"),
  validators.ConflictingObjectAttributes("default_action_allow",
    "default_action_next_policy"),
  validators.ConflictingObjectAttributes("default_action_deny",
    "default_action_next_policy")}
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
  "x-ves-oneof-field-default_action_choice": "[\"default_action_allow\",\"default_action_deny\",\"default_action_next_policy\"]"
}
```

Terraform syntax:

```terraform
deny_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [default_action_allow](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/deny_list/default_action_allow/): complete subsection reference.

- [default_action_deny](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/deny_list/default_action_deny/): complete subsection reference.

- [default_action_next_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/deny_list/default_action_next_policy/): complete subsection reference.

- [dest_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/deny_list/dest_list/): complete subsection reference.

- [http_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/deny_list/http_list/): complete subsection reference.

- [tls_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/deny_list/tls_list/): complete subsection reference.

## Next pages

- [deny_list.default_action_allow](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/deny_list/default_action_allow/)
- [deny_list.default_action_deny](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/deny_list/default_action_deny/)
- [deny_list.default_action_next_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/deny_list/default_action_next_policy/)
- [deny_list.dest_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/deny_list/dest_list/)
- [deny_list.http_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/deny_list/http_list/)
- [deny_list.tls_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/deny_list/tls_list/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/)
- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/)
