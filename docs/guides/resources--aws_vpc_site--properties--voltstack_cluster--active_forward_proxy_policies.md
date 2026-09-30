---
page_title: "voltstack_cluster.active_forward_proxy_policies"
subcategory: "Infrastructure"
description: "voltstack_cluster.active_forward_proxy_policies for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1417, "body_sha256": "sha256:f1f2db669f47f87cf51559ada60ee555d600389fb463c7d620d97277bfdd697e", "canonical_id": "xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster:active_forward_proxy_policies", "child_ids": ["xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster:active_forward_proxy_policies:forward_proxy_policies"], "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster:active_forward_proxy_policies", "parent_id": "xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster", "path": "docs/guides/resources--aws_vpc_site--properties--voltstack_cluster--active_forward_proxy_policies.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["voltstack_cluster", "active_forward_proxy_policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/voltstack_cluster/active_forward_proxy_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster.active_forward_proxy_policies for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# voltstack_cluster.active_forward_proxy_policies

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
- [Property reference](resources--aws_vpc_site--reference.md)
- [voltstack_cluster](resources--aws_vpc_site--properties--voltstack_cluster.md)
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

- [forward_proxy_policies](resources--aws_vpc_site--properties--voltstack_cluster--active_forward_proxy_policies--forward_proxy_policies.md): complete subsection reference.

## Next pages

- [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies](resources--aws_vpc_site--properties--voltstack_cluster--active_forward_proxy_policies--forward_proxy_policies.md)
- [voltstack_cluster](resources--aws_vpc_site--properties--voltstack_cluster.md)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
