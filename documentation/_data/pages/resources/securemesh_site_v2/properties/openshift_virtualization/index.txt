---
page_title: "openshift_virtualization"
subcategory: ""
description: "OpenShift Provider Type."
xcsh_docs: {"aliases": ["openshift virtualization"], "body_bytes": 1145, "body_sha256": "sha256:5da5ab9803e54b5e7e0e5050e0898e706218ae86e03df62916b1ddeb04dbc10d", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "documentation/resources/securemesh_site_v2/properties/openshift_virtualization/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-014.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["openshift_virtualization"], "schema_version": 1, "sections": [{"aliases": ["openshift virtualization not managed"], "anchor": "section", "description": "This section will show nodes associated with this site. Note: For sites that are not orchestrated by F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it will be shown in this section.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["openshift_virtualization", "not_managed"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/openshift_virtualization/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "OpenShift Provider Type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# openshift_virtualization

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- openshift_virtualization

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for openshift virtualization.

Additional upstream details:

OpenShift Provider Type.

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

Terraform syntax:

```terraform
openshift_virtualization {
  # Configure direct properties listed below.
}
```

## Direct properties

- [not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openshift_virtualization/not_managed/): complete subsection reference.
