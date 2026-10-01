---
page_title: "aws"
subcategory: ""
description: "aws for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 3162, "body_sha256": "sha256:3291f6a6987a84470e23e88099006e1e5dd9eb04e7e322167e6d36077f15f2d9", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:aws:not_managed"], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:aws", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "documentation/resources/securemesh_site_v2/properties/aws/index.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["aws"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/aws/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
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

- [aws](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/aws/#section)
- [azure](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/azure/#section)
- [baremetal](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/baremetal/#section)
- [eks_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/eks_k8s/#section)
- [equinix](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/equinix/#section)
- [gcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/gcp/#section)
- [kvm](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/kvm/#section)
- [nutanix](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/nutanix/#section)
- [oci](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/oci/#section)
- [openshift_virtualization](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openshift_virtualization/#section)
- [openstack](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openstack/#section)
- [vmware](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/vmware/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
aws {
  # Configure direct properties listed below.
}
```

## Direct properties

- [not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/aws/not_managed/): complete subsection reference.

## Next pages

- [aws.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/aws/not_managed/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
