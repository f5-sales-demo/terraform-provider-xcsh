---
page_title: "voltstack_cluster.active_network_policies"
subcategory: "Infrastructure"
description: "voltstack_cluster.active_network_policies for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1146, "body_sha256": "sha256:2be60ffdc0db2ce79ff97103885a8a213bcc2b17f28920be2f4af044d8c4d86a", "canonical_id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:active_network_policies", "child_ids": ["xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:active_network_policies:network_policies"], "collection_id": "xcsh-docs:data-sources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:active_network_policies", "parent_id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster", "path": "docs/guides/data-sources--aws_vpc_site--properties--voltstack_cluster--active_network_policies.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["voltstack_cluster", "active_network_policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_vpc_site/properties/voltstack_cluster/active_network_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster.active_network_policies for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# voltstack_cluster.active_network_policies

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md)
- [Property reference](data-sources--aws_vpc_site--reference.md)
- [voltstack_cluster](data-sources--aws_vpc_site--properties--voltstack_cluster.md)
- voltstack_cluster.active_network_policies

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for active network policies.

Upstream description:

List of firewall policy views.

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

## Direct properties

- [network_policies](data-sources--aws_vpc_site--properties--voltstack_cluster--active_network_policies--network_policies.md): complete subsection reference.

## Next pages

- [voltstack_cluster.active_network_policies.network_policies](data-sources--aws_vpc_site--properties--voltstack_cluster--active_network_policies--network_policies.md)
- [voltstack_cluster](data-sources--aws_vpc_site--properties--voltstack_cluster.md)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md)
