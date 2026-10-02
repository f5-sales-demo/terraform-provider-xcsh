---
page_title: "aws"
subcategory: ""
description: "AWS Provider Type."
xcsh_docs: {"aliases": ["aws"], "body_bytes": 3162, "body_sha256": "sha256:3291f6a6987a84470e23e88099006e1e5dd9eb04e7e322167e6d36077f15f2d9", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:aws:not_managed"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:aws", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "documentation/resources/securemesh_site_v2/properties/aws/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["aws"], "schema_version": 1, "sections": [{"aliases": ["not managed"], "anchor": "section", "description": "This section will show nodes associated with this site. Note: For sites that are not orchestrated by F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it will be shown in this section.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:aws:not_managed", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["aws", "not_managed"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/aws/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "AWS Provider Type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
