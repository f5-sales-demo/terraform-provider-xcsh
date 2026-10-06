---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-2130111131232013-1220123201311202-3133000103231121-2322113333330323-0202110311211331-3113022121103001-1301223132322213-2023303310002223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.ip_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311)
- api_rate_limit.server_url_rules.client_matcher.ip_matcher

<a id="canonical-1133122320001010-0313103132200101-3032123312232212-1333331321120222-3002030031001001-2322200013010013-3133033312223001-3101030212321101"></a>

Type: `"object"`. single nested block, Optional.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3323210331221133-1101331223210031-1121220020230021-1100031013301211-3111212132201003-2133201031302112-1030223010202012-2301111001010021"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.ip_matcher`

<a id="canonical-0022201132310221-3313123210033332-2033022113023313-2112300201203312-2210322133030112-1010310030122100-0110311013232332-0101333201100003"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_matcher.invert_matcher` property

Type: `"bool"`. Optional.

Invert IP Matcher. Invert the match result.

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

- [prefix_sets](resources--cdn_loadbalancer--reference--group-006.md#canonical-3313002220000222-1013023302210203-0331011102101332-1321220011231113-1201123222220003-0201321312303122-3000133220202331-0221333322011123): complete subsection reference.

<a id="canonical-3313002220000222-1013023302210203-0331011102101332-1321220011231113-1201123222220003-0201321312303122-3000133220202331-0221333322011123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311)
- [api_rate_limit.server_url_rules.client_matcher.ip_matcher](resources--cdn_loadbalancer--reference--group-006.md#canonical-2130111131232013-1220123201311202-3133000103231121-2322113333330323-0202110311211331-3113022121103001-1301223132322213-2023303310002223)
- api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-0101031020122023-1110211033210023-1110100213222020-0120230323100003-1110010332130031-1103232221320222-1232221333102133-1111131212001023"></a>

Type: `"object"`. list nested block, Optional.

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

<a id="canonical-1133022102000122-0333103331323032-2130303321010330-0130201223023121-0333101132001020-0222010312221001-1333010123222100-3010111322223001"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets`

<a id="canonical-2233301113300221-2100122220032130-1112032112330033-2213003312310012-3110023012330330-1102210002202233-3202002113022323-2211212032131020"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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

<a id="canonical-2303313220202022-0303230112112012-0210303231021210-0302313313220123-0132223020001312-3321030000103331-0120112110122203-3302010000102030"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets.name` property

Type: `"string"`. Optional.

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

<a id="canonical-2320033030212331-3031122120320323-2113333332222111-0201313112320203-3302323013022022-1013322010120231-1323003201131213-3231232232002323"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2311313313232320-2232233221332023-2320220132212233-2001300011003123-3112300113123220-0002121030312313-3131130202033133-0113112120221301"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets.tenant` property

Type: `"string"`. Computed.

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

<a id="canonical-2200033022132330-1013032221120002-1320220311223212-1023032300110220-3223032203020123-3021113222121221-3232023231110332-3012122013032130"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets.uid` property

Type: `"string"`. Computed.

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

<a id="canonical-1121321233302210-2201233030311302-0233012110112111-3301022303130020-3330230211332201-1122013113002133-3100021113102003-1302122332233003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.ip_prefix_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
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

<a id="canonical-1120110333210222-0332210123320230-3313311113311032-2310330332130220-1121201310212101-0201003221222331-0210123120002311-3323133312111321"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.ip_prefix_list`

<a id="canonical-3113331313220101-2331330033023320-3010103032031011-1332002231302203-3011330030323102-3003020013232031-2031300130020310-0001130101320101"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_prefix_list.invert_match` property

Type: `"bool"`. Optional.

Invert Match Result. Invert the match result.

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

<a id="canonical-1020221202303220-1232322313331320-1130110001130121-3321323203120203-0013212023020333-2010303322223102-3302122103313230-1000133213020202"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_prefix_list.ip_prefixes` property

Type: `["list", "string"]`. Optional.

IPv4 Prefix List. List of IPv4 prefix strings.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3000100122201321-1031000322331033-0320100200312313-0310123303011001-0231212022301023-2030131301002101-0231000210030100-2010003030323113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311)
- api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list

<a id="canonical-3100330322300200-1331102120020231-0220330232301003-0313300300130023-2212200103120302-3001021301231221-2212230110112332-0310333103203003"></a>

Type: `"object"`. single nested block, Optional.

IP Threat Category List Type. List of IP threat categories.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3220130222122332-2103321002303133-3102111102203210-3213322010330120-1023021230111320-0201130121022232-0101312231010310-1302303333332233"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list`

<a id="canonical-2100202112303322-1022020213112022-1003103301132313-0331310102130231-2202133110110223-3001310110321032-2132011331211220-1010303103330232"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list.ip_threat_categories` property

Type: `["list", "string"]`. Optional.

\[Enum:
SPAM\_SOURCES|WINDOWS\_EXPLOITS|WEB\_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE\_THREATS|TOR\_PROXY|DENIAL\_OF\_SERVICE|NETWORK\]
The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions. Possible values are \`SPAM\_SOURCES\`, \`WINDOWS\_EXPLOITS\`,
\`WEB\_ATTACKS\`, \`BOTNETS\`, \`SCANNERS\`, \`REPUTATION\`, \`PHISHING\`, \`PROXY\`,
\`MOBILE\_THREATS\`, \`TOR\_PROXY\`, \`DENIAL\_OF\_SERVICE\`, \`NETWORK\`. Defaults to
\`SPAM\_SOURCES\`.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0130003213232033-2202300221200002-2331102232130333-0210030322312120-3203313330213022-0310130203033112-1123331112033221-0112233032020000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311)
- api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-1310201021312320-0113202323013002-3130321313012030-3301211231221331-2003311330120133-3222222113203213-3321210102021302-3301011103013002"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-2213212211312012-0020312302200302-1030311201220011-0011230213300200-1101222223133120-3223233321232132-3000310103332010-1102212223002230"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher`

<a id="canonical-3203012101201101-3311213222002223-2332220001020010-1303122223120133-2033032101020032-0223030310131232-3113030310021210-0323000112311123"></a>

#### `api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher.classes` property

Type: `["list", "string"]`. Optional.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Additional upstream details:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0003022033320311-2113331221103201-2131012003021012-2122130011132211-3103001231201220-1313201033111021-0320303130011032-0100330131210120"></a>

#### `api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0333133110210031-3032123002333010-3211202122132032-0231313230300200-2101002110332300-0010131022300310-0021231111203133-0220002030102101"></a>

#### `api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher.excluded_values` property

Type: `["list", "string"]`. Optional.

A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can
be used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0321031211113300-3020232113001120-2211311332112302-3102201332200033-3301102331030312-2110201132013202-3010222101121131-0103331231111021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.inline_rate_limiter` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- api_rate_limit.server_url_rules.inline_rate_limiter

<a id="canonical-2112001120300200-0213110213211102-2232121213220101-3130111122212112-3021133133212111-0123323133032120-0120122200022121-0321200311032112"></a>

Type: `"object"`. single nested block, Optional.

Inline rate-limiter settings for this domain, base-path, or endpoint rule. Select this field as the
required rate\_limiter\_choice when no stored rate-limiter object is used.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0222320320100022-0220311000321130-0312213233102103-0212200210000111-3223030221120202-1121301101013222-0133221023213210-2220322323110001"></a>

### Direct properties for `api_rate_limit.server_url_rules.inline_rate_limiter`

- [ref_user_id](resources--cdn_loadbalancer--reference--group-006.md#canonical-0013311012330202-0121212300112232-0023221120001330-0102000030031120-3321233120230310-2003230003222132-1020133112230313-2220300233121113): complete subsection reference.

<a id="canonical-3101110210003223-0233211032303021-0013321113303300-1202210220023000-2000302331313302-2332331311032323-2003211210113012-1103230101010110"></a>

<a id="canonical-2222103100102322-2103010311230321-2110311203131231-0111310213030313-3232013030202223-2102032102023310-1230212232121023-1233303103232023"></a>

#### `api_rate_limit.server_url_rules.inline_rate_limiter.threshold` property

Type: `"number"`. Optional.

The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified
period.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0310300201010210-1022130203000310-0021131030231313-2231013210232331-3031023320033201-0033313023332003-3322212113332112-1011320320322021"></a>

#### `api_rate_limit.server_url_rules.inline_rate_limiter.unit` property

Type: `"string"`. Optional.

\[Enum: SECOND|MINUTE|HOUR\] Unit for the period per which the rate limit is applied. - SECOND:
Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR:
Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days. Possible values are
\`SECOND\`, \`MINUTE\`, \`HOUR\`. Defaults to \`SECOND\`.

Provider validators and defaults (from schema source):

```go
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

- [use_http_lb_user_id](resources--cdn_loadbalancer--reference--group-006.md#canonical-2230212032311201-2231312302320020-3212020200322021-0003131301030311-3323121300123202-3022102231011330-1031030320220021-1212332030021031): complete subsection reference.

<a id="canonical-0013311012330202-0121212300112232-0023221120001330-0102000030031120-3321233120230310-2003230003222132-1020133112230313-2220300233121113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.inline_rate_limiter](resources--cdn_loadbalancer--reference--group-006.md#canonical-0321031211113300-3020232113001120-2211311332112302-3102201332200033-3301102331030312-2110201132013202-3010222101121131-0103331231111021)
- api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id

<a id="canonical-2221100320102101-2332302121002013-2031011233230131-0210222330203303-3102102103311033-2010321020322101-3311332322002220-0020101222220000"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1212332111133030-3021110311130222-0331031200303010-3320032111213320-3003210220321123-1313113120133210-2333133201021033-1222030203033223"></a>

### Direct properties for `api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id`

<a id="canonical-1331313131200122-3120223221230121-3023303230012233-3233113332331322-2120032113321311-3313232331031220-0221110330231133-0233131011301210"></a>

#### `api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1123333123011200-3200011301110112-1033002002333103-2201212221220313-2112112200322320-0102231231013123-3132103313310122-2300330203321122"></a>

#### `api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3320323122221231-2313112122010231-1300220332033003-2022212210101223-3030000130310331-3313220012103231-1111123131221322-1131302003103233"></a>

#### `api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2230212032311201-2231312302320020-3212020200322021-0003131301030311-3323121300123202-3022102231011330-1031030320220021-1212332030021031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.inline_rate_limiter](resources--cdn_loadbalancer--reference--group-006.md#canonical-0321031211113300-3020232113001120-2211311332112302-3102201332200033-3301102331030312-2110201132013202-3010222101121131-0103331231111021)
- api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id

<a id="canonical-1123112033221002-0231000020012233-1321001311200030-1312102331121120-0211303221233112-0310011012131132-1000022312010311-3031223201010331"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1210111123202013-2212101212113221-1330130203301223-2331201130033000-0332303011120301-2031002033003132-3203032332303101-2213233012110233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.ref_rate_limiter` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- api_rate_limit.server_url_rules.ref_rate_limiter

<a id="canonical-0133331221033120-3233110330132200-3212312233133233-0010301332311201-2202031302133023-3321323221000000-3303321133030130-2312121000012211"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Additional upstream details:

Reference to a stored rate-limiter object for this scoped rule. Select exactly one of
ref\_rate\_limiter and inline\_rate\_limiter.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3101322130203111-1310010102312003-0210003033101212-0201020011231100-0020133301232101-0330201321312012-0231200313203301-0200110011320311"></a>

### Direct properties for `api_rate_limit.server_url_rules.ref_rate_limiter`

<a id="canonical-3220303000310130-2113300121023031-3021011010133132-1201002210112230-2331033030131022-3331103110210123-3110002333003320-1322112103123133"></a>

#### `api_rate_limit.server_url_rules.ref_rate_limiter.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2332101110033203-1300120231112230-0031201031030212-1213033110121001-2100201201103231-3202111100332123-2013222200232321-0230311102032120"></a>

#### `api_rate_limit.server_url_rules.ref_rate_limiter.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3312001211131301-1303011132203331-1323231023321303-2332230303103333-1230233322021033-2232301030212121-1202022133222311-3001032113322022"></a>

#### `api_rate_limit.server_url_rules.ref_rate_limiter.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- api_rate_limit.server_url_rules.request_matcher

<a id="canonical-0011033133021301-2313223002222133-2222211103133300-2201212331231002-0020103000321310-1221311022111220-0310300313120013-2000012312013102"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for request matcher.

Additional upstream details:

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

<a id="canonical-3212010010100023-0332312230211111-0332132021313201-1102330202122113-0210233132233203-3011103220223213-3231301212010300-3131130003222313"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher`

- [cookie_matchers](resources--cdn_loadbalancer--reference--group-006.md#canonical-3033211112003230-2302031233322223-3011210231211232-3130203102313021-2123313232211002-0033322022102113-0323310000003111-2330301133031030): complete subsection reference.

- [headers](resources--cdn_loadbalancer--reference--group-006.md#canonical-0002122103212131-2020122213023330-0233012102211313-2223031022231311-3110202023102222-1303322023020132-3300313112001132-0021201011102323): complete subsection reference.

- [jwt_claims](resources--cdn_loadbalancer--reference--group-006.md#canonical-2120002012323320-0210320131230122-1101202032031212-0321112122211211-2002131210323022-3122021323013001-1311113321020303-2332232002032320): complete subsection reference.

- [query_params](resources--cdn_loadbalancer--reference--group-006.md#canonical-0122313023112111-2300122102223032-0012011121332013-0133313131303331-3101132233230020-2203113102113133-3002130201213302-3001033321123231): complete subsection reference.

<a id="canonical-3033211112003230-2302031233322223-3011210231211232-3130203102313021-2123313232211002-0033322022102113-0323310000003111-2330301133031030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.cookie_matchers` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-006.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers

<a id="canonical-0101000130230322-2011232033211002-0003000213202020-2312130230100211-2213301330013323-2123301023312131-3202301332100011-0033010020003330"></a>

Type: `"object"`. list nested block, Optional.

A list of predicates for all cookies that need to be matched. The criteria for matching each cookie
is described in individual instances of CookieMatcherType. The actual cookie values are extracted
from the request API as a list of strings for each cookie name. Note that all specified cookie
matcher predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1030222002331123-1221330333200122-2302331232132210-0032032213222000-2210323332023203-3211300100212003-2311021302113210-0323303203123102"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.cookie_matchers`

- [check_not_present](resources--cdn_loadbalancer--reference--group-006.md#canonical-0022023013102010-2123110121321310-0012323310123032-1133003210323313-2002012113210013-1103120130131013-1022022322311101-0313121302122002): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-006.md#canonical-1120300322323132-1200130313310131-0002203230003120-3003321321113332-2222331133131013-3230333133133113-0122033123232211-2031231022331323): complete subsection reference.

<a id="canonical-1112200200312101-1110130302231220-2222323232222101-2303032302221021-3333033210122133-2022320001122201-1210121032122131-2113201002222102"></a>

<a id="canonical-1223322231020201-1100023132300203-1321122322032230-0003120131320333-2002022210102303-1122312020302311-2002301320012302-0123102312232102"></a>

#### `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.invert_matcher` property

Type: `"bool"`. Optional.

Invert Matcher. Invert Match of the expression defined.

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

- [item](resources--cdn_loadbalancer--reference--group-006.md#canonical-2200003022203033-3232100113133101-3330303230031122-3310313132032300-1302323220000332-2100013130213202-1020331333022300-3221312111301310): complete subsection reference.

<a id="canonical-3333011123232033-1113031111122313-2000022130321320-0332313303321110-0302003002211211-3112221331132202-2110122330210303-1230221203022102"></a>

<a id="canonical-2012321100033122-2213011003232213-2231203111013222-1022130331322231-3133112020031121-2031300021030222-0202302202222010-0230220120122320"></a>

#### `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.name` property

Type: `"string"`. Optional.

Cookie Name. A case-sensitive cookie name.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0022023013102010-2123110121321310-0012323310123032-1133003210323313-2002012113210013-1103120130131013-1022022322311101-0313121302122002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-006.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-006.md#canonical-3033211112003230-2302031233322223-3011210231211232-3130203102313021-2123313232211002-0033322022102113-0323310000003111-2330301133031030)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-2221210220332301-0132213302302221-0321121032111120-0323121203103301-2202313022220120-0033203221102333-1210021233111330-1101332333301310"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1120300322323132-1200130313310131-0002203230003120-3003321321113332-2222331133131013-3230333133133113-0122033123232211-2031231022331323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-006.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-006.md#canonical-3033211112003230-2302031233322223-3011210231211232-3130203102313021-2123313232211002-0033322022102113-0323310000003111-2330301133031030)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-1212323132202333-3100231123220010-3221133222020002-1113231032331223-0212001003010231-3300301102120313-1033101201303211-3331112210001302"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2200003022203033-3232100113133101-3330303230031122-3310313132032300-1302323220000332-2100013130213202-1020331333022300-3221312111301310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-006.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-006.md#canonical-3033211112003230-2302031233322223-3011210231211232-3130203102313021-2123313232211002-0033322022102113-0323310000003111-2330301133031030)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item

<a id="canonical-1221033222211231-2200220211030101-2131030103002200-1222101223133313-0110101121231303-3321310110032300-1121122211110310-3002002210312003"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-0032200321210012-0112030300313000-1210211020212012-3230203230311122-0123103231002201-0132023103130202-2032102010100032-0200121222122023"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item`

<a id="canonical-3332322123201122-1212103112233210-0111323302033203-0233023302313332-3013022230330011-1303020320103202-1223302110120010-2313101032231122"></a>

#### `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0220330120232233-3120000033013203-0312113320333201-3111012130213232-2221012000121303-2323220231322032-1110002022121332-2313030021313320"></a>

#### `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2020111230310000-3311112213320130-0233313101223022-2133311122323301-2212133032020110-2320203133100320-0022330120010203-0003331122303011"></a>

#### `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0002122103212131-2020122213023330-0233012102211313-2223031022231311-3110202023102222-1303322023020132-3300313112001132-0021201011102323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.headers` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-006.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- api_rate_limit.server_url_rules.request_matcher.headers

<a id="canonical-1331322133020001-2023222003233133-2323020002012031-2213132122033232-2012032322021212-2232231101003330-0302031011300010-3333121120202033"></a>

Type: `"object"`. list nested block, Optional.

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3130320231032101-3103032122001331-2001023133131030-0131232032303321-1312213230312130-2130332023321310-2232133202231302-0223311222212011"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.headers`

- [check_not_present](resources--cdn_loadbalancer--reference--group-006.md#canonical-3313133023110133-3231332201323201-0232132312101311-3230130032312001-1212102113222112-3202002113133022-0032222013230302-2122333100023121): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-006.md#canonical-1331001203220231-1210232321130113-0002023333310220-1120311111210203-3123123120321320-3330133133100323-0233320030311210-1120022231031332): complete subsection reference.

<a id="canonical-0223302231320013-3222103102122030-2211310321022223-1001201220022012-1110323220030231-3011201301100002-0223103011310102-0112331210113201"></a>

<a id="canonical-2322120222100330-0330222002023001-1110330101102323-0323020320221032-1103323221233312-1213312113312020-2013332113021123-3100310123332220"></a>

#### `api_rate_limit.server_url_rules.request_matcher.headers.invert_matcher` property

Type: `"bool"`. Optional.

Invert Header Matcher. Invert the match result.

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

- [item](resources--cdn_loadbalancer--reference--group-006.md#canonical-0221231223321303-3120300201222020-1120200011202210-3211331311120130-0220332033231000-1011003020230110-0311133003003031-3230012101233210): complete subsection reference.

<a id="canonical-0032203303231223-0303121210231230-1132011333221301-0002123111132023-1330200221002132-0323002200132303-0131132022232113-1132201032230313"></a>

<a id="canonical-0120201012200232-1103021132010202-2202232002320120-0302322212001333-1301000330012021-1300110102220333-1130212330321012-0102132000000220"></a>

#### `api_rate_limit.server_url_rules.request_matcher.headers.name` property

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3313133023110133-3231332201323201-0232132312101311-3230130032312001-1212102113222112-3202002113133022-0032222013230302-2122333100023121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.headers.check_not_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-006.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [api_rate_limit.server_url_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-006.md#canonical-0002122103212131-2020122213023330-0233012102211313-2223031022231311-3110202023102222-1303322023020132-3300313112001132-0021201011102323)
- api_rate_limit.server_url_rules.request_matcher.headers.check_not_present

<a id="canonical-3130030012231213-1303323332300120-2300122330301002-2010300230202330-0310323311223123-1302320002221012-0232112022200321-0303112323320321"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1331001203220231-1210232321130113-0002023333310220-1120311111210203-3123123120321320-3330133133100323-0233320030311210-1120022231031332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.headers.check_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-006.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [api_rate_limit.server_url_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-006.md#canonical-0002122103212131-2020122213023330-0233012102211313-2223031022231311-3110202023102222-1303322023020132-3300313112001132-0021201011102323)
- api_rate_limit.server_url_rules.request_matcher.headers.check_present

<a id="canonical-1333313223232333-2220000213213003-3230020333320332-0032312120101311-2333112202110321-0311332131210013-1312221033113332-2301030203231102"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0221231223321303-3120300201222020-1120200011202210-3211331311120130-0220332033231000-1011003020230110-0311133003003031-3230012101233210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.headers.item` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-006.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [api_rate_limit.server_url_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-006.md#canonical-0002122103212131-2020122213023330-0233012102211313-2223031022231311-3110202023102222-1303322023020132-3300313112001132-0021201011102323)
- api_rate_limit.server_url_rules.request_matcher.headers.item

<a id="canonical-3112002101313321-2112203230130322-1211232321021201-1113113220120331-0233132133031200-2220302302232000-3022111032013232-1313310033221222"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-3122033113202203-0322100003133121-0220110113301312-3223201111133220-3002302212210233-2022333122233222-3001321202001021-1310331300000120"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.headers.item`

<a id="canonical-0021120303301111-3321230013312221-3333002330031320-3000111201110121-1321313123333132-1010123112202003-0330332222111211-2322202310112123"></a>

#### `api_rate_limit.server_url_rules.request_matcher.headers.item.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3313232003311221-0013101301201200-2112302011331201-3121011200102030-2002032212100201-2210320303330011-1132230121031222-0123322322032333"></a>

#### `api_rate_limit.server_url_rules.request_matcher.headers.item.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2320120213002102-0123122013010011-1333220032332112-3132103220303003-3202121200022301-1221103102022221-0100023213222033-3113023210213333"></a>

#### `api_rate_limit.server_url_rules.request_matcher.headers.item.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2120002012323320-0210320131230122-1101202032031212-0321112122211211-2002131210323022-3122021323013001-1311113321020303-2332232002032320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.jwt_claims` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-006.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims

<a id="canonical-3212312113023033-2033123110232322-3113330222122021-2133230113312131-2013113100331130-2130132033012311-2111322322322132-3232120322310311"></a>

Type: `"object"`. list nested block, Optional.

A list of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates
must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3021101030032231-2031103133210100-1300332321211321-0132203022133230-1121320330223332-0203323110212103-1031331302122311-1233133212130031"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.jwt_claims`

- [check_not_present](resources--cdn_loadbalancer--reference--group-006.md#canonical-1110323000122121-3303301302312310-1320200003111323-3032011220203120-2131033201113102-2203313202211313-1132323001133122-1101331012201031): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-006.md#canonical-2322331120020121-0002113312012132-0111202000311011-2100002203231213-3200312100012310-3222010320113103-3313302201132000-3321102322211131): complete subsection reference.

<a id="canonical-0030311220322001-2002010011010112-2030220020200032-3322230202000003-1332331110213131-0313201230132032-0000301110202012-0121010112232301"></a>

<a id="canonical-2033212100030022-3102222302123031-3113102201101122-1012002223332231-0320202103102313-1201131333312020-1120300331322210-0303231331023220"></a>

#### `api_rate_limit.server_url_rules.request_matcher.jwt_claims.invert_matcher` property

Type: `"bool"`. Optional.

Invert Matcher. Invert the match result.

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

- [item](resources--cdn_loadbalancer--reference--group-006.md#canonical-0003022123333033-0123121313130321-2121132222111112-0130011033022321-2232120020210000-0221331031033030-0233013312030201-2011300220130001): complete subsection reference.

<a id="canonical-3002133203310330-3023102321102112-1113123113233013-0000003012322331-1110032031332201-0303223133111203-3131232101130202-0232121133001222"></a>

<a id="canonical-1201320312302021-0130232310200301-0023111001121030-1302033010201033-1303112002131111-0013311133231022-3112131331002000-3012330121311123"></a>

#### `api_rate_limit.server_url_rules.request_matcher.jwt_claims.name` property

Type: `"string"`. Optional.

JWT Claim Name. JWT claim name.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1110323000122121-3303301302312310-1320200003111323-3032011220203120-2131033201113102-2203313202211313-1132323001133122-1101331012201031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-006.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-006.md#canonical-2120002012323320-0210320131230122-1101202032031212-0321112122211211-2002131210323022-3122021323013001-1311113321020303-2332232002032320)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-1211101130033310-1313033332132001-3202330020130010-3210101123230132-0033313020012002-1032011112301230-0310121200201102-1132102321131003"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2322331120020121-0002113312012132-0111202000311011-2100002203231213-3200312100012310-3222010320113103-3313302201132000-3321102322211131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-006.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-006.md#canonical-2120002012323320-0210320131230122-1101202032031212-0321112122211211-2002131210323022-3122021323013001-1311113321020303-2332232002032320)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present

<a id="canonical-2232303200223332-1232123112231210-1222323330102001-0212113332320200-3032302231131122-0011211030023233-2111003302020331-1120010132131022"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0003022123333033-0123121313130321-2121132222111112-0130011033022321-2232120020210000-0221331031033030-0233013312030201-2011300220130001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.jwt_claims.item` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-006.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-006.md#canonical-2120002012323320-0210320131230122-1101202032031212-0321112122211211-2002131210323022-3122021323013001-1311113321020303-2332232002032320)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims.item

<a id="canonical-2121231321210132-2203000303010210-2020300212022333-1002122211333233-0300123232010313-0311311030013211-3113210023010323-2210111100003300"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-2223011333321302-1030203122121202-3113102111010211-1030113233201221-0222033321111122-1011101031020210-3200212302132033-2113002101202000"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.jwt_claims.item`

<a id="canonical-2033301130312200-2332222203233331-1211131212120323-0203021013202221-2222231221100000-3202012200222020-2210101301200303-1022331312110123"></a>

#### `api_rate_limit.server_url_rules.request_matcher.jwt_claims.item.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2012000023310003-1003232330021212-3013210310310320-3302011211033211-3013232323122023-3111112033012133-1301301321210303-3110200210133032"></a>

#### `api_rate_limit.server_url_rules.request_matcher.jwt_claims.item.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2303203011300333-0232111101332020-1122233020331032-1021133232200222-0312233120310132-2203120031031123-2211311131013200-3132333333323133"></a>

#### `api_rate_limit.server_url_rules.request_matcher.jwt_claims.item.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0122313023112111-2300122102223032-0012011121332013-0133313131303331-3101132233230020-2203113102113133-3002130201213302-3001033321123231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.query_params` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-006.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- api_rate_limit.server_url_rules.request_matcher.query_params

<a id="canonical-1231321222222110-0033113223323002-1010111300220300-2311321231203323-0011103330120103-2003220003003202-0132332113321123-0102212330123003"></a>

Type: `"object"`. list nested block, Optional.

A list of predicates for all query parameters that need to be matched. The criteria for matching
each query parameter are described in individual instances of QueryParameterMatcherType. The actual
query parameter values are extracted from the request API as a list of strings for each query
parameter name. Note that all specified query parameter predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1003121030032031-3132021122021303-0323020332113121-2331222033101303-3332012031023103-3100300223322322-0112333201123221-2222121323301003"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.query_params`

- [check_not_present](resources--cdn_loadbalancer--reference--group-006.md#canonical-0031211030333021-3121133300032202-2222032100113212-0222322311221311-3200013103011122-2032001331333233-1121130331332110-2130300203113110): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-006.md#canonical-3312033211310313-3303303032031201-0313212311020003-1002300120123310-3102100123213103-2210322300013230-0023122312323202-2033033022000301): complete subsection reference.

<a id="canonical-3001121131031030-1210110313021110-0001101202111212-2021303200301221-3203132033032103-0122003112113130-3310200121002003-0313310210311130"></a>

<a id="canonical-2312300133230322-2333121021300202-3032123023213122-3121010013112112-0021313012022110-2001012213313121-3200121021323310-1222012100123013"></a>

#### `api_rate_limit.server_url_rules.request_matcher.query_params.invert_matcher` property

Type: `"bool"`. Optional.

Invert Query Parameter Matcher. Invert the match result.

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

- [item](resources--cdn_loadbalancer--reference--group-006.md#canonical-3022303020223232-1111323302302211-3023001310232001-1301133223032302-2031221220101030-3011311302103021-3213330300230233-0302101230233121): complete subsection reference.

<a id="canonical-0202211210031202-0333122030200020-1113311312023111-1200320211101313-2220002033232233-3132003303203231-0301102310033003-3302031030100020"></a>

<a id="canonical-3103120121132101-2120033123312133-0123222011133223-1013031233220313-1300113003212002-3301211113132232-0233033021222013-1330222202321122"></a>

#### `api_rate_limit.server_url_rules.request_matcher.query_params.key` property

Type: `"string"`. Optional.

A case-sensitive HTTP query parameter name.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0031211030333021-3121133300032202-2222032100113212-0222322311221311-3200013103011122-2032001331333233-1121130331332110-2130300203113110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-006.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [api_rate_limit.server_url_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-006.md#canonical-0122313023112111-2300122102223032-0012011121332013-0133313131303331-3101132233230020-2203113102113133-3002130201213302-3001033321123231)
- api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present

<a id="canonical-0012121133222020-1030131122232111-1233302033330021-3230323321023021-0200132320213100-0211123201223001-0330111223232103-0110023212113231"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3312033211310313-3303303032031201-0313212311020003-1002300120123310-3102100123213103-2210322300013230-0023122312323202-2033033022000301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.query_params.check_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-006.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [api_rate_limit.server_url_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-006.md#canonical-0122313023112111-2300122102223032-0012011121332013-0133313131303331-3101132233230020-2203113102113133-3002130201213302-3001033321123231)
- api_rate_limit.server_url_rules.request_matcher.query_params.check_present

<a id="canonical-0201200311312110-1212232202311321-0212020031330121-3202010300333132-2332021223320111-1022232222013110-3333232210023002-0310332111210302"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3022303020223232-1111323302302211-3023001310232001-1301133223032302-2031221220101030-3011311302103021-3213330300230233-0302101230233121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.query_params.item` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-006.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [api_rate_limit.server_url_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-006.md#canonical-0122313023112111-2300122102223032-0012011121332013-0133313131303331-3101132233230020-2203113102113133-3002130201213302-3001033321123231)
- api_rate_limit.server_url_rules.request_matcher.query_params.item

<a id="canonical-1211321221121011-3111113022112222-3321233101000030-2131103233321302-1333200102022230-0203323112310220-2031210012123232-2202132322222232"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-3032203303020112-3032310001220322-2223310303110131-2131033011300111-2031332311033320-2222100123211033-2211033201023102-3210100103230113"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.query_params.item`

<a id="canonical-3301123011303113-3021030010102310-2222131101101210-3100001111313100-3301023303130030-3222332201130332-1223022020133320-2112103000233033"></a>

#### `api_rate_limit.server_url_rules.request_matcher.query_params.item.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3032313103030201-0213233331103302-2200233300320100-3021132312102122-3003111021031122-0032300103111203-2323031032131023-3132211202130112"></a>

#### `api_rate_limit.server_url_rules.request_matcher.query_params.item.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1303303232312323-0123032322311001-2133121121201320-3002220203032222-3003102203010310-1021223130020012-1121123221013221-3233130110022320"></a>

#### `api_rate_limit.server_url_rules.request_matcher.query_params.item.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- api_specification

<a id="canonical-0102031021120331-0211222232131320-1022002322130010-0221000333310320-3012323211302110-0213323320200103-0202010323332222-0330001323100020"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: api\_specification, disable\_api\_definition; Default: disable\_api\_definition\] Settings
for API specification (API definition, OpenAPI validation, etc.).

Provider validators and defaults (from schema source):

```go
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

- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-0102031021120331-0211222232131320-1022002322130010-0221000333310320-3012323211302110-0213323320200103-0202010323332222-0330001323100020)
- [disable_api_definition](resources--cdn_loadbalancer--reference--group-010.md#canonical-1013231301331321-3201321033031021-0303223333210301-2112110233330302-0302100120123301-0132331033332233-0312121102001302-1222030220330231)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
api_specification {
  # Configure direct properties listed below.
}
```
