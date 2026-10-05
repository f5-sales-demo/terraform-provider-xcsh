---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-3231000330200203-1000233102323220-1030123212002110-3200002223210333-2302222010321112-3201011100222033-3131213230203032-3020211100131010"></a>

## Next pages — any_domain / 211213123132 / 4

- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111320203131221-0032003221131003-3030112302030003-3022302323232121-2101301002303122-1022103210221300-2120131223331200-3233133102011000"></a>

## api_rate_limit.server_url_rules.client_matcher — client_matcher / 211020013301 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- api_rate_limit.server_url_rules.client_matcher

<a id="canonical-3022123012300011-2103200333113130-0102301133332210-1211301022221310-0311233230030231-2302321330302021-3222033010131111-2200330020300202"></a>

Type: `"object"`. single nested block, Optional.

Client Matcher. Client conditions for matching a rule.

Upstream description:

Client conditions for matching a rule.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("any_client",
    "client_selector"),
  validators.ConflictingObjectAttributes("any_client",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("any_ip",
    "asn_list"),
  validators.ConflictingObjectAttributes("any_ip",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("asn_list",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("asn_list",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("asn_list",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("asn_matcher",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("asn_matcher",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("client_selector",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("ip_matcher",
    "ip_prefix_list")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-client_choice": "[\"any_client\",\"client_selector\",\"ip_threat_category_list\"]",
  "x-ves-oneof-field-ip_asn_choice": "[\"any_ip\",\"asn_list\",\"asn_matcher\",\"ip_matcher\",\"ip_prefix_list\"]"
}
```

Terraform syntax:

```terraform
client_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0133302333222022-1311013230313311-1021200030130312-1030201221020102-0313320333310000-3023011230130123-2302231313330102-3202300033230222"></a>

## Direct properties — client_matcher / 211020013301 / 3

- [any_client](resources--cdn_loadbalancer--reference--group-005.md#canonical-1001121311313002-0320330200133230-1323323031130323-2203320203133231-0311032312303320-0012313023130210-1330202011323221-1101331223002310): complete subsection reference.

- [any_ip](resources--cdn_loadbalancer--reference--group-005.md#canonical-3211100132121310-2001123001112333-0210021011001302-2030301130323123-1232210201110233-0311333300232100-0311231200201322-2211320102020111): complete subsection reference.

- [asn_list](resources--cdn_loadbalancer--reference--group-005.md#canonical-3021213112110103-3303110222212011-3311210320211323-2033201233112103-1013012103023313-0330003332000011-3100011230000132-0121232220300223): complete subsection reference.

- [asn_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-3123212302233332-0002110301331331-1012311122210102-0301230130022331-3210120210310121-3333320313133003-1131312131002201-1023000203303332): complete subsection reference.

- [client_selector](resources--cdn_loadbalancer--reference--group-005.md#canonical-0011031133302211-0202030230313323-0033223313130231-3010013233323210-1023223112002133-0201303202030302-1023011112031120-3133331110222023): complete subsection reference.

- [ip_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2130111131232013-1220123201311202-3133000103231121-2322113333330323-0202110311211331-3113022121103001-1301223132322213-2023303310002223): complete subsection reference.

- [ip_prefix_list](resources--cdn_loadbalancer--reference--group-005.md#canonical-1121321233302210-2201233030311302-0233012110112111-3301022303130020-3330230211332201-1122013113002133-3100021113102003-1302122332233003): complete subsection reference.

- [ip_threat_category_list](resources--cdn_loadbalancer--reference--group-005.md#canonical-3000100122201321-1031000322331033-0320100200312313-0310123303011001-0231212022301023-2030131301002101-0231000210030100-2010003030323113): complete subsection reference.

- [tls_fingerprint_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-0130003213232033-2202300221200002-2331102232130333-0210030322312120-3203313330213022-0310130203033112-1123331112033221-0112233032020000): complete subsection reference.

<a id="canonical-1120003302103113-0323133202331003-1302023023231132-3211330321310020-2213111303112102-3121123103203200-1122103131103332-3313011030101100"></a>

## Next pages — client_matcher / 211020013301 / 4

- [api_rate_limit.server_url_rules.client_matcher.any_client](resources--cdn_loadbalancer--reference--group-005.md#canonical-1001121311313002-0320330200133230-1323323031130323-2203320203133231-0311032312303320-0012313023130210-1330202011323221-1101331223002310)
- [api_rate_limit.server_url_rules.client_matcher.any_ip](resources--cdn_loadbalancer--reference--group-005.md#canonical-3211100132121310-2001123001112333-0210021011001302-2030301130323123-1232210201110233-0311333300232100-0311231200201322-2211320102020111)
- [api_rate_limit.server_url_rules.client_matcher.asn_list](resources--cdn_loadbalancer--reference--group-005.md#canonical-3021213112110103-3303110222212011-3311210320211323-2033201233112103-1013012103023313-0330003332000011-3100011230000132-0121232220300223)
- [api_rate_limit.server_url_rules.client_matcher.asn_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-3123212302233332-0002110301331331-1012311122210102-0301230130022331-3210120210310121-3333320313133003-1131312131002201-1023000203303332)
- [api_rate_limit.server_url_rules.client_matcher.client_selector](resources--cdn_loadbalancer--reference--group-005.md#canonical-0011031133302211-0202030230313323-0033223313130231-3010013233323210-1023223112002133-0201303202030302-1023011112031120-3133331110222023)
- [api_rate_limit.server_url_rules.client_matcher.ip_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2130111131232013-1220123201311202-3133000103231121-2322113333330323-0202110311211331-3113022121103001-1301223132322213-2023303310002223)
- [api_rate_limit.server_url_rules.client_matcher.ip_prefix_list](resources--cdn_loadbalancer--reference--group-005.md#canonical-1121321233302210-2201233030311302-0233012110112111-3301022303130020-3330230211332201-1122013113002133-3100021113102003-1302122332233003)
- [api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list](resources--cdn_loadbalancer--reference--group-005.md#canonical-3000100122201321-1031000322331033-0320100200312313-0310123303011001-0231212022301023-2030131301002101-0231000210030100-2010003030323113)
- [api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-0130003213232033-2202300221200002-2331102232130333-0210030322312120-3203313330213022-0310130203033112-1123331112033221-0112233032020000)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1001121311313002-0320330200133230-1323323031130323-2203320203133231-0311032312303320-0012313023130210-1330202011323221-1101331223002310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133100023202030-2222120333230032-0102221203103123-0321323223013223-0330122203121111-0133022302322222-1122003120100300-1322032130202301"></a>

## api_rate_limit.server_url_rules.client_matcher.any_client — any_client / 312003323330 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311)
- api_rate_limit.server_url_rules.client_matcher.any_client

<a id="canonical-3123110111101202-0002001213332012-3310023202211032-3123133332001222-2122220231330132-1111232231222211-1113311121013213-1231012332323121"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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
any_client = {}
```

<a id="canonical-2220021022212313-0221213032301130-0023300312130131-2220031121203222-0211201222102120-3210230232313111-1000021132200202-1223121031013123"></a>

## Direct properties — any_client / 312003323330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3221131001022302-3122332000001110-2223223321232322-0131031023111031-0212312022320232-2010102322220323-1302213203220133-2032200301303101"></a>

## Next pages — any_client / 312003323330 / 4

- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3211100132121310-2001123001112333-0210021011001302-2030301130323123-1232210201110233-0311333300232100-0311231200201322-2211320102020111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221012123031030-1132021211032221-2212203001200312-0103013112103002-2230000311322112-3002101233132032-2011031122023202-0222122310202011"></a>

## api_rate_limit.server_url_rules.client_matcher.any_ip — any_ip / 202232003300 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311)
- api_rate_limit.server_url_rules.client_matcher.any_ip

<a id="canonical-0100323203331322-2033223120312221-0210210332103201-3003303023201133-2000032210301022-0112013022110200-1030102100332132-2311321330121233"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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
any_ip = {}
```

<a id="canonical-1312022030101131-3011030033120323-1123132323221023-0202013220203022-0232121300130212-3233233003023110-0131032203322330-2313303031233233"></a>

## Direct properties — any_ip / 202232003300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0022332320000211-2113221013303330-1111102303223123-3220023311101002-1020332023011302-2130200110301031-1301021231030122-3332111332210123"></a>

## Next pages — any_ip / 202232003300 / 4

- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3021213112110103-3303110222212011-3311210320211323-2033201233112103-1013012103023313-0330003332000011-3100011230000132-0121232220300223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122112213112121-1330023112113230-2130031122303021-1202003233320233-1321332212131302-1030221220330132-3031000123313023-0111203031100033"></a>

## api_rate_limit.server_url_rules.client_matcher.asn_list — asn_list / 223120112102 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311)
- api_rate_limit.server_url_rules.client_matcher.asn_list

<a id="canonical-3003310000130221-3211202233131031-3232223231231101-2111033313111333-2110233120230221-2200303001022122-0000211303102333-0022223212030001"></a>

Type: `"object"`. single nested block, Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("as_numbers")}
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
asn_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0312133022011010-2120011321310101-3133221331130031-1113030311201213-3000221230103111-3022303110301010-2200212133302122-0332112213220103"></a>

## Direct properties — asn_list / 223120112102 / 3

<a id="canonical-3213100002313123-0312121303011121-0332213023302222-0003113201231332-3210212233000231-3121011231030210-1300320003223321-1022001301222212"></a>

<a id="canonical-2123201122132200-3233200120020321-2013313223021003-1312020123103100-0313032123022221-0302223032131103-3132233200231232-1202301021233202"></a>

## as_numbers property — asn_list / 223120112102 / 4

Type: `["list", "number"]`. Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3122333101132321-2210210002310110-0233130023023220-1332121100132323-0203301012300112-3101003011102311-0031120301121220-1003012323112000"></a>

## Next pages — asn_list / 223120112102 / 5

- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3123212302233332-0002110301331331-1012311122210102-0301230130022331-3210120210310121-3333320313133003-1131312131002201-1023000203303332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3201310112022032-2023103133113001-2213131111331233-1033123332120001-2332220210110003-1000123322110231-0132301103012120-1210201203232301"></a>

## api_rate_limit.server_url_rules.client_matcher.asn_matcher — asn_matcher / 130313212222 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311)
- api_rate_limit.server_url_rules.client_matcher.asn_matcher

<a id="canonical-2321223333223122-2213320101023223-0100110102303113-0303010202003112-3100101212020320-0100120322232320-3222001210023000-2223212321232100"></a>

Type: `"object"`. single nested block, Optional.

Match any AS number contained in the list of bgp\_asn\_sets.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("asn_sets")}
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
asn_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0202013031031221-3312210022021220-3310231102211330-1121011132333030-3023021200303121-2031222020013111-2032230111000002-1211320211023123"></a>

## Direct properties — asn_matcher / 130313212222 / 3

- [asn_sets](resources--cdn_loadbalancer--reference--group-005.md#canonical-0020011133130001-2133103330222323-3200110323220330-3113332022021002-1101003110330122-1312000313210012-2322333200133203-3022332123223111): complete subsection reference.

<a id="canonical-0221322221023033-3333303232120001-1100220112100331-3101010231121213-1111110120301031-1202200333132233-3101300001311032-2020222331302133"></a>

## Next pages — asn_matcher / 130313212222 / 4

- [api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets](resources--cdn_loadbalancer--reference--group-005.md#canonical-0020011133130001-2133103330222323-3200110323220330-3113332022021002-1101003110330122-1312000313210012-2322333200133203-3022332123223111)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0020011133130001-2133103330222323-3200110323220330-3113332022021002-1101003110330122-1312000313210012-2322333200133203-3022332123223111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3111321032223010-0333313030012033-2220032111203312-2113033310023333-2111301122030223-1033212300120232-3230101321310133-3003300223201300"></a>

## api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets — asn_sets / 022332010113 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311)
- [api_rate_limit.server_url_rules.client_matcher.asn_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-3123212302233332-0002110301331331-1012311122210102-0301230130022331-3210120210310121-3333320313133003-1131312131002201-1023000203303332)
- api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets

<a id="canonical-2011210033032032-1223022200233010-2211012313321331-0311120032133121-2122313023230323-3012231031300333-0003020002310111-1330011233111101"></a>

Type: `"object"`. list nested block, Optional.

List of references to bgp\_asn\_set objects.

Upstream description:

A list of references to bgp\_asn\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
asn_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-2231001010301100-2021301102001323-0112131203313222-0201023023333132-0301202031001102-1002013210210101-2313120132221303-3221302330312233"></a>

## Direct properties — asn_sets / 022332010113 / 3

<a id="canonical-1110003321222232-1102313112131201-1130311223232010-3321321011003023-0011130232033033-2320130002130111-2220203001032232-1101012130010232"></a>

<a id="canonical-2100321122132122-0103320123331123-1332030202131220-0220302000213130-3322032203320223-1132200122303101-0133120202222301-2203211313010330"></a>

## kind property — asn_sets / 022332010113 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1023211130023111-0033131013311213-0230100232020031-3320332222011011-0311331212032231-2200223022002312-2312303021332111-2021322113110003"></a>

<a id="canonical-1200011032231232-0210122322303022-3001003021312200-3131303301010011-2202312130021030-0301133001010021-2311203232102103-0132032232021103"></a>

## name property — asn_sets / 022332010113 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1103311032212223-1010322101001222-2120012323122302-0320002201233023-1101311333033223-0222000312100203-3232210302110212-2002203302032031"></a>

<a id="canonical-0133300233112221-0130313022023231-1002011110110131-1021311100333220-1131100130123203-0232123120212311-0201333322203132-3312222330132312"></a>

## namespace property — asn_sets / 022332010113 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3010213322220013-2003321221111332-0020330222030230-3013032310333032-3101221023030102-3300010032030022-2033021333303330-2211321023001020"></a>

<a id="canonical-3020012120312000-3103021013213320-3101321333122210-1213203323100323-3130301213212211-0302231300222202-1012113100011022-0330313122033101"></a>

## tenant property — asn_sets / 022332010113 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2103003122331323-3331100032121200-1320321201333222-0330223123132303-1100300200213123-0032202020201331-1310311230233021-2031023022300332"></a>

<a id="canonical-1003301123323123-2201301111202220-1032113103221212-0003012332002110-2123300100120333-0230011102030110-0301131012300033-3212333302331230"></a>

## uid property — asn_sets / 022332010113 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3131232320312213-0220032113001210-1100132221102110-3211330313133313-2301221322011022-0332130113132312-2123131021101311-2131100223233303"></a>

## Next pages — asn_sets / 022332010113 / 9

- [api_rate_limit.server_url_rules.client_matcher.asn_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-3123212302233332-0002110301331331-1012311122210102-0301230130022331-3210120210310121-3333320313133003-1131312131002201-1023000203303332)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0011031133302211-0202030230313323-0033223313130231-3010013233323210-1023223112002133-0201303202030302-1023011112031120-3133331110222023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123232023111100-2333031131322033-3112212032113333-3202123020313332-3200203200002300-1122100202113010-3012331030230310-3031233030330201"></a>

## api_rate_limit.server_url_rules.client_matcher.client_selector — client_selector / 021323100302 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311)
- api_rate_limit.server_url_rules.client_matcher.client_selector

<a id="canonical-2330202113000011-1312121010112333-2233133100002331-0000023113122010-3121211212303101-0302102112030031-0331101102230120-2203223000213110"></a>

Type: `"object"`. single nested block, Optional.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
client_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-3312322013133212-2023302233001132-1133330311012231-0003111021123110-3323321232122233-0332222203230021-3300102301313231-2023331101323313"></a>

## Direct properties — client_selector / 021323100302 / 3

<a id="canonical-1211102011020310-0212112023201101-2122123300203101-0310301222021120-2213023313203312-0233233002132021-2330111030112203-1132200023110132"></a>

<a id="canonical-2311132300231111-2310331311322103-2221321220313321-2020332033332220-3303121120302333-1220132212223123-0102300333100030-3012200232330013"></a>

## expressions property — client_selector / 021323100302 / 4

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-3111220212011202-1233231132200021-3323200013322310-0212333312012120-2312210222102233-0012133103001221-3301221321103231-1021220201011210"></a>

## Next pages — client_selector / 021323100302 / 5

- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2130111131232013-1220123201311202-3133000103231121-2322113333330323-0202110311211331-3113022121103001-1301223132322213-2023303310002223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323210331221133-1101331223210031-1121220020230021-1100031013301211-3111212132201003-2133201031302112-1030223010202012-2301111001010021"></a>

## api_rate_limit.server_url_rules.client_matcher.ip_matcher — ip_matcher / 332032020110 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311)
- api_rate_limit.server_url_rules.client_matcher.ip_matcher

<a id="canonical-1133122320001010-0313103132200101-3032123312232212-1333331321120222-3002030031001001-2322200013010013-3133033312223001-3101030212321101"></a>

Type: `"object"`. single nested block, Optional.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("prefix_sets")}
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
ip_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0323201011121312-3111001031232302-2322110332223302-2213303320330311-3000230311033230-2320222321220301-0301023311103021-1112033111313131"></a>

## Direct properties — ip_matcher / 332032020110 / 3

<a id="canonical-0022201132310221-3313123210033332-2033022113023313-2112300201203312-2210322133030112-1010310030122100-0110311013232332-0101333201100003"></a>

<a id="canonical-1030233320331121-1223313310120231-2120222310001123-2333201010202030-0321231101000223-1202202313322012-2333121202323213-2322110323031313"></a>

## invert_matcher property — ip_matcher / 332032020110 / 4

Type: `"bool"`. Optional.

Invert IP Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [prefix_sets](resources--cdn_loadbalancer--reference--group-005.md#canonical-3313002220000222-1013023302210203-0331011102101332-1321220011231113-1201123222220003-0201321312303122-3000133220202331-0221333322011123): complete subsection reference.

<a id="canonical-3033323203212022-2110011231033221-0233311002213200-2022013101233013-2132331222132332-1333230013230330-1003221301120013-2000033303011111"></a>

## Next pages — ip_matcher / 332032020110 / 5

- [api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets](resources--cdn_loadbalancer--reference--group-005.md#canonical-3313002220000222-1013023302210203-0331011102101332-1321220011231113-1201123222220003-0201321312303122-3000133220202331-0221333322011123)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3313002220000222-1013023302210203-0331011102101332-1321220011231113-1201123222220003-0201321312303122-3000133220202331-0221333322011123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1133022102000122-0333103331323032-2130303321010330-0130201223023121-0333101132001020-0222010312221001-1333010123222100-3010111322223001"></a>

## api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets — prefix_sets / 012033230121 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311)
- [api_rate_limit.server_url_rules.client_matcher.ip_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2130111131232013-1220123201311202-3133000103231121-2322113333330323-0202110311211331-3113022121103001-1301223132322213-2023303310002223)
- api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-0101031020122023-1110211033210023-1110100213222020-0120230323100003-1110010332130031-1103232221320222-1232221333102133-1111131212001023"></a>

Type: `"object"`. list nested block, Optional.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
prefix_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-2303313220202022-0303230112112012-0210303231021210-0302313313220123-0132223020001312-3321030000103331-0120112110122203-3302010000102030"></a>

## Direct properties — prefix_sets / 012033230121 / 3

<a id="canonical-2233301113300221-2100122220032130-1112032112330033-2213003312310012-3110023012330330-1102210002202233-3202002113022323-2211212032131020"></a>

<a id="canonical-2320033030212331-3031122120320323-2113333332222111-0201313112320203-3302323013022022-1013322010120231-1323003201131213-3231232232002323"></a>

## kind property — prefix_sets / 012033230121 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2010311313312031-3030200102030331-0310211123231023-2000313331113100-2233312003211230-2012221230121233-3233311221301333-0220222102101102"></a>

<a id="canonical-2311313313232320-2232233221332023-2320220132212233-2001300011003123-3112300113123220-0002121030312313-3131130202033133-0113112120221301"></a>

## name property — prefix_sets / 012033230121 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1001210130300100-0211013111110212-1300113100110233-0113103031001113-2312010121233213-3122101211121132-3122302322113101-1033123110302301"></a>

<a id="canonical-2200033022132330-1013032221120002-1320220311223212-1023032300110220-3223032203020123-3021113222121221-3232023231110332-3012122013032130"></a>

## namespace property — prefix_sets / 012033230121 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2230322021002103-1103130033301003-0231212221211032-2112203312103220-2000132312120131-3200020211033212-0313001110332032-2122200210132320"></a>

<a id="canonical-1021313211320202-3211233103200311-0112322030101212-1131201112030030-1302132002321200-1202002122231132-3000121202132231-1332021110332223"></a>

## tenant property — prefix_sets / 012033230121 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1100302313323012-2303303132230020-0013002222321032-1012302203132030-2002330111030013-1033203311121022-3031101133233020-0322123110002200"></a>

<a id="canonical-3032223300332221-1333002301003322-0102330231233032-2300302032022211-2100303031211122-2000301132222331-1320033212230131-1121312223020021"></a>

## uid property — prefix_sets / 012033230121 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2130303133300221-3111203032121030-3310013112221103-3203221223010033-3313203212032113-0321133031332022-0102011102122021-0010332213300322"></a>

## Next pages — prefix_sets / 012033230121 / 9

- [api_rate_limit.server_url_rules.client_matcher.ip_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2130111131232013-1220123201311202-3133000103231121-2322113333330323-0202110311211331-3113022121103001-1301223132322213-2023303310002223)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1121321233302210-2201233030311302-0233012110112111-3301022303130020-3330230211332201-1122013113002133-3100021113102003-1302122332233003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120110333210222-0332210123320230-3313311113311032-2310330332130220-1121201310212101-0201003221222331-0210123120002311-3323133312111321"></a>

## api_rate_limit.server_url_rules.client_matcher.ip_prefix_list — ip_prefix_list / 322002321003 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311)
- api_rate_limit.server_url_rules.client_matcher.ip_prefix_list

<a id="canonical-2331230113331202-0322001220233032-1100221123320311-3021001302112131-1203113033200222-2012211303003130-2020101132130300-0311012312220230"></a>

Type: `"object"`. single nested block, Optional.

List of IP Prefix strings to match against.

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
ip_prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1020221202303220-1232322313331320-1130110001130121-3321323203120203-0013212023020333-2010303322223102-3302122103313230-1000133213020202"></a>

## Direct properties — ip_prefix_list / 322002321003 / 3

<a id="canonical-3113331313220101-2331330033023320-3010103032031011-1332002231302203-3011330030323102-3003020013232031-2031300130020310-0001130101320101"></a>

<a id="canonical-0233111213100320-0023023233221213-0022120301103320-1233103131101223-2012313032233312-0030121202231231-1000311013121111-2131311221211303"></a>

## invert_match property — ip_prefix_list / 322002321003 / 4

Type: `"bool"`. Optional.

Invert Match Result. Invert the match result.

Upstream description:

Invert the match result.

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

<a id="canonical-3232112202210200-1100032001303302-2232320012101323-2212231002032123-1221132211210020-3013022110320323-2210323312203221-3122230230033311"></a>

<a id="canonical-2330132002320302-3023233210233211-3030301003220203-2202002231221333-2012033120003012-2210212021303131-3333130130102222-1222122321210130"></a>

## ip_prefixes property — ip_prefix_list / 322002321003 / 5

Type: `["list", "string"]`. Optional.

IPv4 Prefix List. List of IPv4 prefix strings.

Upstream description:

List of IPv4 prefix strings.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0221213003102122-3223010300210300-1210113000201302-2022113032121330-2203322213120303-2311103103021123-0322010301313201-0200010311232132"></a>

## Next pages — ip_prefix_list / 322002321003 / 6

- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3000100122201321-1031000322331033-0320100200312313-0310123303011001-0231212022301023-2030131301002101-0231000210030100-2010003030323113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3220130222122332-2103321002303133-3102111102203210-3213322010330120-1023021230111320-0201130121022232-0101312231010310-1302303333332233"></a>

## api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list — ip_threat_category_list / 020233132232 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311)
- api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list

<a id="canonical-3100330322300200-1331102120020231-0220330232301003-0313300300130023-2212200103120302-3001021301231221-2212230110112332-0310333103203003"></a>

Type: `"object"`. single nested block, Optional.

IP Threat Category List Type. List of IP threat categories.

Upstream description:

List of IP threat categories.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("ip_threat_categories")}
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
ip_threat_category_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3132002102330012-0022122100202113-2130223222222210-0232311323121313-3330230003022030-2100232203031310-3332002120221220-3033303231302030"></a>

## Direct properties — ip_threat_category_list / 020233132232 / 3

<a id="canonical-2100202112303322-1022020213112022-1003103301132313-0331310102130231-2202133110110223-3001310110321032-2132011331211220-1010303103330232"></a>

<a id="canonical-0130110312320130-2033000200202121-0030110212203233-2222101000330210-3102122301322303-3203331002131102-1210210112001001-1010033013103023"></a>

## ip_threat_categories property — ip_threat_category_list / 020233132232 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
SPAM\_SOURCES|WINDOWS\_EXPLOITS|WEB\_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE\_THREATS|TOR\_PROXY|DENIAL\_OF\_SERVICE|NETWORK\]
The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions. Possible values are \`SPAM\_SOURCES\`, \`WINDOWS\_EXPLOITS\`,
\`WEB\_ATTACKS\`, \`BOTNETS\`, \`SCANNERS\`, \`REPUTATION\`, \`PHISHING\`, \`PROXY\`,
\`MOBILE\_THREATS\`, \`TOR\_PROXY\`, \`DENIAL\_OF\_SERVICE\`, \`NETWORK\`. Defaults to
\`SPAM\_SOURCES\`.

Upstream description:

The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3302331332320203-3200030122300123-3000231213313013-3203310110000011-3110311113200002-2123033321303112-0223331201122001-0002212120301103"></a>

## Next pages — ip_threat_category_list / 020233132232 / 5

- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0130003213232033-2202300221200002-2331102232130333-0210030322312120-3203313330213022-0310130203033112-1123331112033221-0112233032020000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213212211312012-0020312302200302-1030311201220011-0011230213300200-1101222223133120-3223233321232132-3000310103332010-1102212223002230"></a>

## api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher — tls_fingerprint_matcher / 001230321102 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311)
- api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-1310201021312320-0113202323013002-3130321313012030-3301211231221331-2003311330120133-3222222113203213-3321210102021302-3301011103013002"></a>

Type: `"object"`. single nested block, Optional.

TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are
satisfied..

Upstream description:

A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are satisfied
and the input fingerprint is not one of the excluded values.

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
tls_fingerprint_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0003022033320311-2113331221103201-2131012003021012-2122130011132211-3103001231201220-1313201033111021-0320303130011032-0100330131210120"></a>

## Direct properties — tls_fingerprint_matcher / 001230321102 / 3

<a id="canonical-3203012101201101-3311213222002223-2332220001020010-1303122223120133-2033032101020032-0223030310131232-3113030310021210-0323000112311123"></a>

<a id="canonical-0333133110210031-3032123002333010-3211202122132032-0231313230300200-2101002110332300-0010131022300310-0021231111203133-0220002030102101"></a>

## classes property — tls_fingerprint_matcher / 001230321102 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Upstream description:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3323010032322312-2211331320121303-0033312111301333-2310230003123130-1211103200202021-1210202210102323-1332101003001111-0311330023133011"></a>

<a id="canonical-3033310200031303-3202200301120323-0211220211332100-1111130120213012-1210111032122132-0320002000203203-3021110020221003-2100202013122322"></a>

## exact_values property — tls_fingerprint_matcher / 001230321102 / 5

Type: `["list", "string"]`. Optional.

List of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Upstream description:

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0013121000303010-3220021200023022-1230331032221011-0201222311221102-0321030230130302-0310201303201321-0003012323102023-3133101202302132"></a>

<a id="canonical-2333130131023200-0112130211322130-2022332330100223-3232203200312012-2321313102300000-0133233312220233-2111222201320123-1302300223231010"></a>

## excluded_values property — tls_fingerprint_matcher / 001230321102 / 6

Type: `["list", "string"]`. Optional.

List of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be
used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Upstream description:

A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can
be used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2333103110021331-3220131112012020-1121011102333323-1133131222312131-2230203330100301-3120211121213211-2132220120132311-3201311310302303"></a>

## Next pages — tls_fingerprint_matcher / 001230321102 / 7

- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0321031211113300-3020232113001120-2211311332112302-3102201332200033-3301102331030312-2110201132013202-3010222101121131-0103331231111021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0222320320100022-0220311000321130-0312213233102103-0212200210000111-3223030221120202-1121301101013222-0133221023213210-2220322323110001"></a>

## api_rate_limit.server_url_rules.inline_rate_limiter — inline_rate_limiter / 023323223313 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- api_rate_limit.server_url_rules.inline_rate_limiter

<a id="canonical-2112001120300200-0213110213211102-2232121213220101-3130111122212112-3021133133212111-0123323133032120-0120122200022121-0321200311032112"></a>

Type: `"object"`. single nested block, Optional.

Inline rate-limiter settings for this domain, base-path, or endpoint rule. Select this field as the
required rate\_limiter\_choice when no stored rate-limiter object is used.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("threshold"),
  validators.ConflictingObjectAttributes("ref_user_id",
    "use_http_lb_user_id")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-count_by_choice": "[\"ref_user_id\",\"use_http_lb_user_id\"]"
}
```

Terraform syntax:

```terraform
inline_rate_limiter {
  # Configure direct properties listed below.
}
```

<a id="canonical-2222103100102322-2103010311230321-2110311203131231-0111310213030313-3232013030202223-2102032102023310-1230212232121023-1233303103232023"></a>

## Direct properties — inline_rate_limiter / 023323223313 / 3

- [ref_user_id](resources--cdn_loadbalancer--reference--group-005.md#canonical-0013311012330202-0121212300112232-0023221120001330-0102000030031120-3321233120230310-2003230003222132-1020133112230313-2220300233121113): complete subsection reference.

<a id="canonical-3101110210003223-0233211032303021-0013321113303300-1202210220023000-2000302331313302-2332331311032323-2003211210113012-1103230101010110"></a>

<a id="canonical-0310300201010210-1022130203000310-0021131030231313-2231013210232331-3031023320033201-0033313023332003-3322212113332112-1011320320322021"></a>

## threshold property — inline_rate_limiter / 023323223313 / 4

Type: `"number"`. Optional.

The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified
period.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  }
}
```

<a id="canonical-1131032333200223-2101000311023221-3121312300232221-3313113021113000-2323030201312213-0113231210233033-0233201203113010-0001301300000013"></a>

<a id="canonical-3020333003022311-2110331202233033-3221200201101323-3312131131223222-1301300223003122-2022102220121122-1113101121210232-0111210123310202"></a>

## unit property — inline_rate_limiter / 023323223313 / 5

Type: `"string"`. Optional.

\[Enum: SECOND|MINUTE|HOUR\] Unit for the period per which the rate limit is applied. - SECOND:
Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR:
Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days. Possible values are
\`SECOND\`, \`MINUTE\`, \`HOUR\`. Defaults to \`SECOND\`.

Upstream description:

Unit for the period per which the rate limit is applied.

&#8203;- SECOND: Second

Rate limit period unit is seconds &#8203;- MINUTE: Minute

Rate limit period unit is minutes &#8203;- HOUR: Hour

Rate limit period unit is hours &#8203;- DAY: Day

Rate limit period unit is days.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["HOUR","MINUTE","SECOND"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("SECOND",
    "MINUTE",
    "HOUR"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "SECOND",
  "enum": [
    "SECOND",
    "MINUTE",
    "HOUR"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [use_http_lb_user_id](resources--cdn_loadbalancer--reference--group-005.md#canonical-2230212032311201-2231312302320020-3212020200322021-0003131301030311-3323121300123202-3022102231011330-1031030320220021-1212332030021031): complete subsection reference.

<a id="canonical-2220003222030020-1233321033120001-0212130323112222-1322122211133131-2303103300223310-0332122303021331-0021322032130330-3111133121300221"></a>

## Next pages — inline_rate_limiter / 023323223313 / 6

- [api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id](resources--cdn_loadbalancer--reference--group-005.md#canonical-0013311012330202-0121212300112232-0023221120001330-0102000030031120-3321233120230310-2003230003222132-1020133112230313-2220300233121113)
- [api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id](resources--cdn_loadbalancer--reference--group-005.md#canonical-2230212032311201-2231312302320020-3212020200322021-0003131301030311-3323121300123202-3022102231011330-1031030320220021-1212332030021031)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0013311012330202-0121212300112232-0023221120001330-0102000030031120-3321233120230310-2003230003222132-1020133112230313-2220300233121113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212332111133030-3021110311130222-0331031200303010-3320032111213320-3003210220321123-1313113120133210-2333133201021033-1222030203033223"></a>

## api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id — ref_user_id / 112331131310 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.inline_rate_limiter](resources--cdn_loadbalancer--reference--group-005.md#canonical-0321031211113300-3020232113001120-2211311332112302-3102201332200033-3301102331030312-2110201132013202-3010222101121131-0103331231111021)
- api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id

<a id="canonical-2221100320102101-2332302121002013-2031011233230131-0210222330203303-3102102103311033-2010321020322101-3311332322002220-0020101222220000"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
ref_user_id {
  # Configure direct properties listed below.
}
```

<a id="canonical-1123333123011200-3200011301110112-1033002002333103-2201212221220313-2112112200322320-0102231231013123-3132103313310122-2300330203321122"></a>

## Direct properties — ref_user_id / 112331131310 / 3

<a id="canonical-1331313131200122-3120223221230121-3023303230012233-3233113332331322-2120032113321311-3313232331031220-0221110330231133-0233131011301210"></a>

<a id="canonical-3320323122221231-2313112122010231-1300220332033003-2022212210101223-3030000130310331-3313220012103231-1111123131221322-1131302003103233"></a>

## name property — ref_user_id / 112331131310 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-1313212020221323-0300033320200012-0123031323123133-1303013311321231-0331101130033123-0113002201320323-0333031200010132-1302032323222002"></a>

<a id="canonical-2213023012002331-0222000030020013-0211322200312201-3231232320313302-1033013113201113-0303223103213303-0311031003010331-3222300022133223"></a>

## namespace property — ref_user_id / 112331131310 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0232132330300231-0111032320011321-2000013212132010-2230100000011323-1132130320111112-2012023310230110-0222120231231020-1312333013112012"></a>

<a id="canonical-3023013113223020-2132310202232332-0220113310313210-1313000021101233-2300313012131212-2113300102301203-2132331200131001-0001100320011202"></a>

## tenant property — ref_user_id / 112331131310 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2311211122220100-0111131320033001-3322130330102313-2200022013001030-3022013133213100-0103013020221333-0222223000333133-3313231230302201"></a>

## Next pages — ref_user_id / 112331131310 / 7

- [api_rate_limit.server_url_rules.inline_rate_limiter](resources--cdn_loadbalancer--reference--group-005.md#canonical-0321031211113300-3020232113001120-2211311332112302-3102201332200033-3301102331030312-2110201132013202-3010222101121131-0103331231111021)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2230212032311201-2231312302320020-3212020200322021-0003131301030311-3323121300123202-3022102231011330-1031030320220021-1212332030021031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3100331300223002-0130211333321022-2212113003210231-0221120031001223-0310221220032000-1322232103013300-0301102001102211-0032231223020310"></a>

## api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id — use_http_lb_user_id / 110022313210 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.inline_rate_limiter](resources--cdn_loadbalancer--reference--group-005.md#canonical-0321031211113300-3020232113001120-2211311332112302-3102201332200033-3301102331030312-2110201132013202-3010222101121131-0103331231111021)
- api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id

<a id="canonical-1123112033221002-0231000020012233-1321001311200030-1312102331121120-0211303221233112-0310011012131132-1000022312010311-3031223201010331"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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
use_http_lb_user_id = {}
```

<a id="canonical-1101223333311232-2003230113213322-1303103332313003-2032012110211223-3321021112332020-0033011132321321-2200302223033330-2032012123233122"></a>

## Direct properties — use_http_lb_user_id / 110022313210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0332000301230312-1333220302310002-3030133123012220-2011022210331200-3211101223203031-0233211221012133-1310332021210022-3301313123032302"></a>

## Next pages — use_http_lb_user_id / 110022313210 / 4

- [api_rate_limit.server_url_rules.inline_rate_limiter](resources--cdn_loadbalancer--reference--group-005.md#canonical-0321031211113300-3020232113001120-2211311332112302-3102201332200033-3301102331030312-2110201132013202-3010222101121131-0103331231111021)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1210111123202013-2212101212113221-1330130203301223-2331201130033000-0332303011120301-2031002033003132-3203032332303101-2213233012110233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101322130203111-1310010102312003-0210003033101212-0201020011231100-0020133301232101-0330201321312012-0231200313203301-0200110011320311"></a>

## api_rate_limit.server_url_rules.ref_rate_limiter — ref_rate_limiter / 303001121222 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- api_rate_limit.server_url_rules.ref_rate_limiter

<a id="canonical-0133331221033120-3233110330132200-3212312233133233-0010301332311201-2202031302133023-3321323221000000-3303321133030130-2312121000012211"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

Reference to a stored rate-limiter object for this scoped rule. Select exactly one of
ref\_rate\_limiter and inline\_rate\_limiter.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
ref_rate_limiter {
  # Configure direct properties listed below.
}
```

<a id="canonical-2332101110033203-1300120231112230-0031201031030212-1213033110121001-2100201201103231-3202111100332123-2013222200232321-0230311102032120"></a>

## Direct properties — ref_rate_limiter / 303001121222 / 3

<a id="canonical-3220303000310130-2113300121023031-3021011010133132-1201002210112230-2331033030131022-3331103110210123-3110002333003320-1322112103123133"></a>

<a id="canonical-3312001211131301-1303011132203331-1323231023321303-2332230303103333-1230233322021033-2232301030212121-1202022133222311-3001032113322022"></a>

## name property — ref_rate_limiter / 303001121222 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-0022000301213133-3131320033000011-0133100101123313-2322022132032131-2223230101103302-1020300310023313-2310030230210002-0000010131323221"></a>

<a id="canonical-3012231332121111-1323310222000302-1200233130310022-3111210133022102-0133210110312233-0332330102332021-3231311013233033-2000110001130023"></a>

## namespace property — ref_rate_limiter / 303001121222 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3131101233001030-3321310132330203-2123120313210230-2020011113223323-2123100211101112-3221203221011230-3231022120313012-1302111112033023"></a>

<a id="canonical-0023013111210110-0330031030332201-2112212111303212-2121033120101030-2110322110103013-3302113032132022-1330323010211013-1013133302331212"></a>

## tenant property — ref_rate_limiter / 303001121222 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0323212113233322-0111231110212131-3132032032222012-1102221110133011-0010122110303001-0021310312311122-2031112301323112-0302320330312300"></a>

## Next pages — ref_rate_limiter / 303001121222 / 7

- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3212010010100023-0332312230211111-0332132021313201-1102330202122113-0210233132233203-3011103220223213-3231301212010300-3131130003222313"></a>

## api_rate_limit.server_url_rules.request_matcher — request_matcher / 112311210322 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- api_rate_limit.server_url_rules.request_matcher

<a id="canonical-0011033133021301-2313223002222133-2222211103133300-2201212331231002-0020103000321310-1221311022111220-0310300313120013-2000012312013102"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for request matcher.

Upstream description:

Request conditions for matching a rule.

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
request_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-3120203012003012-3211310022100323-0332232232033320-1010200003103110-1023012233221002-0002102310211331-2112230333032320-3301112121003132"></a>

## Direct properties — request_matcher / 112311210322 / 3

- [cookie_matchers](resources--cdn_loadbalancer--reference--group-005.md#canonical-3033211112003230-2302031233322223-3011210231211232-3130203102313021-2123313232211002-0033322022102113-0323310000003111-2330301133031030): complete subsection reference.

- [headers](resources--cdn_loadbalancer--reference--group-005.md#canonical-0002122103212131-2020122213023330-0233012102211313-2223031022231311-3110202023102222-1303322023020132-3300313112001132-0021201011102323): complete subsection reference.

- [jwt_claims](resources--cdn_loadbalancer--reference--group-005.md#canonical-2120002012323320-0210320131230122-1101202032031212-0321112122211211-2002131210323022-3122021323013001-1311113321020303-2332232002032320): complete subsection reference.

- [query_params](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122313023112111-2300122102223032-0012011121332013-0133313131303331-3101132233230020-2203113102113133-3002130201213302-3001033321123231): complete subsection reference.

<a id="canonical-2332300221330021-1012220130102322-2023210132313132-2230123113021023-0113020113132300-0311303212333231-0223220000133211-2331213000332230"></a>

## Next pages — request_matcher / 112311210322 / 4

- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-005.md#canonical-3033211112003230-2302031233322223-3011210231211232-3130203102313021-2123313232211002-0033322022102113-0323310000003111-2330301133031030)
- [api_rate_limit.server_url_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-005.md#canonical-0002122103212131-2020122213023330-0233012102211313-2223031022231311-3110202023102222-1303322023020132-3300313112001132-0021201011102323)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-005.md#canonical-2120002012323320-0210320131230122-1101202032031212-0321112122211211-2002131210323022-3122021323013001-1311113321020303-2332232002032320)
- [api_rate_limit.server_url_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122313023112111-2300122102223032-0012011121332013-0133313131303331-3101132233230020-2203113102113133-3002130201213302-3001033321123231)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3033211112003230-2302031233322223-3011210231211232-3130203102313021-2123313232211002-0033322022102113-0323310000003111-2330301133031030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030222002331123-1221330333200122-2302331232132210-0032032213222000-2210323332023203-3211300100212003-2311021302113210-0323303203123102"></a>

## api_rate_limit.server_url_rules.request_matcher.cookie_matchers — cookie_matchers / 111101110001 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers

<a id="canonical-0101000130230322-2011232033211002-0003000213202020-2312130230100211-2213301330013323-2123301023312131-3202301332100011-0033010020003330"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for all cookies that need to be matched. The criteria for matching each cookie is
described in individual instances of CookieMatcherType. The actual cookie values are extracted from
the request API as a list of strings for each cookie name.

Upstream description:

A list of predicates for all cookies that need to be matched. The criteria for matching each cookie
is described in individual instances of CookieMatcherType. The actual cookie values are extracted
from the request API as a list of strings for each cookie name. Note that all specified cookie
matcher predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
cookie_matchers {
  # Configure direct properties listed below.
}
```

<a id="canonical-1223322231020201-1100023132300203-1321122322032230-0003120131320333-2002022210102303-1122312020302311-2002301320012302-0123102312232102"></a>

## Direct properties — cookie_matchers / 111101110001 / 3

- [check_not_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-0022023013102010-2123110121321310-0012323310123032-1133003210323313-2002012113210013-1103120130131013-1022022322311101-0313121302122002): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-1120300322323132-1200130313310131-0002203230003120-3003321321113332-2222331133131013-3230333133133113-0122033123232211-2031231022331323): complete subsection reference.

<a id="canonical-1112200200312101-1110130302231220-2222323232222101-2303032302221021-3333033210122133-2022320001122201-1210121032122131-2113201002222102"></a>

<a id="canonical-2012321100033122-2213011003232213-2231203111013222-1022130331322231-3133112020031121-2031300021030222-0202302202222010-0230220120122320"></a>

## invert_matcher property — cookie_matchers / 111101110001 / 4

Type: `"bool"`. Optional.

Invert Matcher. Invert Match of the expression defined.

Upstream description:

Invert Match of the expression defined.

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

- [item](resources--cdn_loadbalancer--reference--group-005.md#canonical-2200003022203033-3232100113133101-3330303230031122-3310313132032300-1302323220000332-2100013130213202-1020331333022300-3221312111301310): complete subsection reference.

<a id="canonical-3333011123232033-1113031111122313-2000022130321320-0332313303321110-0302003002211211-3112221331132202-2110122330210303-1230221203022102"></a>

<a id="canonical-0212220001313113-2113132211312322-3310020132210012-0122103012123012-0131103100320300-0213023032132012-1331120323122313-1302100120220100"></a>

## name property — cookie_matchers / 111101110001 / 5

Type: `"string"`. Optional.

Cookie Name. A case-sensitive cookie name.

Upstream description:

A case-sensitive cookie name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-3330311320011321-2302230213223031-2000101210122012-0312001011233131-0333332110020203-3100002302330030-3133302210330223-1221310201211231"></a>

## Next pages — cookie_matchers / 111101110001 / 6

- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-0022023013102010-2123110121321310-0012323310123032-1133003210323313-2002012113210013-1103120130131013-1022022322311101-0313121302122002)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-1120300322323132-1200130313310131-0002203230003120-3003321321113332-2222331133131013-3230333133133113-0122033123232211-2031231022331323)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item](resources--cdn_loadbalancer--reference--group-005.md#canonical-2200003022203033-3232100113133101-3330303230031122-3310313132032300-1302323220000332-2100013130213202-1020331333022300-3221312111301310)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0022023013102010-2123110121321310-0012323310123032-1133003210323313-2002012113210013-1103120130131013-1022022322311101-0313121302122002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2323000332113303-0222310210233201-1001113230332131-0133320333113100-1223310332030123-2313113100013212-1121111202203131-3112030311112333"></a>

## api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_present — check_not_present / 331301233021 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-005.md#canonical-3033211112003230-2302031233322223-3011210231211232-3130203102313021-2123313232211002-0033322022102113-0323310000003111-2330301133031030)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-2221210220332301-0132213302302221-0321121032111120-0323121203103301-2202313022220120-0033203221102333-1210021233111330-1101332333301310"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

Upstream description:

This can be used for messages where no values are needed.

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
check_not_present = {}
```

<a id="canonical-2133010310113131-1202112333031310-0220331312332230-3023032200000300-0310113201023033-3320031030220320-2221112213302323-2132113011320112"></a>

## Direct properties — check_not_present / 331301233021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1103300011012101-0121120132331320-2311003022103212-3102231012011222-0220112132020211-2111111321200013-3232320233102322-3103222020032112"></a>

## Next pages — check_not_present / 331301233021 / 4

- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-005.md#canonical-3033211112003230-2302031233322223-3011210231211232-3130203102313021-2123313232211002-0033322022102113-0323310000003111-2330301133031030)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1120300322323132-1200130313310131-0002203230003120-3003321321113332-2222331133131013-3230333133133113-0122033123232211-2031231022331323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2322130330231301-2321233212212301-0133123300010123-1110021013120203-3221120313011021-1211033113001330-1003021023230203-3000201023201333"></a>

## api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present — check_present / 202100320101 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-005.md#canonical-3033211112003230-2302031233322223-3011210231211232-3130203102313021-2123313232211002-0033322022102113-0323310000003111-2330301133031030)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-1212323132202333-3100231123220010-3221133222020002-1113231032331223-0212001003010231-3300301102120313-1033101201303211-3331112210001302"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

Upstream description:

This can be used for messages where no values are needed.

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
check_present = {}
```

<a id="canonical-2221113101212100-2303101320313200-1110110131100301-0311211220203120-2213220130322212-1030223302201210-1320211133102020-0233103020200123"></a>

## Direct properties — check_present / 202100320101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2312313303310110-2122311110120102-2311011232300322-0010123232001023-3313002233213300-0121220333020011-2231223110310330-2322201010223230"></a>

## Next pages — check_present / 202100320101 / 4

- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-005.md#canonical-3033211112003230-2302031233322223-3011210231211232-3130203102313021-2123313232211002-0033322022102113-0323310000003111-2330301133031030)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2200003022203033-3232100113133101-3330303230031122-3310313132032300-1302323220000332-2100013130213202-1020331333022300-3221312111301310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032200321210012-0112030300313000-1210211020212012-3230203230311122-0123103231002201-0132023103130202-2032102010100032-0200121222122023"></a>

## api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item — item / 031130132133 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-005.md#canonical-3033211112003230-2302031233322223-3011210231211232-3130203102313021-2123313232211002-0033322022102113-0323310000003111-2330301133031030)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item

<a id="canonical-1221033222211231-2200220211030101-2131030103002200-1222101223133313-0110101121231303-3321310110032300-1121122211110310-3002002210312003"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-0220330120232233-3120000033013203-0312113320333201-3111012130213232-2221012000121303-2323220231322032-1110002022121332-2313030021313320"></a>

## Direct properties — item / 031130132133 / 3

<a id="canonical-3332322123201122-1212103112233210-0111323302033203-0233023302313332-3013022230330011-1303020320103202-1223302110120010-2313101032231122"></a>

<a id="canonical-2020111230310000-3311112213320130-0233313101223022-2133311122323301-2212133032020110-2320203133100320-0022330120010203-0003331122303011"></a>

## exact_values property — item / 031130132133 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0202203113130203-0132320302012012-0230210120023011-3211103000002021-3111122303301101-1201330233013020-2021321331212130-0032200221011133"></a>

<a id="canonical-3111313303332022-0223301333123213-3002322032110211-3002213122133112-1110003232003010-0123331022011220-3003310101031300-2003233121110211"></a>

## regex_values property — item / 031130132133 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3212000331131221-2222003010021132-3130133212133013-2021110201321023-0103320222311121-2222211002213101-1233031003302010-3203310132201031"></a>

<a id="canonical-3010123010123133-0023220020222321-1130133330123122-0130213332031023-1131130100312320-0331112113102311-3323033211223113-0022331020023232"></a>

## transformers property — item / 031130132133 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0020000030301230-1101221020222033-3100201022002123-3103213301300230-2012222212103121-3022112311302000-1011211001101312-1223103222303022"></a>

## Next pages — item / 031130132133 / 7

- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-005.md#canonical-3033211112003230-2302031233322223-3011210231211232-3130203102313021-2123313232211002-0033322022102113-0323310000003111-2330301133031030)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0002122103212131-2020122213023330-0233012102211313-2223031022231311-3110202023102222-1303322023020132-3300313112001132-0021201011102323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3130320231032101-3103032122001331-2001023133131030-0131232032303321-1312213230312130-2130332023321310-2232133202231302-0223311222212011"></a>

## api_rate_limit.server_url_rules.request_matcher.headers — headers / 012112131111 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- api_rate_limit.server_url_rules.request_matcher.headers

<a id="canonical-1331322133020001-2023222003233133-2323020002012031-2213132122033232-2012032322021212-2232231101003330-0302031011300010-3333121120202033"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for various HTTP headers that need to match. The criteria for matching each HTTP
header are described in individual HeaderMatcherType instances. The actual HTTP header values are
extracted from the request API as a list of strings for each HTTP header type.

Upstream description:

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-2322120222100330-0330222002023001-1110330101102323-0323020320221032-1103323221233312-1213312113312020-2013332113021123-3100310123332220"></a>

## Direct properties — headers / 012112131111 / 3

- [check_not_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-3313133023110133-3231332201323201-0232132312101311-3230130032312001-1212102113222112-3202002113133022-0032222013230302-2122333100023121): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-1331001203220231-1210232321130113-0002023333310220-1120311111210203-3123123120321320-3330133133100323-0233320030311210-1120022231031332): complete subsection reference.

<a id="canonical-0223302231320013-3222103102122030-2211310321022223-1001201220022012-1110323220030231-3011201301100002-0223103011310102-0112331210113201"></a>

<a id="canonical-0120201012200232-1103021132010202-2202232002320120-0302322212001333-1301000330012021-1300110102220333-1130212330321012-0102132000000220"></a>

## invert_matcher property — headers / 012112131111 / 4

Type: `"bool"`. Optional.

Invert Header Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [item](resources--cdn_loadbalancer--reference--group-005.md#canonical-0221231223321303-3120300201222020-1120200011202210-3211331311120130-0220332033231000-1011003020230110-0311133003003031-3230012101233210): complete subsection reference.

<a id="canonical-0032203303231223-0303121210231230-1132011333221301-0002123111132023-1330200221002132-0323002200132303-0131132022232113-1132201032230313"></a>

<a id="canonical-1033210311012210-1123230210022023-3321103222030112-1033211121222231-1111102122110300-3113131321002211-0302211022031312-0102131110233011"></a>

## name property — headers / 012112131111 / 5

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-2231200232110032-3002100111033130-2002033303332133-2233001201203222-2320030133122133-1020330232131333-2300301132322130-1120313101030102"></a>

## Next pages — headers / 012112131111 / 6

- [api_rate_limit.server_url_rules.request_matcher.headers.check_not_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-3313133023110133-3231332201323201-0232132312101311-3230130032312001-1212102113222112-3202002113133022-0032222013230302-2122333100023121)
- [api_rate_limit.server_url_rules.request_matcher.headers.check_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-1331001203220231-1210232321130113-0002023333310220-1120311111210203-3123123120321320-3330133133100323-0233320030311210-1120022231031332)
- [api_rate_limit.server_url_rules.request_matcher.headers.item](resources--cdn_loadbalancer--reference--group-005.md#canonical-0221231223321303-3120300201222020-1120200011202210-3211331311120130-0220332033231000-1011003020230110-0311133003003031-3230012101233210)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3313133023110133-3231332201323201-0232132312101311-3230130032312001-1212102113222112-3202002113133022-0032222013230302-2122333100023121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013111201323000-0313012202230032-0320223312102330-2031220232133330-3003321113100221-0231333033100301-1210203303210302-0320021302230201"></a>

## api_rate_limit.server_url_rules.request_matcher.headers.check_not_present — check_not_present / 230330210032 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [api_rate_limit.server_url_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-005.md#canonical-0002122103212131-2020122213023330-0233012102211313-2223031022231311-3110202023102222-1303322023020132-3300313112001132-0021201011102323)
- api_rate_limit.server_url_rules.request_matcher.headers.check_not_present

<a id="canonical-3130030012231213-1303323332300120-2300122330301002-2010300230202330-0310323311223123-1302320002221012-0232112022200321-0303112323320321"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

Upstream description:

This can be used for messages where no values are needed.

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
check_not_present = {}
```

<a id="canonical-2212213323130001-3123010030132203-2123322231110100-3300221201010332-0021220211002220-0212103002320013-1330211223113002-3211211221100230"></a>

## Direct properties — check_not_present / 230330210032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1302120103310023-1210300112030111-2202310302022210-0010001301230011-1201303200222232-2132032322301013-0130030203220102-1013210322323023"></a>

## Next pages — check_not_present / 230330210032 / 4

- [api_rate_limit.server_url_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-005.md#canonical-0002122103212131-2020122213023330-0233012102211313-2223031022231311-3110202023102222-1303322023020132-3300313112001132-0021201011102323)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1331001203220231-1210232321130113-0002023333310220-1120311111210203-3123123120321320-3330133133100323-0233320030311210-1120022231031332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122020110022033-3123210221311100-2312202311013233-0330223121320330-2131001002130232-2013310322130213-0200301300201312-1333113220300112"></a>

## api_rate_limit.server_url_rules.request_matcher.headers.check_present — check_present / 121233101023 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [api_rate_limit.server_url_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-005.md#canonical-0002122103212131-2020122213023330-0233012102211313-2223031022231311-3110202023102222-1303322023020132-3300313112001132-0021201011102323)
- api_rate_limit.server_url_rules.request_matcher.headers.check_present

<a id="canonical-1333313223232333-2220000213213003-3230020333320332-0032312120101311-2333112202110321-0311332131210013-1312221033113332-2301030203231102"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

Upstream description:

This can be used for messages where no values are needed.

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
check_present = {}
```

<a id="canonical-2131020321132221-1322100210023300-1002313331113210-3002023033102233-0131100300003303-1233211013312210-3212213111213101-2032223010312122"></a>

## Direct properties — check_present / 121233101023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0132300323001300-2322321000232121-1223311223310130-3201222213132033-0203221330333032-1312312331030030-1332233121213102-2002101200020032"></a>

## Next pages — check_present / 121233101023 / 4

- [api_rate_limit.server_url_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-005.md#canonical-0002122103212131-2020122213023330-0233012102211313-2223031022231311-3110202023102222-1303322023020132-3300313112001132-0021201011102323)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0221231223321303-3120300201222020-1120200011202210-3211331311120130-0220332033231000-1011003020230110-0311133003003031-3230012101233210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122033113202203-0322100003133121-0220110113301312-3223201111133220-3002302212210233-2022333122233222-3001321202001021-1310331300000120"></a>

## api_rate_limit.server_url_rules.request_matcher.headers.item — item / 100332022320 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [api_rate_limit.server_url_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-005.md#canonical-0002122103212131-2020122213023330-0233012102211313-2223031022231311-3110202023102222-1303322023020132-3300313112001132-0021201011102323)
- api_rate_limit.server_url_rules.request_matcher.headers.item

<a id="canonical-3112002101313321-2112203230130322-1211232321021201-1113113220120331-0233132133031200-2220302302232000-3022111032013232-1313310033221222"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-3313232003311221-0013101301201200-2112302011331201-3121011200102030-2002032212100201-2210320303330011-1132230121031222-0123322322032333"></a>

## Direct properties — item / 100332022320 / 3

<a id="canonical-0021120303301111-3321230013312221-3333002330031320-3000111201110121-1321313123333132-1010123112202003-0330332222111211-2322202310112123"></a>

<a id="canonical-2320120213002102-0123122013010011-1333220032332112-3132103220303003-3202121200022301-1221103102022221-0100023213222033-3113023210213333"></a>

## exact_values property — item / 100332022320 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2232332231032321-3330222201301303-2113202300231301-1222003030321100-1032010130003001-1132112010201302-2021003301010100-1001313231231330"></a>

<a id="canonical-1012221133010131-2110230111113112-3232110312132233-2312033002233133-0303120033033001-1322231213103111-0303030113220311-0010121133101221"></a>

## regex_values property — item / 100332022320 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0203010333012201-1320211020310103-3131012111122033-0202301103113023-1333313220310320-2232312033331313-3122030300301222-1132203021323330"></a>

<a id="canonical-1320233313210320-2012210110000331-2000233013130213-2200132103201230-0200202311102312-2233120220030101-0223120022120031-2302222123002312"></a>

## transformers property — item / 100332022320 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1302202222100031-2221012033111122-0121101100101111-1123220303223222-3013333330120112-0103123121331033-3120220100022112-3221231101311230"></a>

## Next pages — item / 100332022320 / 7

- [api_rate_limit.server_url_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-005.md#canonical-0002122103212131-2020122213023330-0233012102211313-2223031022231311-3110202023102222-1303322023020132-3300313112001132-0021201011102323)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2120002012323320-0210320131230122-1101202032031212-0321112122211211-2002131210323022-3122021323013001-1311113321020303-2332232002032320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3021101030032231-2031103133210100-1300332321211321-0132203022133230-1121320330223332-0203323110212103-1031331302122311-1233133212130031"></a>

## api_rate_limit.server_url_rules.request_matcher.jwt_claims — jwt_claims / 220130110212 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims

<a id="canonical-3212312113023033-2033123110232322-3113330222122021-2133230113312131-2013113100331130-2130132033012311-2111322322322132-3232120322310311"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings.

Upstream description:

A list of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates
must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
jwt_claims {
  # Configure direct properties listed below.
}
```

<a id="canonical-2033212100030022-3102222302123031-3113102201101122-1012002223332231-0320202103102313-1201131333312020-1120300331322210-0303231331023220"></a>

## Direct properties — jwt_claims / 220130110212 / 3

- [check_not_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-1110323000122121-3303301302312310-1320200003111323-3032011220203120-2131033201113102-2203313202211313-1132323001133122-1101331012201031): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-2322331120020121-0002113312012132-0111202000311011-2100002203231213-3200312100012310-3222010320113103-3313302201132000-3321102322211131): complete subsection reference.

<a id="canonical-0030311220322001-2002010011010112-2030220020200032-3322230202000003-1332331110213131-0313201230132032-0000301110202012-0121010112232301"></a>

<a id="canonical-1201320312302021-0130232310200301-0023111001121030-1302033010201033-1303112002131111-0013311133231022-3112131331002000-3012330121311123"></a>

## invert_matcher property — jwt_claims / 220130110212 / 4

Type: `"bool"`. Optional.

Invert Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [item](resources--cdn_loadbalancer--reference--group-005.md#canonical-0003022123333033-0123121313130321-2121132222111112-0130011033022321-2232120020210000-0221331031033030-0233013312030201-2011300220130001): complete subsection reference.

<a id="canonical-3002133203310330-3023102321102112-1113123113233013-0000003012322331-1110032031332201-0303223133111203-3131232101130202-0232121133001222"></a>

<a id="canonical-3230002223121123-1131031330312103-2113321113320020-3331130301200003-2121233302222123-3013012201303222-2312301202101023-1132030020000233"></a>

## name property — jwt_claims / 220130110212 / 5

Type: `"string"`. Optional.

JWT Claim Name. JWT claim name.

Upstream description:

JWT claim name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-3230232323032322-3321012132122130-2312013130221013-3020213032233101-2230300321210132-2233130311331030-1122111221012020-3030223220122323"></a>

## Next pages — jwt_claims / 220130110212 / 6

- [api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-1110323000122121-3303301302312310-1320200003111323-3032011220203120-2131033201113102-2203313202211313-1132323001133122-1101331012201031)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-2322331120020121-0002113312012132-0111202000311011-2100002203231213-3200312100012310-3222010320113103-3313302201132000-3321102322211131)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims.item](resources--cdn_loadbalancer--reference--group-005.md#canonical-0003022123333033-0123121313130321-2121132222111112-0130011033022321-2232120020210000-0221331031033030-0233013312030201-2011300220130001)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1110323000122121-3303301302312310-1320200003111323-3032011220203120-2131033201113102-2203313202211313-1132323001133122-1101331012201031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030320230112020-0212212012010331-2103102223123201-2311131213303200-3120302130132310-0010301103103301-1012312313220020-1033100300102221"></a>

## api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present — check_not_present / 121000332212 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-005.md#canonical-2120002012323320-0210320131230122-1101202032031212-0321112122211211-2002131210323022-3122021323013001-1311113321020303-2332232002032320)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-1211101130033310-1313033332132001-3202330020130010-3210101123230132-0033313020012002-1032011112301230-0310121200201102-1132102321131003"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

Upstream description:

This can be used for messages where no values are needed.

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
check_not_present = {}
```

<a id="canonical-3023022123330200-0110212333332323-1330010220131021-3033032302002203-3021010112022031-0221320203110212-3310233212213022-1230313230030100"></a>

## Direct properties — check_not_present / 121000332212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2030332332132023-2121222010210232-0232321030223211-1222001103121000-0010222002312230-3301013022032333-0023131022021102-0212333311131110"></a>

## Next pages — check_not_present / 121000332212 / 4

- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-005.md#canonical-2120002012323320-0210320131230122-1101202032031212-0321112122211211-2002131210323022-3122021323013001-1311113321020303-2332232002032320)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2322331120020121-0002113312012132-0111202000311011-2100002203231213-3200312100012310-3222010320113103-3313302201132000-3321102322211131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1211003110201311-1123320213131332-1021232231023031-1101230231201033-1312111020333132-1210201203112022-3323022120221010-0002102221231301"></a>

## api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present — check_present / 002010011312 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-005.md#canonical-2120002012323320-0210320131230122-1101202032031212-0321112122211211-2002131210323022-3122021323013001-1311113321020303-2332232002032320)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present

<a id="canonical-2232303200223332-1232123112231210-1222323330102001-0212113332320200-3032302231131122-0011211030023233-2111003302020331-1120010132131022"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

Upstream description:

This can be used for messages where no values are needed.

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
check_present = {}
```

<a id="canonical-2033112031321120-1332332030212132-0133212000120100-0111013332103230-1012122233323033-3131102210331122-3203102002223313-1002310230330012"></a>

## Direct properties — check_present / 002010011312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3322232310012330-2331232011103101-0202223230331332-0031323102011330-3123333032132031-2021112312200113-1200033320303031-1323131132031202"></a>

## Next pages — check_present / 002010011312 / 4

- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-005.md#canonical-2120002012323320-0210320131230122-1101202032031212-0321112122211211-2002131210323022-3122021323013001-1311113321020303-2332232002032320)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0003022123333033-0123121313130321-2121132222111112-0130011033022321-2232120020210000-0221331031033030-0233013312030201-2011300220130001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2223011333321302-1030203122121202-3113102111010211-1030113233201221-0222033321111122-1011101031020210-3200212302132033-2113002101202000"></a>

## api_rate_limit.server_url_rules.request_matcher.jwt_claims.item — item / 112030311033 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-005.md#canonical-2120002012323320-0210320131230122-1101202032031212-0321112122211211-2002131210323022-3122021323013001-1311113321020303-2332232002032320)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims.item

<a id="canonical-2121231321210132-2203000303010210-2020300212022333-1002122211333233-0300123232010313-0311311030013211-3113210023010323-2210111100003300"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-2012000023310003-1003232330021212-3013210310310320-3302011211033211-3013232323122023-3111112033012133-1301301321210303-3110200210133032"></a>

## Direct properties — item / 112030311033 / 3

<a id="canonical-2033301130312200-2332222203233331-1211131212120323-0203021013202221-2222231221100000-3202012200222020-2210101301200303-1022331312110123"></a>

<a id="canonical-2303203011300333-0232111101332020-1122233020331032-1021133232200222-0312233120310132-2203120031031123-2211311131013200-3132333333323133"></a>

## exact_values property — item / 112030311033 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3032112131111232-3000012023033001-1020001232102223-2011010130111322-2031020212102131-1200220221201302-1200000132233102-3310232132002323"></a>

<a id="canonical-1320301200320020-2010000311221030-3320031203232333-2210303012210020-3131111330123103-1103033033021332-0012010313330013-0103232233123231"></a>

## regex_values property — item / 112030311033 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2103020130003303-1310020221031202-3120010313003320-2320323331100123-0120331121203011-3333013312210020-1220202222323323-0031102323033213"></a>

<a id="canonical-3210021323332230-1002030003131221-1200003033202021-0031021203303023-1201221132032321-1233331313003300-1002021120221102-0331110222110020"></a>

## transformers property — item / 112030311033 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0220131031302221-2013000313300310-2122220123131131-0201023012300010-0233132302003010-1312310230323331-0313310302232220-2231130213211110"></a>

## Next pages — item / 112030311033 / 7

- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-005.md#canonical-2120002012323320-0210320131230122-1101202032031212-0321112122211211-2002131210323022-3122021323013001-1311113321020303-2332232002032320)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0122313023112111-2300122102223032-0012011121332013-0133313131303331-3101132233230020-2203113102113133-3002130201213302-3001033321123231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003121030032031-3132021122021303-0323020332113121-2331222033101303-3332012031023103-3100300223322322-0112333201123221-2222121323301003"></a>

## api_rate_limit.server_url_rules.request_matcher.query_params — query_params / 323231310312 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- api_rate_limit.server_url_rules.request_matcher.query_params

<a id="canonical-1231321222222110-0033113223323002-1010111300220300-2311321231203323-0011103330120103-2003220003003202-0132332113321123-0102212330123003"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for all query parameters that need to be matched. The criteria for matching each
query parameter are described in individual instances of QueryParameterMatcherType. The actual query
parameter values are extracted from the request API as a list of strings for each query..

Upstream description:

A list of predicates for all query parameters that need to be matched. The criteria for matching
each query parameter are described in individual instances of QueryParameterMatcherType. The actual
query parameter values are extracted from the request API as a list of strings for each query
parameter name. Note that all specified query parameter predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("key"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
query_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-2312300133230322-2333121021300202-3032123023213122-3121010013112112-0021313012022110-2001012213313121-3200121021323310-1222012100123013"></a>

## Direct properties — query_params / 323231310312 / 3

- [check_not_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-0031211030333021-3121133300032202-2222032100113212-0222322311221311-3200013103011122-2032001331333233-1121130331332110-2130300203113110): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-3312033211310313-3303303032031201-0313212311020003-1002300120123310-3102100123213103-2210322300013230-0023122312323202-2033033022000301): complete subsection reference.

<a id="canonical-3001121131031030-1210110313021110-0001101202111212-2021303200301221-3203132033032103-0122003112113130-3310200121002003-0313310210311130"></a>

<a id="canonical-3103120121132101-2120033123312133-0123222011133223-1013031233220313-1300113003212002-3301211113132232-0233033021222013-1330222202321122"></a>

## invert_matcher property — query_params / 323231310312 / 4

Type: `"bool"`. Optional.

Invert Query Parameter Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [item](resources--cdn_loadbalancer--reference--group-005.md#canonical-3022303020223232-1111323302302211-3023001310232001-1301133223032302-2031221220101030-3011311302103021-3213330300230233-0302101230233121): complete subsection reference.

<a id="canonical-0202211210031202-0333122030200020-1113311312023111-1200320211101313-2220002033232233-3132003303203231-0301102310033003-3302031030100020"></a>

<a id="canonical-1121012113200030-1130200201321121-1130313233120110-3112120001310303-3223300222313123-2322131301321212-1123120332113331-1033321011300221"></a>

## key property — query_params / 323231310312 / 5

Type: `"string"`. Optional.

Case-sensitive HTTP query parameter name.

Upstream description:

A case-sensitive HTTP query parameter name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-1202001201302233-2031320130000130-0333333211303302-2222300103102233-2033220323103032-3121033322030101-1212200023133013-3000230331303200"></a>

## Next pages — query_params / 323231310312 / 6

- [api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-0031211030333021-3121133300032202-2222032100113212-0222322311221311-3200013103011122-2032001331333233-1121130331332110-2130300203113110)
- [api_rate_limit.server_url_rules.request_matcher.query_params.check_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-3312033211310313-3303303032031201-0313212311020003-1002300120123310-3102100123213103-2210322300013230-0023122312323202-2033033022000301)
- [api_rate_limit.server_url_rules.request_matcher.query_params.item](resources--cdn_loadbalancer--reference--group-005.md#canonical-3022303020223232-1111323302302211-3023001310232001-1301133223032302-2031221220101030-3011311302103021-3213330300230233-0302101230233121)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0031211030333021-3121133300032202-2222032100113212-0222322311221311-3200013103011122-2032001331333233-1121130331332110-2130300203113110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3111322011110221-1113223022211002-3332020320031321-0001021132101230-0203312321310110-2301132220303230-1111130021121211-1201322101210320"></a>

## api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present — check_not_present / 031011123131 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [api_rate_limit.server_url_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122313023112111-2300122102223032-0012011121332013-0133313131303331-3101132233230020-2203113102113133-3002130201213302-3001033321123231)
- api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present

<a id="canonical-0012121133222020-1030131122232111-1233302033330021-3230323321023021-0200132320213100-0211123201223001-0330111223232103-0110023212113231"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

Upstream description:

This can be used for messages where no values are needed.

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
check_not_present = {}
```

<a id="canonical-0231023232330211-3232023202300013-1322232211022020-1302210103200103-2202112212232330-1313321312102101-3010310200222310-3311332123110222"></a>

## Direct properties — check_not_present / 031011123131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2302313021010122-2123113101022120-1301313001222330-3330302311333010-0231001223201200-0302123302210203-3311311112012132-3021311130033311"></a>

## Next pages — check_not_present / 031011123131 / 4

- [api_rate_limit.server_url_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122313023112111-2300122102223032-0012011121332013-0133313131303331-3101132233230020-2203113102113133-3002130201213302-3001033321123231)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3312033211310313-3303303032031201-0313212311020003-1002300120123310-3102100123213103-2210322300013230-0023122312323202-2033033022000301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1332131130111321-0120131120120221-0103232131132133-1311110333120111-1123212022031313-0311113013300330-1213011000020332-3201122121213210"></a>

## api_rate_limit.server_url_rules.request_matcher.query_params.check_present — check_present / 313233133122 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [api_rate_limit.server_url_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122313023112111-2300122102223032-0012011121332013-0133313131303331-3101132233230020-2203113102113133-3002130201213302-3001033321123231)
- api_rate_limit.server_url_rules.request_matcher.query_params.check_present

<a id="canonical-0201200311312110-1212232202311321-0212020031330121-3202010300333132-2332021223320111-1022232222013110-3333232210023002-0310332111210302"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

Upstream description:

This can be used for messages where no values are needed.

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
check_present = {}
```

<a id="canonical-1310001023310031-1030330103200202-3113021123022301-3100010300001030-3132020123222220-2100331131232100-3333002130121103-0313231100323232"></a>

## Direct properties — check_present / 313233133122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1010300331111301-0013001131130100-0203122202012232-0233201202301013-1323323311212001-0110303023213000-2203021013231111-2222003312111012"></a>

## Next pages — check_present / 313233133122 / 4

- [api_rate_limit.server_url_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122313023112111-2300122102223032-0012011121332013-0133313131303331-3101132233230020-2203113102113133-3002130201213302-3001033321123231)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3022303020223232-1111323302302211-3023001310232001-1301133223032302-2031221220101030-3011311302103021-3213330300230233-0302101230233121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032203303020112-3032310001220322-2223310303110131-2131033011300111-2031332311033320-2222100123211033-2211033201023102-3210100103230113"></a>

## api_rate_limit.server_url_rules.request_matcher.query_params.item — item / 200121010132 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [api_rate_limit.server_url_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122313023112111-2300122102223032-0012011121332013-0133313131303331-3101132233230020-2203113102113133-3002130201213302-3001033321123231)
- api_rate_limit.server_url_rules.request_matcher.query_params.item

<a id="canonical-1211321221121011-3111113022112222-3321233101000030-2131103233321302-1333200102022230-0203323112310220-2031210012123232-2202132322222232"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-3032313103030201-0213233331103302-2200233300320100-3021132312102122-3003111021031122-0032300103111203-2323031032131023-3132211202130112"></a>

## Direct properties — item / 200121010132 / 3

<a id="canonical-3301123011303113-3021030010102310-2222131101101210-3100001111313100-3301023303130030-3222332201130332-1223022020133320-2112103000233033"></a>

<a id="canonical-1303303232312323-0123032322311001-2133121121201320-3002220203032222-3003102203010310-1021223130020012-1121123221013221-3233130110022320"></a>

## exact_values property — item / 200121010132 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2121221113022030-1101131012330010-0023200330200332-3132312310000023-1221221313113002-0001310132113000-1212111031001031-0120112101012203"></a>

<a id="canonical-3021200011113020-2203310210332311-2001013110213011-0223020230232021-2230323131332210-3313200131310122-1201201330221320-2301210320212101"></a>

## regex_values property — item / 200121010132 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2213203210333323-2130113302001000-0212231201232203-1023011312330321-2130033032002021-0312102332230211-0123131033011121-1120303000103310"></a>

<a id="canonical-0102132000300113-0032330231023121-0330232021201011-1333303002302333-1110002002113033-0311021332310003-0123133013311222-1221130002022212"></a>

## transformers property — item / 200121010132 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1322012120012230-3120010102313131-1213012212031011-3113303211013022-1122032221333302-3122023211211330-3123231002102313-3310020123302121"></a>

## Next pages — item / 200121010132 / 7

- [api_rate_limit.server_url_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122313023112111-2300122102223032-0012011121332013-0133313131303331-3101132233230020-2203113102113133-3002130201213302-3001033321123231)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221000001313202-1033302122333012-2301030213020202-1123103023112120-3030003201212031-2002112201322021-3212231122031111-3211021332111113"></a>

## api_specification — api_specification / 112003322322 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- api_specification

<a id="canonical-0102031021120331-0211222232131320-1022002322130010-0221000333310320-3012323211302110-0213323320200103-0202010323332222-0330001323100020"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: api\_specification, disable\_api\_definition; Default: disable\_api\_definition\] Settings
for API specification (API definition, OpenAPI validation, etc.).

Upstream description:

Settings for API specification (API definition, OpenAPI validation, etc.)

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("validation_all_spec_endpoints",
    "validation_custom_list"),
  validators.ConflictingObjectAttributes("validation_all_spec_endpoints",
    "validation_disabled"),
  validators.ConflictingObjectAttributes("validation_custom_list",
    "validation_disabled")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-validation_target_choice": "[\"validation_all_spec_endpoints\",\"validation_custom_list\",\"validation_disabled\"]"
}
```

OneOf alternatives in this subsection:

- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-0102031021120331-0211222232131320-1022002322130010-0221000333310320-3012323211302110-0213323320200103-0202010323332222-0330001323100020)
- [disable_api_definition](resources--cdn_loadbalancer--reference--group-010.md#canonical-1013231301331321-3201321033031021-0303223333210301-2112110233330302-0302100120123301-0132331033332233-0312121102001302-1222030220330231)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
api_specification {
  # Configure direct properties listed below.
}
```

<a id="canonical-1121333230123010-3232012302103032-1110320311231002-2100010020131213-0011001323301021-1131111221230210-1111112233233312-1011121231333100"></a>

## Direct properties — api_specification / 112003322322 / 3

- [api_definition](resources--cdn_loadbalancer--reference--group-005.md#canonical-3332331233113011-1302210123002122-1102002302230210-0222311320311221-0202010322311020-0233111223111000-2110301333002222-0313013021210020): complete subsection reference.

- [validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333): complete subsection reference.

- [validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133): complete subsection reference.

- [validation_disabled](resources--cdn_loadbalancer--reference--group-007.md#canonical-3201130202332300-2123131032222110-0303233101102201-2330112312130010-0032100122031021-0111232221032121-0030201333002210-1203121111323232): complete subsection reference.

<a id="canonical-2122222033321022-1233201031000313-2221202322330022-0221301331112202-2221303001332321-3222310013012320-0231022222211113-2313011232310321"></a>

## Next pages — api_specification / 112003322322 / 4

- [api_specification.api_definition](resources--cdn_loadbalancer--reference--group-005.md#canonical-3332331233113011-1302210123002122-1102002302230210-0222311320311221-0202010322311020-0233111223111000-2110301333002222-0313013021210020)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_disabled](resources--cdn_loadbalancer--reference--group-007.md#canonical-3201130202332300-2123131032222110-0303233101102201-2330112312130010-0032100122031021-0111232221032121-0030201333002210-1203121111323232)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3332331233113011-1302210123002122-1102002302230210-0222311320311221-0202010322311020-0233111223111000-2110301333002222-0313013021210020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0333303212002302-2321203231022011-1103210012011322-1031220202313130-0323031022010000-3133330311121302-2313130300213303-0231021301121222"></a>

## api_specification.api_definition — api_definition / 133313201110 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- api_specification.api_definition

<a id="canonical-0132023233102013-3301203333023110-1023321032301230-3033120000333332-1212220201310133-0330322333220231-3223221302131210-1230320020313321"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
api_definition {
  # Configure direct properties listed below.
}
```

<a id="canonical-0121123211310200-0032110233021313-3231230203210233-2232210300121110-3012033312031210-0220233312031211-3102333210110320-3122230101011302"></a>

## Direct properties — api_definition / 133313201110 / 3

<a id="canonical-3320110100101323-1333110203233132-2023010323210303-3332210011021132-2023002101212022-3203321321002312-0033101122123203-0222231330103201"></a>

<a id="canonical-2300333111002022-3031012012121213-1103011223203023-3302232323001223-0012301031223120-1311101022113030-3113223330321330-3121211322233332"></a>

## name property — api_definition / 133313201110 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-0112011330312200-0233130300332031-3100032113100212-2332031003303023-1133330300331033-2210100231002000-3120213310303122-2103213313111303"></a>

<a id="canonical-0320231023121331-1030030122211103-1022220230101333-1232220213011113-2123111322011320-1303300330303330-0312301103130212-3332130020022002"></a>

## namespace property — api_definition / 133313201110 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0121331331012103-3133202301030032-2332002201012132-0221311210311330-0301020223230311-2112310202032322-1222203222003110-2112312010201300"></a>

<a id="canonical-1122011020222232-3323001112212020-3222100321103221-3110112321023033-1232320120030000-0211331002133001-2212332020132202-0302032103023233"></a>

## tenant property — api_definition / 133313201110 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2231223212123333-3222022100020023-1021320213123111-0012030010331223-1010203030303030-2323322111232213-1332221201200021-2130102032321123"></a>

## Next pages — api_definition / 133313201110 / 7

- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203233100311121-2013310310323311-2023332231011021-1012231120021033-0122031131310332-3130200323332222-0301203322231332-1313112012121122"></a>

## api_specification.validation_all_spec_endpoints — validation_all_spec_endpoints / 311322330022 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- api_specification.validation_all_spec_endpoints

<a id="canonical-2202223010323020-2322320100113123-0021130200300331-0112323201310023-2230200100311001-2221311213022211-1222302010231032-0130202221030121"></a>

Type: `"object"`. single nested block, Optional.

API Inventory. Settings for API Inventory validation.

Upstream description:

Settings for API Inventory validation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-oversized_body_choice": "[]"
}
```

Terraform syntax:

```terraform
validation_all_spec_endpoints {
  # Configure direct properties listed below.
}
```

<a id="canonical-0103221013300020-2100000013202121-0120012212111201-0201021023211320-0313012130010210-2133320300020010-1113212221111122-2132112001232200"></a>

## Direct properties — validation_all_spec_endpoints / 311322330022 / 3

- [fall_through_mode](resources--cdn_loadbalancer--reference--group-005.md#canonical-1132202011301023-1311003221212003-0223220113211003-0022210103300313-0200312302211323-3013210300201133-2131013023312002-3210220003000332): complete subsection reference.

- [settings](resources--cdn_loadbalancer--reference--group-006.md#canonical-2212122210213200-2011131031130121-1310310331303200-0003121202031101-2212233312312021-1333022320112232-0110003003230311-3212032312203310): complete subsection reference.

- [validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-0032000232202023-1021303223203232-2311231222130001-0301110330022200-1322133020231031-0122001231011131-3132131303303103-1013212230000131): complete subsection reference.

<a id="canonical-2321133012022322-1001313323321103-3031330100113200-0132032223310213-2313010313033031-1213001323321103-3021020012320011-0101201230232331"></a>

## Next pages — validation_all_spec_endpoints / 311322330022 / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--cdn_loadbalancer--reference--group-005.md#canonical-1132202011301023-1311003221212003-0223220113211003-0022210103300313-0200312302211323-3013210300201133-2131013023312002-3210220003000332)
- [api_specification.validation_all_spec_endpoints.settings](resources--cdn_loadbalancer--reference--group-006.md#canonical-2212122210213200-2011131031130121-1310310331303200-0003121202031101-2212233312312021-1333022320112232-0110003003230311-3212032312203310)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-0032000232202023-1021303223203232-2311231222130001-0301110330022200-1322133020231031-0122001231011131-3132131303303103-1013212230000131)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1132202011301023-1311003221212003-0223220113211003-0022210103300313-0200312302211323-3013210300201133-2131013023312002-3210220003000332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2101232011322332-3001313010030122-2202002233012131-1112023231112123-1221231233001313-1100001120020311-0331110330112321-3231213002323312"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode — fall_through_mode / 213032112121 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- api_specification.validation_all_spec_endpoints.fall_through_mode

<a id="canonical-3010321102320013-1120320322131003-3023111131303010-1333122101011123-3220112213113121-3032131133021302-1333310031023031-1110233200122231"></a>

Type: `"object"`. single nested block, Optional.

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules).

Upstream description:

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules)

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("fall_through_mode_allow",
    "fall_through_mode_custom")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-fall_through_mode_choice": "[\"fall_through_mode_allow\",\"fall_through_mode_custom\"]"
}
```

Terraform syntax:

```terraform
fall_through_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-1103133320201013-0003113312122131-0002002233233013-1032202222130103-2320332123000021-3330103212303003-1220132020110230-0333310021033011"></a>

## Direct properties — fall_through_mode / 213032112121 / 3

- [fall_through_mode_allow](resources--cdn_loadbalancer--reference--group-005.md#canonical-1112311230100333-2313002332302102-0031001300211321-2232211332010232-3221220002131020-2023103121122032-0200131013302101-2233320323300012): complete subsection reference.

- [fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-005.md#canonical-1313112202023113-1100031331202211-3202200203100020-1313201021000212-3033331010021322-2232130310013013-0313201103013202-3001021303232013): complete subsection reference.

<a id="canonical-1233123120113100-0022100122223012-1011203332120110-3321200000310210-2330300013133330-3323302201321002-3321222013130330-0222310010211010"></a>

## Next pages — fall_through_mode / 213032112121 / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_allow](resources--cdn_loadbalancer--reference--group-005.md#canonical-1112311230100333-2313002332302102-0031001300211321-2232211332010232-3221220002131020-2023103121122032-0200131013302101-2233320323300012)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-005.md#canonical-1313112202023113-1100031331202211-3202200203100020-1313201021000212-3033331010021322-2232130310013013-0313201103013202-3001021303232013)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1112311230100333-2313002332302102-0031001300211321-2232211332010232-3221220002131020-2023103121122032-0200131013302101-2233320323300012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0223012311110113-0320031121133313-1112220213033330-3220131321030030-0313001021232000-0012322331321213-0200313121212112-2203020112130012"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_allow — fall_through_mode_allow / 302311323321 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--cdn_loadbalancer--reference--group-005.md#canonical-1132202011301023-1311003221212003-0223220113211003-0022210103300313-0200312302211323-3013210300201133-2131013023312002-3210220003000332)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_allow

<a id="canonical-3110330022103213-1000221320230102-0221232120133033-1301213221021012-1032301223230231-2200321021032033-0100303102200230-1102032022123302"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for fall through mode allow.

Upstream description:

This can be used for messages where no values are needed.

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
fall_through_mode_allow = {}
```

<a id="canonical-1013211311121100-1112011220312122-0020211130212001-3021002020103102-0333211211130021-3002033003102100-2031121100312222-3203033102032023"></a>

## Direct properties — fall_through_mode_allow / 302311323321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1211001313121233-3213032330211131-0333102332120110-3123012230112222-3102210011031223-2133233210023133-2223331223213013-0113010133321110"></a>

## Next pages — fall_through_mode_allow / 302311323321 / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--cdn_loadbalancer--reference--group-005.md#canonical-1132202011301023-1311003221212003-0223220113211003-0022210103300313-0200312302211323-3013210300201133-2131013023312002-3210220003000332)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1313112202023113-1100031331202211-3202200203100020-1313201021000212-3033331010021322-2232130310013013-0313201103013202-3001021303232013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013121210321003-2201031333030003-0020322211232310-3333220010012012-0101223111321210-3012322132300331-3130232221330330-3122111100232121"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom — fall_through_mode_custom / 120022110101 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--cdn_loadbalancer--reference--group-005.md#canonical-1132202011301023-1311003221212003-0223220113211003-0022210103300313-0200312302211323-3013210300201133-2131013023312002-3210220003000332)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom

<a id="canonical-1100000102332202-0023301233232201-1012001213320022-2030103333231302-0203130122320031-3310110303021301-3123121010222210-0012133013332111"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for fall through mode custom.

Upstream description:

Define the fall through settings.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("open_api_validation_rules")}
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
fall_through_mode_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-0301220132233222-1321132011110123-0113300300230332-2100311022111312-3230022021310123-1033021213031321-3332102010033120-2332100131333132"></a>

## Direct properties — fall_through_mode_custom / 120022110101 / 3

- [open_api_validation_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0123300033001212-0103231020012121-0122330200112313-1113202010311313-0313101101311223-0121320230132013-3110202300123132-3323213313332100): complete subsection reference.

<a id="canonical-3101011213022020-3300223221130201-0103233210200111-2102300022331232-1231030133320122-0122021231323230-0123010203310330-2001231003133213"></a>

## Next pages — fall_through_mode_custom / 120022110101 / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0123300033001212-0103231020012121-0122330200112313-1113202010311313-0313101101311223-0121320230132013-3110202300123132-3323213313332100)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--cdn_loadbalancer--reference--group-005.md#canonical-1132202011301023-1311003221212003-0223220113211003-0022210103300313-0200312302211323-3013210300201133-2131013023312002-3210220003000332)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0123300033001212-0103231020012121-0122330200112313-1113202010311313-0313101101311223-0121320230132013-3110202300123132-3323213313332100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212020030131210-0120332123112303-2021023012301033-3032033200130211-2310202233103133-0103031213312103-1112113021223000-0201332021230131"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules — open_api_validation_rules / 121020312123 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--cdn_loadbalancer--reference--group-005.md#canonical-1132202011301023-1311003221212003-0223220113211003-0022210103300313-0200312302211323-3013210300201133-2131013023312002-3210220003000332)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-005.md#canonical-1313112202023113-1100031331202211-3202200203100020-1313201021000212-3033331010021322-2232130310013013-0313201103013202-3001021303232013)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules

<a id="canonical-0110010023210300-0103132130022200-1210332122023122-3231223000021230-1021130322203310-3131021021011003-2202332301223323-0301113330121123"></a>

Type: `"object"`. list nested block, Optional.

Custom Fall Through Rule List. Rule or policy definition

Upstream description:

Rule or policy definition

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("action_block",
    "action_report"),
  validators.ConflictingListObjectAttributes("action_block",
    "action_skip"),
  validators.ConflictingListObjectAttributes("action_report",
    "action_skip"),
  validators.ConflictingListObjectAttributes("api_endpoint",
    "api_group"),
  validators.ConflictingListObjectAttributes("api_endpoint",
    "base_path"),
  validators.ConflictingListObjectAttributes("api_group",
    "base_path")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 15,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 15,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
open_api_validation_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-0101130230022032-3110100322021102-1003213213132030-1113012320121031-3003232333231031-3200002201001200-0200110102223103-1323221313100202"></a>

## Direct properties — open_api_validation_rules / 121020312123 / 3

- [action_block](resources--cdn_loadbalancer--reference--group-005.md#canonical-3022021220103010-3102230030122201-0202213022330321-0022123120302221-0332000121111211-0100310301330132-2323131231202100-2023300332012223): complete subsection reference.

- [action_report](resources--cdn_loadbalancer--reference--group-005.md#canonical-2132011323120203-2013132122023310-0133011203202102-1011022211221121-3121031311220000-0310203001323302-1130201120032311-2200103301011121): complete subsection reference.

- [action_skip](resources--cdn_loadbalancer--reference--group-005.md#canonical-0133032033320222-3111131132223303-0321131322312202-0013130131231010-2013103300303010-2321321013031322-1203321123222202-2202121220201232): complete subsection reference.

- [api_endpoint](resources--cdn_loadbalancer--reference--group-006.md#canonical-1103111311121333-0113333123303023-1322223010120310-3112011320132332-2222310112112311-0233131102323212-0321130302221232-2011300231311323): complete subsection reference.

<a id="canonical-2311222322200332-0303101101122110-1203120303101303-0010202031031021-0213001020213310-0033003301212130-2132101012302032-3130211202100110"></a>

<a id="canonical-3233010012121132-1011113333222031-3022320023330333-1112011213302213-0000103300011202-1322312230003303-0012320120031030-0332300130103303"></a>

## api_group property — open_api_validation_rules / 121020312123 / 4

Type: `"string"`. Optional.

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-0021010110031132-2301313010310103-2223223320210030-2211113213120102-2021231102032030-0322321310212130-0332202231021223-1102212330102021"></a>

<a id="canonical-3103333323303030-0221132321313203-1022133000003233-0320121312322201-2013221330122000-2113102223303020-3020302110113310-3312231120022213"></a>

## base_path property — open_api_validation_rules / 121020312123 / 5

Type: `"string"`. Optional.

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [metadata](resources--cdn_loadbalancer--reference--group-006.md#canonical-3122022232220301-3022311110311231-2132233022101212-1033021212201102-2302313211003333-2321332322111311-2313032223301201-3112030031020321): complete subsection reference.

<a id="canonical-2101222300101301-3220321002313000-0112111201321312-0233102033132021-1331231111121123-3310110321210133-0030130013223212-1221012212000313"></a>

## Next pages — open_api_validation_rules / 121020312123 / 6

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block](resources--cdn_loadbalancer--reference--group-005.md#canonical-3022021220103010-3102230030122201-0202213022330321-0022123120302221-0332000121111211-0100310301330132-2323131231202100-2023300332012223)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report](resources--cdn_loadbalancer--reference--group-005.md#canonical-2132011323120203-2013132122023310-0133011203202102-1011022211221121-3121031311220000-0310203001323302-1130201120032311-2200103301011121)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip](resources--cdn_loadbalancer--reference--group-005.md#canonical-0133032033320222-3111131132223303-0321131322312202-0013130131231010-2013103300303010-2321321013031322-1203321123222202-2202121220201232)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint](resources--cdn_loadbalancer--reference--group-006.md#canonical-1103111311121333-0113333123303023-1322223010120310-3112011320132332-2222310112112311-0233131102323212-0321130302221232-2011300231311323)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata](resources--cdn_loadbalancer--reference--group-006.md#canonical-3122022232220301-3022311110311231-2132233022101212-1033021212201102-2302313211003333-2321332322111311-2313032223301201-3112030031020321)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-005.md#canonical-1313112202023113-1100031331202211-3202200203100020-1313201021000212-3033331010021322-2232130310013013-0313201103013202-3001021303232013)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3022021220103010-3102230030122201-0202213022330321-0022123120302221-0332000121111211-0100310301330132-2323131231202100-2023300332012223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0132231000122302-2112112002121300-2011320223133133-2203131212113132-2123022123230132-0203230131211202-1200300311011300-1121012202312200"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block — action_block / 322102113023 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--cdn_loadbalancer--reference--group-005.md#canonical-1132202011301023-1311003221212003-0223220113211003-0022210103300313-0200312302211323-3013210300201133-2131013023312002-3210220003000332)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-005.md#canonical-1313112202023113-1100031331202211-3202200203100020-1313201021000212-3033331010021322-2232130310013013-0313201103013202-3001021303232013)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0123300033001212-0103231020012121-0122330200112313-1113202010311313-0313101101311223-0121320230132013-3110202300123132-3323213313332100)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block

<a id="canonical-1020030013131122-3011313121220303-0023112320321223-1000302330011011-2031022010133111-1002323231113231-0120131313032203-2122201020130031"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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
action_block = {}
```

<a id="canonical-3123003000323312-0013121122133031-3332002313301030-0320130333321212-2012112220123012-3101100211120202-0113011113313113-2111213102121112"></a>

## Direct properties — action_block / 322102113023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2220313030113130-3223231022003313-0323032301321233-2121310010321102-1021120132313110-1102301300302130-1111222332002123-1021112013220131"></a>

## Next pages — action_block / 322102113023 / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0123300033001212-0103231020012121-0122330200112313-1113202010311313-0313101101311223-0121320230132013-3110202300123132-3323213313332100)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2132011323120203-2013132122023310-0133011203202102-1011022211221121-3121031311220000-0310203001323302-1130201120032311-2200103301011121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132203333011013-1020110123102323-1213332203101130-0012030002123012-2303113223332132-1311111332031330-2223322123013203-0202302320023332"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report — action_report / 131323200011 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--cdn_loadbalancer--reference--group-005.md#canonical-1132202011301023-1311003221212003-0223220113211003-0022210103300313-0200312302211323-3013210300201133-2131013023312002-3210220003000332)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-005.md#canonical-1313112202023113-1100031331202211-3202200203100020-1313201021000212-3033331010021322-2232130310013013-0313201103013202-3001021303232013)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0123300033001212-0103231020012121-0122330200112313-1113202010311313-0313101101311223-0121320230132013-3110202300123132-3323213313332100)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report

<a id="canonical-3021302223123330-3323011220130231-1211223330310120-0330110232001030-3112210032321003-2030211132120232-2032221132002100-3132010211311111"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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
action_report = {}
```

<a id="canonical-3002221222302312-1200101101322013-0113323310032102-3020101210022202-0221222321022032-1033021012022221-0133131320211123-1232120330130100"></a>

## Direct properties — action_report / 131323200011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2033101212300330-1033110322201020-1200101003001331-2022131231221130-0002311313013021-0003320223131102-0203311012012031-0331312103022112"></a>

## Next pages — action_report / 131323200011 / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0123300033001212-0103231020012121-0122330200112313-1113202010311313-0313101101311223-0121320230132013-3110202300123132-3323213313332100)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0133032033320222-3111131132223303-0321131322312202-0013130131231010-2013103300303010-2321321013031322-1203321123222202-2202121220201232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210221132022132-0031011122210301-0001121222012203-1202231303330123-2133002123230322-0123230331231133-1210012001130000-2331320121203332"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip — action_skip / 221200020002 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--cdn_loadbalancer--reference--group-005.md#canonical-1132202011301023-1311003221212003-0223220113211003-0022210103300313-0200312302211323-3013210300201133-2131013023312002-3210220003000332)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-005.md#canonical-1313112202023113-1100031331202211-3202200203100020-1313201021000212-3033331010021322-2232130310013013-0313201103013202-3001021303232013)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0123300033001212-0103231020012121-0122330200112313-1113202010311313-0313101101311223-0121320230132013-3110202300123132-3323213313332100)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip

<a id="canonical-0013222000310230-3002102300221210-0002121321021101-0030311301303132-0133010001231201-2101233303022002-2123122220232331-0333032103021123"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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
action_skip = {}
```

<a id="canonical-1330213101131003-3231322200030132-3100023010100332-1110230201001011-2102033210101222-1300020200221211-0231111010332332-0002222223132113"></a>

## Direct properties — action_skip / 221200020002 / 3

This is an empty object or choice marker. It has no direct properties.
