---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_partner_customer_support_comments."
xcsh_docs: {"aliases": ["partner customer support comments"], "body_bytes": 4613, "body_sha256": "sha256:735c2027796cda921ee70efc0e55c8ceb7162f6fac560befb367c471c16f4374", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:partner_customer_support_comments:properties:comments"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:partner_customer_support_comments:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:partner_customer_support_comments:reference", "parent_id": "xcsh-docs:data-sources:partner_customer_support_comments:fundamentals", "path": "documentation/data-sources/partner_customer_support_comments/properties/index.md", "product": "distributed-cloud", "provider_name": "partner_customer_support_comments", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-0233221211110202-0203013030322003-2221102332003211-2100233032223323-1132312111031303-3332333322001111-1213311201313222-2030303112323312", "registry_path": "docs/guides/data-sources--partner_customer_support_comments--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["all comments returned"], "anchor": "schema-all_comments_returned", "description": "Indicates if all comments for the issue have been returned.", "document_id": "xcsh-docs:data-sources:partner_customer_support_comments:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["all_comments_returned"], "syntax": "attribute", "type": "bool"}, {"aliases": ["comments"], "anchor": "section", "description": "List of comments on the customer support ticket.", "document_id": "xcsh-docs:data-sources:partner_customer_support_comments:properties:comments", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["comments"], "syntax": "attribute", "type": "object"}, {"aliases": ["created until timestamp"], "anchor": "schema-created_until_timestamp", "description": "Filter to retrieve comments created up to the specified timestamp.", "document_id": "xcsh-docs:data-sources:partner_customer_support_comments:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["created_until_timestamp"], "syntax": "attribute", "type": "string"}, {"aliases": ["tp id"], "anchor": "schema-tp_id", "description": "Tpid ID assigned to this ticket by Third Party.", "document_id": "xcsh-docs:data-sources:partner_customer_support_comments:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tp_id"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/partner_customer_support_comments/properties/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Property reference for xcsh_partner_customer_support_comments.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": [], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_partner_customer_support_comments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/partner_customer_support_comments/)
- Property reference

## Direct properties

<a id="schema-all_comments_returned"></a>

### all_comments_returned property

Type: `"bool"`. Computed.

Indicates if all comments for the issue have been returned.

- [comments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/partner_customer_support_comments/properties/comments/): complete subsection reference.

<a id="schema-created_until_timestamp"></a>

### created_until_timestamp property

Type: `"string"`. Optional.

Filter to retrieve comments created up to the specified timestamp.

<a id="schema-tp_id"></a>

### tp_id property

Type: `"string"`. Required.

Tpid ID assigned to this ticket by Third Party.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `all_comments_returned` | [all_comments_returned](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/partner_customer_support_comments/properties/#schema-all_comments_returned) |
| `comments` | [comments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/partner_customer_support_comments/properties/comments/#section) |
| `comments.attachment_ids` | [comments.attachment_ids](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/partner_customer_support_comments/properties/comments/#schema-comments--attachment_ids) |
| `comments.attachments_info` | [comments.attachments_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/partner_customer_support_comments/properties/comments/attachments_info/#section) |
| `comments.attachments_info.attachment` | [comments.attachments_info.attachment](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/partner_customer_support_comments/properties/comments/attachments_info/#schema-comments--attachments_info--attachment) |
| `comments.attachments_info.content_type` | [comments.attachments_info.content_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/partner_customer_support_comments/properties/comments/attachments_info/#schema-comments--attachments_info--content_type) |
| `comments.attachments_info.filename` | [comments.attachments_info.filename](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/partner_customer_support_comments/properties/comments/attachments_info/#schema-comments--attachments_info--filename) |
| `comments.attachments_info.tp_id` | [comments.attachments_info.tp_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/partner_customer_support_comments/properties/comments/attachments_info/#schema-comments--attachments_info--tp_id) |
| `comments.author_email` | [comments.author_email](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/partner_customer_support_comments/properties/comments/#schema-comments--author_email) |
| `comments.author_name` | [comments.author_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/partner_customer_support_comments/properties/comments/#schema-comments--author_name) |
| `comments.comment_id` | [comments.comment_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/partner_customer_support_comments/properties/comments/#schema-comments--comment_id) |
| `comments.created_at` | [comments.created_at](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/partner_customer_support_comments/properties/comments/#schema-comments--created_at) |
| `comments.html` | [comments.html](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/partner_customer_support_comments/properties/comments/#schema-comments--html) |
| `comments.plain_text` | [comments.plain_text](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/partner_customer_support_comments/properties/comments/#schema-comments--plain_text) |
| `created_until_timestamp` | [created_until_timestamp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/partner_customer_support_comments/properties/#schema-created_until_timestamp) |
| `tp_id` | [tp_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/partner_customer_support_comments/properties/#schema-tp_id) |
