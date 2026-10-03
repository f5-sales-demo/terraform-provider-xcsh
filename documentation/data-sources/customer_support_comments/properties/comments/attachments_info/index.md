---
page_title: "comments.attachments_info"
subcategory: ""
description: "Information about any attachments (such as screenshots, plain text files) the comment can have."
xcsh_docs: {"aliases": ["comments attachments info"], "body_bytes": 1768, "body_sha256": "sha256:9e8a5e63fda1d29d0a2bbbc0f6ccecd47f6a3c046ea8160cb244948c2d9f52e1", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:customer_support_comments:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:customer_support_comments:properties:comments:attachments_info", "parent_id": "xcsh-docs:data-sources:customer_support_comments:properties:comments", "path": "documentation/data-sources/customer_support_comments/properties/comments/attachments_info/index.md", "product": "distributed-cloud", "provider_name": "customer_support_comments", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1123201102321012-0010233330122000-2321201222222020-2123030310113230-2200012331021303-2311011102210123-3300232002210220-0110001102200201", "registry_path": "docs/guides/data-sources--customer_support_comments--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["comments", "attachments_info"], "schema_version": 1, "sections": [{"aliases": ["comments attachments info attachment"], "anchor": "schema-comments--attachments_info--attachment", "description": "Any binary attachment (such as screenshots, plain text files, PDFs) encoded as base64 if used over HTTP.", "document_id": "xcsh-docs:data-sources:customer_support_comments:properties:comments:attachments_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["comments", "attachments_info", "attachment"], "syntax": "attribute", "type": "string"}, {"aliases": ["comments attachments info content type"], "anchor": "schema-comments--attachments_info--content_type", "description": "MIME content type of the attachment. Helps the UI to properly display the data.", "document_id": "xcsh-docs:data-sources:customer_support_comments:properties:comments:attachments_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["comments", "attachments_info", "content_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["comments attachments info filename"], "anchor": "schema-comments--attachments_info--filename", "description": "Filename of the attachment as provided by the caller.", "document_id": "xcsh-docs:data-sources:customer_support_comments:properties:comments:attachments_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["comments", "attachments_info", "filename"], "syntax": "attribute", "type": "string"}, {"aliases": ["comments attachments info tp id"], "anchor": "schema-comments--attachments_info--tp_id", "description": "Optional ID as assigned by the third-party actually storing the data.", "document_id": "xcsh-docs:data-sources:customer_support_comments:properties:comments:attachments_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["comments", "attachments_info", "tp_id"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/customer_support_comments/properties/comments/attachments_info/index.txt", "spec_pin_digest": "sha256:62f71ec22260bc99f65753ef4581eb9e0dec1b65c506bb0d53099db73a05e19e", "summary": "Information about any attachments (such as screenshots, plain text files) the comment can have.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.2", "schema_components": [], "target_commit": "1a0b5141f4589ffaf7bb696a4369a16ee74ae2ff"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# comments.attachments_info

Breadcrumbs:

- [xcsh_customer_support_comments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/customer_support_comments/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/customer_support_comments/properties/)
- [comments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/customer_support_comments/properties/comments/)
- comments.attachments_info

<a id="section"></a>

Type: `"list"`. Computed.

Information about any attachments (such as screenshots, plain text files) the comment can have.

## Direct properties

<a id="schema-comments--attachments_info--attachment"></a>

### attachment property

Type: `"string"`. Computed.

Any binary attachment (such as screenshots, plain text files, PDFs) encoded as base64 if used over
HTTP.

<a id="schema-comments--attachments_info--content_type"></a>

### content_type property

Type: `"string"`. Computed.

MIME content type of the attachment. Helps the UI to properly display the data.

<a id="schema-comments--attachments_info--filename"></a>

### filename property

Type: `"string"`. Computed.

Filename of the attachment as provided by the caller.

<a id="schema-comments--attachments_info--tp_id"></a>

### tp_id property

Type: `"string"`. Computed.

Optional ID as assigned by the third-party actually storing the data.

## Next pages

- [comments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/customer_support_comments/properties/comments/)
- [xcsh_customer_support_comments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/customer_support_comments/)
