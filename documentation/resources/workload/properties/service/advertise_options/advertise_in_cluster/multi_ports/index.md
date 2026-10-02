---
page_title: "service.advertise_options.advertise_in_cluster.multi_ports"
subcategory: "Container"
description: "Multiple ports."
xcsh_docs: {"aliases": ["service advertise options advertise in cluster multi ports"], "body_bytes": 2180, "body_sha256": "sha256:a46644a646c76db782c689ed39295f1b193311798dd2595e50e77f1295bb9e39", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:service:advertise_options:advertise_in_cluster:multi_ports:ports"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_in_cluster:multi_ports", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_in_cluster", "path": "documentation/resources/workload/properties/service/advertise_options/advertise_in_cluster/multi_ports/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-3233102030323303-0010001111332222-2221330300220032-2322233002300322-1030033321333231-1200300001103302-1331010302322302-2231113103111232", "registry_path": "docs/guides/resources--workload--reference--group-008.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_in_cluster.multi_ports:RequiredObjectAttributes:ports", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_in_cluster:multi_ports:ports", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "advertise_options", "advertise_in_cluster", "multi_ports"], "schema_version": 1, "sections": [{"aliases": ["ports"], "anchor": "section", "description": "Ports to advertise.", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_in_cluster:multi_ports:ports", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-service--advertise_options--advertise_in_cluster--multi_ports--ports--name", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_in_cluster.multi_ports.ports:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_in_cluster:multi_ports:ports", "type": "requires"}], "schema_path": ["service", "advertise_options", "advertise_in_cluster", "multi_ports", "ports"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_in_cluster/multi_ports/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Multiple ports.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_in_cluster.multi_ports

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/)
- [service.advertise_options.advertise_in_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_in_cluster/)
- service.advertise_options.advertise_in_cluster.multi_ports

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Multiple Ports. Multiple ports.

Upstream description:

Multiple ports.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ports")}
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
multi_ports {
  # Configure direct properties listed below.
}
```

## Direct properties

- [ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_in_cluster/multi_ports/ports/): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_in_cluster.multi_ports.ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_in_cluster/multi_ports/ports/)
- [service.advertise_options.advertise_in_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_in_cluster/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
