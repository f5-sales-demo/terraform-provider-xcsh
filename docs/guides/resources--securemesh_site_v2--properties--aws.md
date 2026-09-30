---
page_title: "aws"
subcategory: ""
description: "aws for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 2143, "body_sha256": "sha256:5f0c4487bf29721b67190def970b4b9a95871656dce1f3a2a252bd72b7621dd1", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:aws", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:aws:not_managed"], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:aws", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "docs/guides/resources--securemesh_site_v2--properties--aws.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["aws"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/aws/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# aws

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- aws

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

- [aws](resources--securemesh_site_v2--properties--aws.md#section)
- [azure](resources--securemesh_site_v2--properties--azure.md#section)
- [baremetal](resources--securemesh_site_v2--properties--baremetal.md#section)
- [eks_k8s](resources--securemesh_site_v2--properties--eks_k8s.md#section)
- [equinix](resources--securemesh_site_v2--properties--equinix.md#section)
- [gcp](resources--securemesh_site_v2--properties--gcp.md#section)
- [kvm](resources--securemesh_site_v2--properties--kvm.md#section)
- [nutanix](resources--securemesh_site_v2--properties--nutanix.md#section)
- [oci](resources--securemesh_site_v2--properties--oci.md#section)
- [openshift_virtualization](resources--securemesh_site_v2--properties--openshift_virtualization.md#section)
- [openstack](resources--securemesh_site_v2--properties--openstack.md#section)
- [vmware](resources--securemesh_site_v2--properties--vmware.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
aws {
  # Configure direct properties listed below.
}
```

## Direct properties

- [not_managed](resources--securemesh_site_v2--properties--aws--not_managed.md): complete subsection reference.

## Next pages

- [aws.not_managed](resources--securemesh_site_v2--properties--aws--not_managed.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
