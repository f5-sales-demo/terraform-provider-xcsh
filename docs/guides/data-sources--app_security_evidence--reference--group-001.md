---
page_title: "xcsh_app_security_evidence reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_security_evidence reference."
---

# xcsh_app_security_evidence reference

<a id="canonical-160f5b6ebef4ab35b1510dc3be35d8747129089705ae396e4d0fa10d3481c008"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ed94591335db757325afb2c243c22c395bef96067f2a6b0a5937642a60ac7c41"></a>

## Property reference — Property reference / 43d9d47f6963 / 2

Breadcrumbs:

- [xcsh_app_security_evidence](../data-sources/app_security_evidence.md#canonical-d3c80b619742ede9129532b0a717cd17f1a275a4761a59fe318edc65c1f1d29b)
- Property reference

<a id="canonical-21f12f67771b8438f2d452fab9f47d500c9933b9166d4798f249f4a86e638ef1"></a>

## Direct properties — Property reference / 43d9d47f6963 / 3

- [aggs](data-sources--app_security_evidence--reference--group-001.md#canonical-b2bb1e6ea5b5f2e31b06ab04584089531be04aa30eafc40c07f4055e2fd5c152): complete subsection reference.

<a id="canonical-f7045f6c9fcb4f7df7abb15d0bfd391e1126faa6c244ebb6582640f746074bd1"></a>

<a id="canonical-0df31544e32217989c7ce5490a49f99774ae062f77d05834899e43df7bcf034b"></a>

## end_time property — Property reference / 43d9d47f6963 / 4

Type: `"string"`. Optional.

Fetch security evidence whose timestamp &lt;= end\_time format: unix\_timestamp|RFC 3339 Optional:
If not specified, then the end\_time will be evaluated to start\_time+10m If start\_time is not
specified, then the end\_time will be evaluated to &lt;current time&gt;.

<a id="canonical-e39b97d021e8fa6a363516f2cc87a93d0dc727c57a51d9816234ff45483e4435"></a>

<a id="canonical-045ae25527a9a6e556960e6660c6c0a959018537d2d7735cb97fbdc52ade5647"></a>

## evidences property — Property reference / 43d9d47f6963 / 5

Type: `["list", "string"]`. Computed.

List of security evidences that matched the query. Contains no more than 500 messages.

- [last_sort_values](data-sources--app_security_evidence--reference--group-001.md#canonical-c464f9b61ce5454934ee254847fdf71eb652ca1c54bd81bfe8f225adf2cb2cfa): complete subsection reference.

<a id="canonical-1f96f931f5552f76e3a599daf6fc59760888ddfdf876be80ce1981c7972384f0"></a>

<a id="canonical-bb7ab803460f966c3bba12555baa2275b9684adfa9c16fee215cf515596857c1"></a>

## limit property — Property reference / 43d9d47f6963 / 6

Type: `"number"`. Optional.

Limits the number of security evidence returned in the response Optional: If not specified, first or
last 500 security evidence that matches the query (depending on the sort order) will be returned in
the response. The maximum value for limit is 500.

<a id="canonical-a88adb4d0ab944d5dbf8e375aa65ce9ffab47d2d28f34caf188a2e7d94268de0"></a>

<a id="canonical-92d4e3890ee06de44a4200f3bef639bcdf455b4be1d99b43c3b1f98c150ee3a0"></a>

## namespace property — Property reference / 43d9d47f6963 / 7

Type: `"string"`. Required.

Namespace fetch security evidence for a given namespace.

<a id="canonical-cf62373f9154b14e616b1ae05dc4dd8a8f16eea2856cd3a39cac24097dbf8519"></a>

<a id="canonical-df6d0ebd4f40192aceadc98cef087b02e5bc1385fe7d3f9a8c7e828a40bd8efc"></a>

## query property — Property reference / 43d9d47f6963 / 8

Type: `"string"`. Optional.

Query is used to specify the list of matchers syntax for query := \{\[&lt;matcher&gt;\]\}
&lt;matcher&gt; := &lt;field\_name&gt;&lt;operator&gt;'&lt;value&gt;' &lt;field\_name&gt; := string
One or more of these fields in the security evidence may be specified in the query. Domain - domain
endpoint - endpoint evidence\_id - evidence ID..

<a id="canonical-3a624ff958cb0d00c0d759cd19bb977530ab2cc7a3b53e6afc0415749e6280a0"></a>

<a id="canonical-90b1fe5f0ef9038a796067e0f1a98f027e4a6e1cbbd170c6d1e51f366618a374"></a>

## search_after property — Property reference / 43d9d47f6963 / 9

Type: `"bool"`. Optional.

Search After is used to retrieve large number of log messages (or all log messages) that matches the
query. If search\_after is set to true, the sort\_values in the response can be used in the API to
fetch the next batch of logs. The number of messages in each batch is determined by the limit field.

<a id="canonical-62ea08c0dd2301e0a0030304aa69af60f9578f8758d5de2b878db8e5502e65d7"></a>

<a id="canonical-4713e69a5ee2cd35ca98a5010c91ddda2ef685a551c0e2a7b1c0081567773c08"></a>

## sort property — Property reference / 43d9d47f6963 / 10

Type: `"string"`. Optional.

\[Enum: DESCENDING|ASCENDING\] Sort algorithm Sort in descending order Sort in ascending order.
Possible values are \`DESCENDING\`, \`ASCENDING\`. Defaults to \`DESCENDING\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DESCENDING",
    "ASCENDING"),
}
```

<a id="canonical-fe9da1f5906c25f33021879a83770f34903c5989607ef0bcae0b296ba91946eb"></a>

<a id="canonical-aab7335e440299870dbc75218c06fe8cf7edffca2ad07781bf629962006abaa6"></a>

## sort_by property — Property reference / 43d9d47f6963 / 11

Type: `"string"`. Optional.

Optional: default is sort by last\_event\_time.

- [sort_values](data-sources--app_security_evidence--reference--group-001.md#canonical-f6d83f3f31a4204a3c8cdf851a03387dc2801f62c7f874ecc4e280f3d1434a57): complete subsection reference.

<a id="canonical-bbb527d4688ba1d67fce8ea9acb4a64fd5401234a09ff7d08f311a45ce8c7dd4"></a>

<a id="canonical-80e5eb2171931899f774625698730db1157ca4546bb277d1116a836e00394b35"></a>

## start_time property — Property reference / 43d9d47f6963 / 12

Type: `"string"`. Optional.

Fetch security evidence whose timestamp &gt;= start\_time format: unix\_timestamp|RFC 3339 Optional:
If not specified, then the start\_time will be evaluated to end\_time-10m If end\_time is not
specified, then the start\_time will be evaluated to &lt;current time&gt;-10m.

<a id="canonical-1d82df4663f567e2fa5346346c69644ed00bf9e489dbd7b903b658dd9b222048"></a>

<a id="canonical-460abdf81ab8ebdaac355b235a762585697b6a496828242db9c00c4831d64c50"></a>

## total_hits property — Property reference / 43d9d47f6963 / 13

Type: `"string"`. Computed.

Total number of security events that matched the query.

<a id="canonical-9baafa8be6900d7a13e631ea6764dd18c66ea1755ca0a184fdfde9e4161f2be9"></a>

## All schema paths — Property reference / 43d9d47f6963 / 14

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `aggs` | [aggs](data-sources--app_security_evidence--reference--group-001.md#canonical-ad66de7dff3bb7c6af44eb0c8c2b3a8d4996dd7ff7ed9d61e8f264bdb52631f3) |
| `end_time` | [end_time](data-sources--app_security_evidence--reference--group-001.md#canonical-f7045f6c9fcb4f7df7abb15d0bfd391e1126faa6c244ebb6582640f746074bd1) |
| `evidences` | [evidences](data-sources--app_security_evidence--reference--group-001.md#canonical-e39b97d021e8fa6a363516f2cc87a93d0dc727c57a51d9816234ff45483e4435) |
| `last_sort_values` | [last_sort_values](data-sources--app_security_evidence--reference--group-001.md#canonical-02c14f2aebbd6c7d83b7793c0fde2d5a3f822ce3e958a9d7babeade009cc852f) |
| `last_sort_values.last_doc_id` | [last_sort_values.last_doc_id](data-sources--app_security_evidence--reference--group-001.md#canonical-d977d5125f29431bc106e73205e450058c5ce1d7b1002369d7c8b52467c928af) |
| `last_sort_values.last_timestamp` | [last_sort_values.last_timestamp](data-sources--app_security_evidence--reference--group-001.md#canonical-c5117b77de12e883891e69c613dbabebb4fe0b07523707a1e152f741736518ed) |
| `limit` | [limit](data-sources--app_security_evidence--reference--group-001.md#canonical-1f96f931f5552f76e3a599daf6fc59760888ddfdf876be80ce1981c7972384f0) |
| `namespace` | [namespace](data-sources--app_security_evidence--reference--group-001.md#canonical-a88adb4d0ab944d5dbf8e375aa65ce9ffab47d2d28f34caf188a2e7d94268de0) |
| `query` | [query](data-sources--app_security_evidence--reference--group-001.md#canonical-cf62373f9154b14e616b1ae05dc4dd8a8f16eea2856cd3a39cac24097dbf8519) |
| `search_after` | [search_after](data-sources--app_security_evidence--reference--group-001.md#canonical-3a624ff958cb0d00c0d759cd19bb977530ab2cc7a3b53e6afc0415749e6280a0) |
| `sort` | [sort](data-sources--app_security_evidence--reference--group-001.md#canonical-62ea08c0dd2301e0a0030304aa69af60f9578f8758d5de2b878db8e5502e65d7) |
| `sort_by` | [sort_by](data-sources--app_security_evidence--reference--group-001.md#canonical-fe9da1f5906c25f33021879a83770f34903c5989607ef0bcae0b296ba91946eb) |
| `sort_values` | [sort_values](data-sources--app_security_evidence--reference--group-001.md#canonical-4d7c6d537a402b40613089630a072b2f0a898153717df66dbc84d3ef4b55081a) |
| `sort_values.last_doc_id` | [sort_values.last_doc_id](data-sources--app_security_evidence--reference--group-001.md#canonical-bce97ea847c1166f77b64caf3f95b1504567837567c30f06e4428594f29d578a) |
| `sort_values.last_timestamp` | [sort_values.last_timestamp](data-sources--app_security_evidence--reference--group-001.md#canonical-3648713bb0a27f87e137b8d7337a791170cc20f29cc2e0f121af446580de9afa) |
| `start_time` | [start_time](data-sources--app_security_evidence--reference--group-001.md#canonical-bbb527d4688ba1d67fce8ea9acb4a64fd5401234a09ff7d08f311a45ce8c7dd4) |
| `total_hits` | [total_hits](data-sources--app_security_evidence--reference--group-001.md#canonical-1d82df4663f567e2fa5346346c69644ed00bf9e489dbd7b903b658dd9b222048) |

<a id="canonical-4ac3d8678b43bad9c808486f1d82d8bb7be02513a1aede7c215fb0d4af1cf534"></a>

## Next pages — Property reference / 43d9d47f6963 / 15

- [aggs](data-sources--app_security_evidence--reference--group-001.md#canonical-b2bb1e6ea5b5f2e31b06ab04584089531be04aa30eafc40c07f4055e2fd5c152)
- [last_sort_values](data-sources--app_security_evidence--reference--group-001.md#canonical-c464f9b61ce5454934ee254847fdf71eb652ca1c54bd81bfe8f225adf2cb2cfa)
- [sort_values](data-sources--app_security_evidence--reference--group-001.md#canonical-f6d83f3f31a4204a3c8cdf851a03387dc2801f62c7f874ecc4e280f3d1434a57)
- [xcsh_app_security_evidence](../data-sources/app_security_evidence.md#canonical-d3c80b619742ede9129532b0a717cd17f1a275a4761a59fe318edc65c1f1d29b)

<a id="canonical-b2bb1e6ea5b5f2e31b06ab04584089531be04aa30eafc40c07f4055e2fd5c152"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4df2d8986460b5f4b51a6983c7490b3fad172f24ffb2485c6a4fcbbb5a232d42"></a>

## aggs — aggs / f0a3b5687c3b / 2

Breadcrumbs:

- [xcsh_app_security_evidence](../data-sources/app_security_evidence.md#canonical-d3c80b619742ede9129532b0a717cd17f1a275a4761a59fe318edc65c1f1d29b)
- [Property reference](data-sources--app_security_evidence--reference--group-001.md#canonical-160f5b6ebef4ab35b1510dc3be35d8747129089705ae396e4d0fa10d3481c008)
- aggs

<a id="canonical-ad66de7dff3bb7c6af44eb0c8c2b3a8d4996dd7ff7ed9d61e8f264bdb52631f3"></a>

Type: `"single"`. Optional.

Aggregations provide summary/analytics data over the security evidence response. If the number of
security evidence that matched the query is large and cannot be returned in a single response
message, user can GET helpful insights/summary using aggregations. The aggregations are key'ed by..

<a id="canonical-92d33880244b319b4135221620f607cfab48db544b37e622f6b038a1ddd21bc9"></a>

## Direct properties — aggs / f0a3b5687c3b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f47f8914ac57e32eac2de7d084c7ae9607712cc0babb35337ff887e21600a2c2"></a>

## Next pages — aggs / f0a3b5687c3b / 4

- [Property reference](data-sources--app_security_evidence--reference--group-001.md#canonical-160f5b6ebef4ab35b1510dc3be35d8747129089705ae396e4d0fa10d3481c008)
- [xcsh_app_security_evidence](../data-sources/app_security_evidence.md#canonical-d3c80b619742ede9129532b0a717cd17f1a275a4761a59fe318edc65c1f1d29b)

<a id="canonical-c464f9b61ce5454934ee254847fdf71eb652ca1c54bd81bfe8f225adf2cb2cfa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c617b66199c0e4ae776c082fe8e61d9366151b48bae9770c00a84dbcc4eb72d6"></a>

## last_sort_values — last_sort_values / 5f485ac26499 / 2

Breadcrumbs:

- [xcsh_app_security_evidence](../data-sources/app_security_evidence.md#canonical-d3c80b619742ede9129532b0a717cd17f1a275a4761a59fe318edc65c1f1d29b)
- [Property reference](data-sources--app_security_evidence--reference--group-001.md#canonical-160f5b6ebef4ab35b1510dc3be35d8747129089705ae396e4d0fa10d3481c008)
- last_sort_values

<a id="canonical-02c14f2aebbd6c7d83b7793c0fde2d5a3f822ce3e958a9d7babeade009cc852f"></a>

Type: `"single"`. Computed.

These are timestamp and doc\_id values returned by elastic search in the search request. Client is
expected to set these values in a subsequent request to GET the next page of results.

<a id="canonical-42177fd1c05c456037efabebfe06ff7ae0fe5924b77a4d55461d2687836ffc6a"></a>

## Direct properties — last_sort_values / 5f485ac26499 / 3

<a id="canonical-d977d5125f29431bc106e73205e450058c5ce1d7b1002369d7c8b52467c928af"></a>

<a id="canonical-fc667cbc6f8991e99e0980599384323a8a71576f7df23e85190edcd7901b78ab"></a>

## last_doc_id property — last_sort_values / 5f485ac26499 / 4

Type: `"string"`. Computed.

Unique UUID generated by elastic search.

<a id="canonical-c5117b77de12e883891e69c613dbabebb4fe0b07523707a1e152f741736518ed"></a>

<a id="canonical-24ffafc95ccead075b5c12d47ea357a6353be221c137e511a77ee92a62025486"></a>

## last_timestamp property — last_sort_values / 5f485ac26499 / 5

Type: `"number"`. Computed.

Configuration parameter for last timestamp.

<a id="canonical-1fa92c968573e48d6dfd97b436be86b41b14df31d4698d4e8355c406f049c46e"></a>

## Next pages — last_sort_values / 5f485ac26499 / 6

- [Property reference](data-sources--app_security_evidence--reference--group-001.md#canonical-160f5b6ebef4ab35b1510dc3be35d8747129089705ae396e4d0fa10d3481c008)
- [xcsh_app_security_evidence](../data-sources/app_security_evidence.md#canonical-d3c80b619742ede9129532b0a717cd17f1a275a4761a59fe318edc65c1f1d29b)

<a id="canonical-f6d83f3f31a4204a3c8cdf851a03387dc2801f62c7f874ecc4e280f3d1434a57"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2ef6183b3bd8730b2a4072b1a66802ec1e7a74071e1cbb0ad91cb3e8061997c9"></a>

## sort_values — sort_values / 7e6212ddae9a / 2

Breadcrumbs:

- [xcsh_app_security_evidence](../data-sources/app_security_evidence.md#canonical-d3c80b619742ede9129532b0a717cd17f1a275a4761a59fe318edc65c1f1d29b)
- [Property reference](data-sources--app_security_evidence--reference--group-001.md#canonical-160f5b6ebef4ab35b1510dc3be35d8747129089705ae396e4d0fa10d3481c008)
- sort_values

<a id="canonical-4d7c6d537a402b40613089630a072b2f0a898153717df66dbc84d3ef4b55081a"></a>

Type: `"single"`. Optional.

These are timestamp and doc\_id values returned by elastic search in the search request. Client is
expected to set these values in a subsequent request to GET the next page of results.

<a id="canonical-fa693bf0c64d9b0f4155fdfb5fb4bf758b32a00bb5b996be3bc722471c9f4294"></a>

## Direct properties — sort_values / 7e6212ddae9a / 3

<a id="canonical-bce97ea847c1166f77b64caf3f95b1504567837567c30f06e4428594f29d578a"></a>

<a id="canonical-0a28747eb5e82af2a19910fc18efeed3b108eb0d88e59cdac126807f55cc69dd"></a>

## last_doc_id property — sort_values / 7e6212ddae9a / 4

Type: `"string"`. Optional.

Unique UUID generated by elastic search.

<a id="canonical-3648713bb0a27f87e137b8d7337a791170cc20f29cc2e0f121af446580de9afa"></a>

<a id="canonical-c8ed01ea72e344723b188ec49536f92427c58fd8434514a2f75c2e4de009405d"></a>

## last_timestamp property — sort_values / 7e6212ddae9a / 5

Type: `"number"`. Optional.

Configuration parameter for last timestamp.

<a id="canonical-2f7e385f050deb50a12b947fbe80bf6eda0e772eb3401c25898cc37717b39a8b"></a>

## Next pages — sort_values / 7e6212ddae9a / 6

- [Property reference](data-sources--app_security_evidence--reference--group-001.md#canonical-160f5b6ebef4ab35b1510dc3be35d8747129089705ae396e4d0fa10d3481c008)
- [xcsh_app_security_evidence](../data-sources/app_security_evidence.md#canonical-d3c80b619742ede9129532b0a717cd17f1a275a4761a59fe318edc65c1f1d29b)
