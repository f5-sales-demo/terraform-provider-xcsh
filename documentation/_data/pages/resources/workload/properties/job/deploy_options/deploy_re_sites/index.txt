---
page_title: "job.deploy_options.deploy_re_sites"
subcategory: "Container"
description: "This defines a way to deploy a workload on specific Regional Edge sites."
xcsh_docs: {"aliases": ["job deploy options deploy re sites"], "body_bytes": 1888, "body_sha256": "sha256:efe77f671ecd2266e16c621383b97ba9a99f185603de16ba9bba0ac9e02287f3", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:job:deploy_options:deploy_re_sites:site"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_re_sites", "parent_id": "xcsh-docs:resources:workload:properties:job:deploy_options", "path": "documentation/resources/workload/properties/job/deploy_options/deploy_re_sites/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3202012002300303-1220012332222101-0310120133000321-0300120210310013-2101111121220301-0321202013303020-1302001321033011-1133320020002000", "registry_path": "docs/guides/resources--workload--reference--group-004.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "job.deploy_options.deploy_re_sites:RequiredObjectAttributes:site", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_re_sites:site", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["job", "deploy_options", "deploy_re_sites"], "schema_version": 1, "sections": [{"aliases": ["site"], "anchor": "section", "description": "Which regional edge sites should this workload be deployed.", "document_id": "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_re_sites:site", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-job--deploy_options--deploy_re_sites--site--name", "enforcement": "provider-schema", "group": "job.deploy_options.deploy_re_sites.site:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_re_sites:site", "type": "requires"}], "schema_path": ["job", "deploy_options", "deploy_re_sites", "site"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/job/deploy_options/deploy_re_sites/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This defines a way to deploy a workload on specific Regional Edge sites.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job.deploy_options.deploy_re_sites

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [job](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/)
- [job.deploy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/deploy_options/)
- job.deploy_options.deploy_re_sites

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to deploy a workload on specific Regional Edge sites.

Upstream description:

This defines a way to deploy a workload on specific Regional Edge sites.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("site")}
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
deploy_re_sites {
  # Configure direct properties listed below.
}
```

## Direct properties

- [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/deploy_options/deploy_re_sites/site/): complete subsection reference.

## Next pages

- [job.deploy_options.deploy_re_sites.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/deploy_options/deploy_re_sites/site/)
- [job.deploy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/deploy_options/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
