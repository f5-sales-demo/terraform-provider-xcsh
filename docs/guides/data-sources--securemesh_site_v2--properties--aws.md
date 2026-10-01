---
page_title: "aws"
subcategory: ""
description: "aws for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 2184, "body_sha256": "sha256:2ce90917a357f8f9acc3ea38e2f973850bf4fbc410329851faca1ef79740e0ab", "canonical_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:aws", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:aws:not_managed"], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:aws", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:reference", "path": "docs/guides/data-sources--securemesh_site_v2--properties--aws.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["aws"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/aws/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- aws

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: aws, azure, baremetal, eks\_k8s, equinix, gcp, kvm, nutanix, oci,
openshift\_virtualization, openstack, vmware\] AWS Provider Type. AWS Provider Type.

Upstream description:

AWS Provider Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-orchestration_choice": "[\"not_managed\"]"
}
```

OneOf alternatives in this subsection:

- [aws](data-sources--securemesh_site_v2--properties--aws.md#section)
- [azure](data-sources--securemesh_site_v2--properties--azure.md#section)
- [baremetal](data-sources--securemesh_site_v2--properties--baremetal.md#section)
- [eks_k8s](data-sources--securemesh_site_v2--properties--eks_k8s.md#section)
- [equinix](data-sources--securemesh_site_v2--properties--equinix.md#section)
- [gcp](data-sources--securemesh_site_v2--properties--gcp.md#section)
- [kvm](data-sources--securemesh_site_v2--properties--kvm.md#section)
- [nutanix](data-sources--securemesh_site_v2--properties--nutanix.md#section)
- [oci](data-sources--securemesh_site_v2--properties--oci.md#section)
- [openshift_virtualization](data-sources--securemesh_site_v2--properties--openshift_virtualization.md#section)
- [openstack](data-sources--securemesh_site_v2--properties--openstack.md#section)
- [vmware](data-sources--securemesh_site_v2--properties--vmware.md#section)

Select alternatives according to the provider validators above.

## Direct properties

- [not_managed](data-sources--securemesh_site_v2--properties--aws--not_managed.md): complete subsection reference.

## Next pages

- [aws.not_managed](data-sources--securemesh_site_v2--properties--aws--not_managed.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
