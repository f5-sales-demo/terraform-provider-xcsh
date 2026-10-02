---
page_title: "voltstack_cluster.active_forward_proxy_policies"
subcategory: "Infrastructure"
description: "Ordered List of Forward Proxy Policies active."
xcsh_docs: {"aliases": ["voltstack cluster active forward proxy policies"], "body_bytes": 1871, "body_sha256": "sha256:46324a3e9fb2749111800cadb578295b84e00862489a45d3ec396ff4e500ced5", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster:active_forward_proxy_policies:forward_proxy_policies"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster:active_forward_proxy_policies", "parent_id": "xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster", "path": "documentation/resources/aws_vpc_site/properties/voltstack_cluster/active_forward_proxy_policies/index.md", "product": "distributed-cloud", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0220112232330111-1201221332013323-3130113332313130-3332103123231212-3301233011310123-0130201132101023-0000212222120332-3232313200012311", "registry_path": "docs/guides/resources--aws_vpc_site--reference--group-004.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster.active_forward_proxy_policies:RequiredObjectAttributes:forward_proxy_policies", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster:active_forward_proxy_policies:forward_proxy_policies", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["voltstack_cluster", "active_forward_proxy_policies"], "schema_version": 1, "sections": [{"aliases": ["forward proxy policies"], "anchor": "section", "description": "Ordered List of Forward Proxy Policies active.", "document_id": "xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster:active_forward_proxy_policies:forward_proxy_policies", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-voltstack_cluster--active_forward_proxy_policies--forward_proxy_policies--name", "enforcement": "provider-schema", "group": "voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster:active_forward_proxy_policies:forward_proxy_policies", "type": "requires"}], "schema_path": ["voltstack_cluster", "active_forward_proxy_policies", "forward_proxy_policies"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/voltstack_cluster/active_forward_proxy_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Ordered List of Forward Proxy Policies active.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster.active_forward_proxy_policies

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/voltstack_cluster/)
- voltstack_cluster.active_forward_proxy_policies

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("forward_proxy_policies")}
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
active_forward_proxy_policies {
  # Configure direct properties listed below.
}
```

## Direct properties

- [forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/voltstack_cluster/active_forward_proxy_policies/forward_proxy_policies/): complete subsection reference.

## Next pages

- [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/voltstack_cluster/active_forward_proxy_policies/forward_proxy_policies/)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/voltstack_cluster/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
