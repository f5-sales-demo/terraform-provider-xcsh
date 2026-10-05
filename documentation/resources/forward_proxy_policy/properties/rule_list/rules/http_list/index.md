---
page_title: "rule_list.rules.http_list"
subcategory: "Security"
description: "URLListType."
xcsh_docs: {"aliases": ["rule list rules http list"], "body_bytes": 1682, "body_sha256": "sha256:759b075f4a9a97d09b812d3fbadfa6c51c3a31c8155d49f85403cf95884f5a80", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:http_list:http_list"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:forward_proxy_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:http_list", "parent_id": "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules", "path": "documentation/resources/forward_proxy_policy/properties/rule_list/rules/http_list/index.md", "product": "distributed-cloud", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2221002323000201-0300021002302123-3110200203220001-1211121002132203-0201221220221001-1211313312213133-0310230023120021-1331032103220220", "registry_path": "docs/guides/resources--forward_proxy_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rule_list", "rules", "http_list"], "schema_version": 1, "sections": [{"aliases": ["rule list rules http list http list"], "anchor": "section", "description": "URLs for HTTP connections.", "document_id": "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:http_list:http_list", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-rule_list--rules--http_list--http_list--exact_value", "enforcement": "provider-schema", "group": "rule_list.rules.http_list.http_list:ConflictingListObjectAttributes:exact_value,regex_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:http_list:http_list", "type": "conflicts"}, {"anchor": "schema-rule_list--rules--http_list--http_list--exact_value", "enforcement": "provider-schema", "group": "rule_list.rules.http_list.http_list:ConflictingListObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:http_list:http_list", "type": "conflicts"}, {"anchor": "schema-rule_list--rules--http_list--http_list--path_exact_value", "enforcement": "provider-schema", "group": "rule_list.rules.http_list.http_list:ConflictingListObjectAttributes:any_path,path_exact_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:http_list:http_list", "type": "conflicts"}, {"anchor": "schema-rule_list--rules--http_list--http_list--path_exact_value", "enforcement": "provider-schema", "group": "rule_list.rules.http_list.http_list:ConflictingListObjectAttributes:path_exact_value,path_prefix_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:http_list:http_list", "type": "conflicts"}, {"anchor": "schema-rule_list--rules--http_list--http_list--path_exact_value", "enforcement": "provider-schema", "group": "rule_list.rules.http_list.http_list:ConflictingListObjectAttributes:path_exact_value,path_regex_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:http_list:http_list", "type": "conflicts"}, {"anchor": "schema-rule_list--rules--http_list--http_list--path_prefix_value", "enforcement": "provider-schema", "group": "rule_list.rules.http_list.http_list:ConflictingListObjectAttributes:any_path,path_prefix_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:http_list:http_list", "type": "conflicts"}, {"anchor": "schema-rule_list--rules--http_list--http_list--path_prefix_value", "enforcement": "provider-schema", "group": "rule_list.rules.http_list.http_list:ConflictingListObjectAttributes:path_exact_value,path_prefix_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:http_list:http_list", "type": "conflicts"}, {"anchor": "schema-rule_list--rules--http_list--http_list--path_prefix_value", "enforcement": "provider-schema", "group": "rule_list.rules.http_list.http_list:ConflictingListObjectAttributes:path_prefix_value,path_regex_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:http_list:http_list", "type": "conflicts"}, {"anchor": "schema-rule_list--rules--http_list--http_list--path_regex_value", "enforcement": "provider-schema", "group": "rule_list.rules.http_list.http_list:ConflictingListObjectAttributes:any_path,path_regex_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:http_list:http_list", "type": "conflicts"}, {"anchor": "schema-rule_list--rules--http_list--http_list--path_regex_value", "enforcement": "provider-schema", "group": "rule_list.rules.http_list.http_list:ConflictingListObjectAttributes:path_exact_value,path_regex_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:http_list:http_list", "type": "conflicts"}, {"anchor": "schema-rule_list--rules--http_list--http_list--path_regex_value", "enforcement": "provider-schema", "group": "rule_list.rules.http_list.http_list:ConflictingListObjectAttributes:path_prefix_value,path_regex_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:http_list:http_list", "type": "conflicts"}, {"anchor": "schema-rule_list--rules--http_list--http_list--regex_value", "enforcement": "provider-schema", "group": "rule_list.rules.http_list.http_list:ConflictingListObjectAttributes:exact_value,regex_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:http_list:http_list", "type": "conflicts"}, {"anchor": "schema-rule_list--rules--http_list--http_list--regex_value", "enforcement": "provider-schema", "group": "rule_list.rules.http_list.http_list:ConflictingListObjectAttributes:regex_value,suffix_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:http_list:http_list", "type": "conflicts"}, {"anchor": "schema-rule_list--rules--http_list--http_list--suffix_value", "enforcement": "provider-schema", "group": "rule_list.rules.http_list.http_list:ConflictingListObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:http_list:http_list", "type": "conflicts"}, {"anchor": "schema-rule_list--rules--http_list--http_list--suffix_value", "enforcement": "provider-schema", "group": "rule_list.rules.http_list.http_list:ConflictingListObjectAttributes:regex_value,suffix_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:http_list:http_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.http_list.http_list:ConflictingListObjectAttributes:any_path,path_exact_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:http_list:http_list:any_path", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.http_list.http_list:ConflictingListObjectAttributes:any_path,path_prefix_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:http_list:http_list:any_path", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.http_list.http_list:ConflictingListObjectAttributes:any_path,path_regex_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:http_list:http_list:any_path", "type": "conflicts"}], "schema_path": ["rule_list", "rules", "http_list", "http_list"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/forward_proxy_policy/properties/rule_list/rules/http_list/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "URLListType.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.http_list

Breadcrumbs:

- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/)
- rule_list.rules.http_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

URLListType.

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
http_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [http_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/http_list/http_list/): complete subsection reference.

## Next pages

- [rule_list.rules.http_list.http_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/http_list/http_list/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/)
- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/)
