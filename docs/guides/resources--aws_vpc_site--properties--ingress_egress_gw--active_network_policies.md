---
page_title: "ingress_egress_gw.active_network_policies"
subcategory: "Infrastructure"
description: "ingress_egress_gw.active_network_policies for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1412, "body_sha256": "sha256:cbd1125347d0016db3a2a9c39116690065a3c0126fb0ab9ae6d7cb80a85ac444", "canonical_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:active_network_policies", "child_ids": ["xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:active_network_policies:network_policies"], "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:active_network_policies", "parent_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw", "path": "docs/guides/resources--aws_vpc_site--properties--ingress_egress_gw--active_network_policies.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw", "active_network_policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/ingress_egress_gw/active_network_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw.active_network_policies for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ingress_egress_gw.active_network_policies

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
- [Property reference](resources--aws_vpc_site--reference.md)
- [ingress_egress_gw](resources--aws_vpc_site--properties--ingress_egress_gw.md)
- ingress_egress_gw.active_network_policies

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for active network policies.

Upstream description:

List of firewall policy views.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("network_policies")}
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
active_network_policies {
  # Configure direct properties listed below.
}
```

## Direct properties

- [network_policies](resources--aws_vpc_site--properties--ingress_egress_gw--active_network_policies--network_policies.md): complete subsection reference.

## Next pages

- [ingress_egress_gw.active_network_policies.network_policies](resources--aws_vpc_site--properties--ingress_egress_gw--active_network_policies--network_policies.md)
- [ingress_egress_gw](resources--aws_vpc_site--properties--ingress_egress_gw.md)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
