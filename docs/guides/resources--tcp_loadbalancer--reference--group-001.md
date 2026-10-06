---
page_title: "xcsh_tcp_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_tcp_loadbalancer reference."
---

# xcsh_tcp_loadbalancer reference

<a id="canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- Property reference

<a id="canonical-1331102000101233-2020000131212322-1232303122103300-1010023320210220-1010233220023321-1332310030003312-0330123230133322-0332300222220101"></a>

### Direct properties for `xcsh_tcp_loadbalancer`

- [active_service_policies](resources--tcp_loadbalancer--reference--group-001.md#canonical-2123332010300020-1023231311030232-1300031122101011-1032023203210220-1200333010322303-0323133123222003-1020130133223133-1112131030203102): complete subsection reference.

- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-0101013130323102-0111020001013232-1223003210233021-3112202320110020-1310022023311310-1003000113120300-2312202022202013-0222311212101130): complete subsection reference.

- [advertise_on_public](resources--tcp_loadbalancer--reference--group-002.md#canonical-1101101122210003-2330111331303010-2130211300213202-2120322121300012-2033321200023022-3103313203330222-2302000021033232-1221001021130221): complete subsection reference.

- [advertise_on_public_default_vip](resources--tcp_loadbalancer--reference--group-002.md#canonical-3333100232210200-1000122300031311-3012302322223020-3032022132323133-3310012110333012-2333013330001120-1130222201132013-1120103002011310): complete subsection reference.

<a id="canonical-2131000230031113-0213111121102030-2003021131030103-3123220303022321-0113120030212112-3201102230220033-0230300013302323-2002123223220133"></a>

<a id="canonical-2311101102202323-2133011201002310-2303031013033223-2212330300331131-3233332331133001-3313312312210030-1321300032012331-3131332221231110"></a>

#### `annotations` property

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.values.string.max_len": "1024",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 1024,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [default_lb_with_sni](resources--tcp_loadbalancer--reference--group-002.md#canonical-1233300113221313-3110203321303320-1132100212321211-3233111001220210-2202120330231303-1320322300203123-3132233222313123-1030123011100010): complete subsection reference.

<a id="canonical-0210300032302033-1010012212111101-2300313023100113-1033222123022332-1320011331011221-1101131323010232-3310233030122121-0122030021213023"></a>

<a id="canonical-3132212223111232-1301320200110212-0211230001012101-3202333200031300-2222211012210101-3333220122122211-0103322312202301-2122320033122311"></a>

#### `description` property

Type: `"string"`. Optional.

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="canonical-3211333200201110-3103030023122311-3210331003222030-0221231311222210-1310311121332312-3031311331201132-0221203200022123-0330120013130312"></a>

<a id="canonical-2213331210102122-0132321121331230-2213210031030313-2303232320223332-2102033203013200-1200110333222121-1232301021220233-1301333233030321"></a>

#### `disable` property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Additional upstream details:

A value of true will administratively disable the object.

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

<a id="canonical-1301013030310032-1321231213232210-3131002120333323-1213223112311112-2121332113202211-1303233101201021-0123333210323001-3013010113311202"></a>

<a id="canonical-1231003303223223-1111231203113121-1032210030100022-3132203112132331-2102311231210310-3131101002303021-1021301100120123-1233133212123203"></a>

#### `dns_volterra_managed` property

Type: `"bool"`. Optional, Computed.

DNS records for domains will be managed automatically by F5 Distributed Cloud. This requires the
domain to be delegated to F5XC using the Delegated Domain feature. Defaults to \`false\`. Server
applies default when omitted.

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

- [do_not_advertise](resources--tcp_loadbalancer--reference--group-002.md#canonical-2322213221112110-0231131200021100-2323031112011321-2113223331230331-1110202310201022-0323112230111011-3203121033311100-1130120122122120): complete subsection reference.

- [do_not_retract_cluster](resources--tcp_loadbalancer--reference--group-002.md#canonical-2322113213320330-0102033300000032-1130020131303320-1130033301222310-0100131123323030-1300033212333030-2033201321130131-2130003330320013): complete subsection reference.

<a id="canonical-2311022230003011-3323222012332031-2330233100031023-0230001112122130-3201032212210102-0320003312010303-3032020001331300-2111232131211301"></a>

<a id="canonical-1120033010000332-2032010022203012-1303312212010230-0021221313033022-3113221130100002-1030300221232210-2112032310221030-0020213011022332"></a>

#### `domains` property

Type: `["list", "string"]`. Optional.

A list of Domains (host/authority header) that will be matched to this Load Balancer.

Supported Domains and search order: &#8203;1. Exact Domain names: www&#46;example.com. &#8203;2.
Domains starting with a Wildcard: \*.example.com.

Not supported Domains: &#8203;- Just a Wildcard: \* &#8203;- A Wildcard and TLD with no root Domain:
\*.com. &#8203;- A Wildcard not matching a whole DNS label. E.g. \*.example.com and
\*.bar.example.com are valid Wildcards however \*bar.example.com, \*-bar.example.com, and
bar\*.example.com are all invalid.

Additional notes: A Wildcard will not match empty string. E.g. \*.example.com will match
bar.example.com and baz-bar.example.com but not .example.com. The longest Wildcards match first.
Only a single virtual host in the entire route configuration can match on \*. Also a Domain must be
unique across all virtual hosts within an advertise policy.

Domains are also used for SNI matching if SNI is activated on the given TCP Load Balancer. Domains
also indicate the list of names for which DNS resolution will be automatically resolved to IP
addresses by the system.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 32),
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
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [hash_policy_choice_least_active](resources--tcp_loadbalancer--reference--group-002.md#canonical-3231312101111010-1230322232102131-1310111121030233-1233021020330300-2012003120131022-3230203210022022-1331310221131200-1310111032002330): complete subsection reference.

- [hash_policy_choice_random](resources--tcp_loadbalancer--reference--group-002.md#canonical-2133201212213303-3220313201223303-0031130110231110-1213131231101232-3321233020122003-3302010033002303-0202202110320320-1222200211330032): complete subsection reference.

- [hash_policy_choice_round_robin](resources--tcp_loadbalancer--reference--group-002.md#canonical-2011303133011102-2021121122120220-1133330230131122-0002121330231010-2103102030013022-2020210201310012-0000100210012131-3213113301223032): complete subsection reference.

- [hash_policy_choice_source_ip_stickiness](resources--tcp_loadbalancer--reference--group-002.md#canonical-3123323133211320-2223021102110033-2023122331312031-1312102001331031-2023122213321233-0303012321200222-1233301221213131-3010122320223012): complete subsection reference.

<a id="canonical-2231213110320012-1202221203123003-0013222123323322-3031302333332031-3120011310001121-1033013100102011-3321133323010033-2013320030203312"></a>

<a id="canonical-3031033210132203-3131331020102020-1213103313112331-0022212112033111-0133000222322201-3123212030001232-1132221122023022-1032122202232322"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0131112213231103-1132301023000303-3213220011211213-2303000030010012-2103121033122102-3000213220103223-2110220113310300-3213310301303312"></a>

<a id="canonical-2212213133300132-0111311232302301-2200203213022022-0212001031330023-0000333311202223-2133020113301313-3311123202103120-2313102131132212"></a>

#### `idle_timeout` property

Type: `"number"`. Optional, Computed.

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.
Server applies default when omitted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(4147200000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 4147200000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "4147200000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "4147200000"
  }
}
```

<a id="canonical-0332111223013101-3013220013022010-0033233202300100-1023320112113122-3120331020031032-2333222013003333-3122020112320123-1210310312210212"></a>

<a id="canonical-3301202101203030-2232032120332301-1001130201232213-3221111330131222-1030320312231131-1301120332013222-1332002210001213-0333102300013223"></a>

#### `labels` property

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
filtering.

Additional upstream details:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

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

<a id="canonical-0103102333022111-2021213011032103-3032103010200230-3202330110112210-3123320133213130-2203103322200013-2322003300231211-3113133010101313"></a>

<a id="canonical-1212132021111331-3323113133330030-2001310011023100-3221032333200100-2033011310303031-3110033003121031-2222033103103212-1330000220222320"></a>

#### `listen_port` property

Type: `"number"`. Optional, Computed.

\[OneOf: listen\_port, port\_ranges\] Exclusive with \[port\_ranges\] Listen Port for this load
balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

OneOf alternatives in this subsection:

- [listen_port](resources--tcp_loadbalancer--reference--group-001.md#canonical-0103102333022111-2021213011032103-3032103010200230-3202330110112210-3123320133213130-2203103322200013-2322003300231211-3113133010101313)
- [port_ranges](resources--tcp_loadbalancer--reference--group-001.md#canonical-2210103003232120-0020010122112210-3121130230322133-1310230033123311-3003032123212233-0300132332322301-1122113023321201-0203320231013212)

Select alternatives according to the provider validators above.

<a id="canonical-3220032130201011-1110211113033333-0223302122130231-1311203302233132-3303103312301022-0111301321303011-0211311021133323-1033033032033000"></a>

<a id="canonical-1212101221110132-3111302302321311-3021212101001000-0102003113021323-3130222120100033-2002000231320312-0003122132210222-3202220232103130"></a>

#### `name` property

Type: `"string"`. Required.

Name of the TCP Load Balancer. Must be unique within the namespace.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-0200000111122231-0232213210213212-2103020023032012-2232022333102131-3333130002021332-1012100102113301-3110330013322223-0003302022331330"></a>

<a id="canonical-0231021323230002-0122221302120033-1311300210032012-3231113212212302-3030231332000132-1230310222112012-1232132321012332-2312123323001103"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the TCP Load Balancer is created.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
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

- [no_service_policies](resources--tcp_loadbalancer--reference--group-002.md#canonical-0211303031122133-3002101102330112-3303102231112031-0323212233100110-3023331222023022-0122331220121021-0201223330103012-0220122230220001): complete subsection reference.

- [no_sni](resources--tcp_loadbalancer--reference--group-002.md#canonical-0332121320022012-0131110131002221-3211220033203220-3333221011312320-1303130303113032-2011002222333033-3001212201031002-1010101011201223): complete subsection reference.

- [origin_pools_weights](resources--tcp_loadbalancer--reference--group-002.md#canonical-2100010003132033-1331210120313310-0212003223033302-1320122132000102-1302201010312200-3333221331300211-3332023322113011-0102012310233122): complete subsection reference.

<a id="canonical-2210103003232120-0020010122112210-3121130230322133-1310230033123311-3003032123212233-0300132332322301-1122113023321201-0203320231013212"></a>

<a id="canonical-1311202131311230-2323301123022303-1000032132031000-2321211231222111-1310322200133310-3232303130113122-2321210222000030-1300233320111203"></a>

#### `port_ranges` property

Type: `"string"`. Optional, Computed.

Exclusive with \[listen\_port\] A string containing a comma separated list of port ranges. Each port
range consists of a single port or two ports separated by '-'.

Additional upstream details:

Each port range consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

- [retract_cluster](resources--tcp_loadbalancer--reference--group-003.md#canonical-2203330232231221-0231100222101101-2032200123103102-3022031012302100-0322123221223231-2330022031223301-3202022122302011-3102020113332333): complete subsection reference.

- [service_policies_from_namespace](resources--tcp_loadbalancer--reference--group-003.md#canonical-1302000233122001-0330320220303103-0201012112213303-1133102211110130-2333133012210201-3013203312103310-2212130213030102-1122113100303223): complete subsection reference.

- [sni](resources--tcp_loadbalancer--reference--group-003.md#canonical-1211312003133330-1223211300322220-2113300203322232-3010131033013022-2332011121303131-3303223200010132-3300002312210210-2000310102320322): complete subsection reference.

- [tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-2220232031020123-0333221323000013-0330311311100200-2313132103001132-1300312202210020-1103313303133002-0011300213213202-1132012101023211): complete subsection reference.

- [timeouts](resources--tcp_loadbalancer--reference--group-003.md#canonical-1223011333230122-2303021022201212-1030021321212322-1311000300020110-1130221333112233-3223223303001030-2231022021122211-0113200021221011): complete subsection reference.

- [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120): complete subsection reference.

- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-003.md#canonical-1020000333221312-0212323132120330-0112301302023312-1010022111201312-0212120203321302-3011031322120320-0022031132030332-3001230220212331): complete subsection reference.

<a id="canonical-0213132103121213-2111113310033002-3231130131101112-3130300123312303-0131032101103110-1201311001132312-1300213233121001-3233020130101100"></a>

### All schema paths for `xcsh_tcp_loadbalancer`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `active_service_policies` | [active_service_policies](resources--tcp_loadbalancer--reference--group-001.md#canonical-1211213311100203-3013023320103303-3223003322203320-0331213010022010-0001333301100210-0022302033003002-2302133212310230-3112000003210013) |
| `active_service_policies.policies` | [active_service_policies.policies](resources--tcp_loadbalancer--reference--group-001.md#canonical-1021220010111222-0012031311020213-0212322031031001-2302322202223003-0023101312211013-1323013313031122-1021102000213020-3032233133313022) |
| `active_service_policies.policies.name` | [active_service_policies.policies.name](resources--tcp_loadbalancer--reference--group-001.md#canonical-3302102322212223-0220012013110110-3111012133113323-3202233121111113-3313103212223113-2310001323123201-3311030311110132-1131020303313020) |
| `active_service_policies.policies.namespace` | [active_service_policies.policies.namespace](resources--tcp_loadbalancer--reference--group-001.md#canonical-1003233022033031-1030322201031222-3033132222330123-2233223202130130-1301012022301010-0303322113300322-1200112210010010-3212333223002110) |
| `active_service_policies.policies.tenant` | [active_service_policies.policies.tenant](resources--tcp_loadbalancer--reference--group-001.md#canonical-2111310323203122-2030220313001320-2100222123223211-3320102230102102-1312313123002322-3033210322301002-0012221213302131-3000012301202300) |
| `advertise_custom` | [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-1301323022031230-1003003001020123-1132201332123232-2311223310032312-1132021010021312-1231021033101001-1131013213121211-2103230021312203) |
| `advertise_custom.advertise_where` | [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-3131003031113023-0220010230102313-0300110222112030-3300113232132233-3331103203103011-0320002201322122-1212023321021000-3111233322220223) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public` | [advertise_custom.advertise_where.advertise_dualstack_on_public](resources--tcp_loadbalancer--reference--group-001.md#canonical-0221013131212033-0013120323321102-3323220031131213-0222130200222221-0313231022200022-1301131202233120-2032213103011301-1210211312102100) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip](resources--tcp_loadbalancer--reference--group-001.md#canonical-3332001223333303-3233000200022303-2321323123102312-0022303021030000-1021310312230322-0322113012112313-1330103032002000-1333133110121111) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name](resources--tcp_loadbalancer--reference--group-001.md#canonical-0221322202223201-1123100322330313-3231133311311312-0221021122230232-1123221302020202-0011011332210010-0202331311303202-3033311202212310) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace](resources--tcp_loadbalancer--reference--group-001.md#canonical-3323102333211120-0101210012201003-1032323201330203-1221331133232232-3322103100033303-3032103302303030-3033200031331220-0003310000200202) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant](resources--tcp_loadbalancer--reference--group-001.md#canonical-0313202121202003-1200320220033121-1132201322020102-2023101031023232-0123102323011122-0121233030020222-1000310330221213-2321023232233002) |
| `advertise_custom.advertise_where.advertise_on_public` | [advertise_custom.advertise_where.advertise_on_public](resources--tcp_loadbalancer--reference--group-001.md#canonical-1201200023021201-0022030201023003-0103320033131331-0000301102200231-3001022032020300-3322212200133032-1321112001233010-0323231311113011) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip` | [advertise_custom.advertise_where.advertise_on_public.public_ip](resources--tcp_loadbalancer--reference--group-001.md#canonical-0332233033123312-0102121003221012-3222221121020122-0022020210311233-3232202002303010-2130110133003003-3030030103200003-2231121111132312) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip.name` | [advertise_custom.advertise_where.advertise_on_public.public_ip.name](resources--tcp_loadbalancer--reference--group-001.md#canonical-3300313210230301-0332223103002122-0311202001021020-0202020033202022-1010202003133022-1003203121122110-2110101113323103-2323321320231003) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip.namespace` | [advertise_custom.advertise_where.advertise_on_public.public_ip.namespace](resources--tcp_loadbalancer--reference--group-001.md#canonical-2231213000202020-3002220300131232-3011131123013301-1202102012211022-1332312311233223-1233322230233110-2323023033020120-3123221101013032) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip.tenant` | [advertise_custom.advertise_where.advertise_on_public.public_ip.tenant](resources--tcp_loadbalancer--reference--group-001.md#canonical-1203002101333331-0311103033111033-0223230232212203-2333012003102112-3200110333121003-2132232223320131-0313222321002102-3202022220331121) |
| `advertise_custom.advertise_where.advertise_v6_on_public` | [advertise_custom.advertise_where.advertise_v6_on_public](resources--tcp_loadbalancer--reference--group-001.md#canonical-2303333110033321-0232231022302200-2111310303132201-2102232021330023-2200131322010331-0101020022101002-1133232131120010-1223330011311310) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip](resources--tcp_loadbalancer--reference--group-001.md#canonical-0012012132323332-1022112231133122-1200032003331230-1113113012222130-2211203331020323-0000321232333311-1003220212031020-0021121120312310) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name](resources--tcp_loadbalancer--reference--group-001.md#canonical-1313133112303321-3313300322202013-3030231212010212-0110101300113302-2133303021100311-2323031022033111-2023112313111301-0020022022320211) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace](resources--tcp_loadbalancer--reference--group-001.md#canonical-2333012203131021-2231333212012003-2220103302332302-1023123111320211-2133332130133321-1010310201023112-1210022322132320-0331023111131201) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant](resources--tcp_loadbalancer--reference--group-001.md#canonical-0110320130203131-3103330310223131-0021011322231100-1232131013013012-0223100001111230-3121320321200320-3100033130133321-1311321302213131) |
| `advertise_custom.advertise_where.port` | [advertise_custom.advertise_where.port](resources--tcp_loadbalancer--reference--group-001.md#canonical-2002133223111111-1332301233021023-2023320112110320-0132101131320300-1222322222313100-3222021031112003-3122113122212320-0301310021002200) |
| `advertise_custom.advertise_where.port_ranges` | [advertise_custom.advertise_where.port_ranges](resources--tcp_loadbalancer--reference--group-001.md#canonical-2020233021133322-1223231322123313-0013211212333111-3122032101023223-0122112102132110-1011000101200201-3101130220022100-3122220301122023) |
| `advertise_custom.advertise_where.site` | [advertise_custom.advertise_where.site](resources--tcp_loadbalancer--reference--group-001.md#canonical-0233002212112233-3111330321132031-3121222333110322-3021330031300232-1303220232320133-0011323303112302-3032101121311222-1110223323033002) |
| `advertise_custom.advertise_where.site.ip` | [advertise_custom.advertise_where.site.ip](resources--tcp_loadbalancer--reference--group-001.md#canonical-1113230300201002-1332212231203013-3010303033032200-3233112012302130-3032011322120312-1222100231330122-0210122212300312-2310030320331022) |
| `advertise_custom.advertise_where.site.network` | [advertise_custom.advertise_where.site.network](resources--tcp_loadbalancer--reference--group-001.md#canonical-3011111001110212-1023211021302000-2211212223201003-1103100111331111-2000310332023020-3121110310133200-2023320121220011-2230231312223101) |
| `advertise_custom.advertise_where.site.site` | [advertise_custom.advertise_where.site.site](resources--tcp_loadbalancer--reference--group-001.md#canonical-0333313201001022-2313031213321231-0132021101002003-3230023031113000-3120220323200230-0231030202030323-1113111301131021-1333031320113202) |
| `advertise_custom.advertise_where.site.site.name` | [advertise_custom.advertise_where.site.site.name](resources--tcp_loadbalancer--reference--group-001.md#canonical-1210232312302010-1302132201332010-1312211222003320-3100311111101220-2311132320321202-2313233201230211-2022110212303031-0322300213200133) |
| `advertise_custom.advertise_where.site.site.namespace` | [advertise_custom.advertise_where.site.site.namespace](resources--tcp_loadbalancer--reference--group-001.md#canonical-2033122031331020-0100100121330230-3210002110302121-2013122231122233-2231223131310330-1003313112033312-1300213020002000-1223010232332120) |
| `advertise_custom.advertise_where.site.site.tenant` | [advertise_custom.advertise_where.site.site.tenant](resources--tcp_loadbalancer--reference--group-001.md#canonical-3113113123012313-0313130312223311-1012221013103101-1122211320033332-3000303220103300-1001301130011202-3323100001120222-2102302122302230) |
| `advertise_custom.advertise_where.use_default_port` | [advertise_custom.advertise_where.use_default_port](resources--tcp_loadbalancer--reference--group-001.md#canonical-2120120222220122-3323012000120222-2202102233331233-2101022230221031-0201321302033102-3300022313330233-1200131130212300-2311310103212210) |
| `advertise_custom.advertise_where.virtual_network` | [advertise_custom.advertise_where.virtual_network](resources--tcp_loadbalancer--reference--group-001.md#canonical-2203310231222103-1020202323302322-3300011300002121-2100232333130220-2302211013332332-2321310111100322-3013310130031021-2023131333112201) |
| `advertise_custom.advertise_where.virtual_network.default_v6_vip` | [advertise_custom.advertise_where.virtual_network.default_v6_vip](resources--tcp_loadbalancer--reference--group-002.md#canonical-2300230330233301-3220031330202032-0311013320002003-1311323021331311-1020131330131030-0132021021132223-3032203123132010-2312313323010301) |
| `advertise_custom.advertise_where.virtual_network.default_vip` | [advertise_custom.advertise_where.virtual_network.default_vip](resources--tcp_loadbalancer--reference--group-002.md#canonical-2213333200110023-3303310232003132-0012011013120201-3110100303102200-1031012011220110-1121102000032333-1232133120312133-2211203032231001) |
| `advertise_custom.advertise_where.virtual_network.specific_v6_vip` | [advertise_custom.advertise_where.virtual_network.specific_v6_vip](resources--tcp_loadbalancer--reference--group-001.md#canonical-0021222232011222-1023122310013311-1120212012210322-3212002123023303-1000220300330333-0022100330301332-2330301010203333-3022032031232230) |
| `advertise_custom.advertise_where.virtual_network.specific_vip` | [advertise_custom.advertise_where.virtual_network.specific_vip](resources--tcp_loadbalancer--reference--group-002.md#canonical-0222332231021212-3213220030121313-1300022001031030-0212122201001200-0301322230020123-1112103233131312-1120020332322011-2133131021022333) |
| `advertise_custom.advertise_where.virtual_network.virtual_network` | [advertise_custom.advertise_where.virtual_network.virtual_network](resources--tcp_loadbalancer--reference--group-002.md#canonical-2013303220021031-3010201222213310-3201121113013223-0220120113123300-3330010232230213-3022010032100010-1111302213333313-2200300113123021) |
| `advertise_custom.advertise_where.virtual_network.virtual_network.name` | [advertise_custom.advertise_where.virtual_network.virtual_network.name](resources--tcp_loadbalancer--reference--group-002.md#canonical-2111332202231002-2133301010003332-0220213322013133-2113011322121202-1031113200311232-0212130101220311-1320110122122020-0213111313222221) |
| `advertise_custom.advertise_where.virtual_network.virtual_network.namespace` | [advertise_custom.advertise_where.virtual_network.virtual_network.namespace](resources--tcp_loadbalancer--reference--group-002.md#canonical-2210312203132311-0302010313032130-0210310312223100-0000203122103221-1032322202020311-0231112112201013-2112311330322220-2212131223210002) |
| `advertise_custom.advertise_where.virtual_network.virtual_network.tenant` | [advertise_custom.advertise_where.virtual_network.virtual_network.tenant](resources--tcp_loadbalancer--reference--group-002.md#canonical-3300312221232222-1021131103013010-1323211203211030-3110322221320310-2011220332222010-3223333312322313-3321211330202223-2111332103230301) |
| `advertise_custom.advertise_where.virtual_site` | [advertise_custom.advertise_where.virtual_site](resources--tcp_loadbalancer--reference--group-002.md#canonical-2101030020321030-0020230030332110-0302121201011233-0232202000223022-0333133322220100-2311021122323033-0333031032200232-1120303320320100) |
| `advertise_custom.advertise_where.virtual_site.network` | [advertise_custom.advertise_where.virtual_site.network](resources--tcp_loadbalancer--reference--group-002.md#canonical-2301303022310123-2132120022313212-0201313013313312-2201213323302000-1313200131130131-1221312100003112-2131122320220230-0021302313232211) |
| `advertise_custom.advertise_where.virtual_site.virtual_site` | [advertise_custom.advertise_where.virtual_site.virtual_site](resources--tcp_loadbalancer--reference--group-002.md#canonical-2211322122230300-3333302211031120-3102332010002203-3121213331101203-0030103030113021-0133303121013211-1013332322203013-2130123113201220) |
| `advertise_custom.advertise_where.virtual_site.virtual_site.name` | [advertise_custom.advertise_where.virtual_site.virtual_site.name](resources--tcp_loadbalancer--reference--group-002.md#canonical-3311113020331212-3021133312100032-0231203321312322-0320022012011303-0213300123223311-1033211010201131-1312222322203331-0322330230121022) |
| `advertise_custom.advertise_where.virtual_site.virtual_site.namespace` | [advertise_custom.advertise_where.virtual_site.virtual_site.namespace](resources--tcp_loadbalancer--reference--group-002.md#canonical-3201212132120202-3233100321103132-0212301123212310-2233230200031313-2233032313031033-3031221310033103-3023013200221313-1220130231121022) |
| `advertise_custom.advertise_where.virtual_site.virtual_site.tenant` | [advertise_custom.advertise_where.virtual_site.virtual_site.tenant](resources--tcp_loadbalancer--reference--group-002.md#canonical-0020020132032100-3013000112130012-3213222301110023-0003110112110013-2131312112220222-1011001123310230-2321001211323311-2121102223230230) |
| `advertise_custom.advertise_where.virtual_site_with_vip` | [advertise_custom.advertise_where.virtual_site_with_vip](resources--tcp_loadbalancer--reference--group-002.md#canonical-0331011202312231-2001120230211033-2322202110213221-3101121130032223-3101133203330023-3113302012121122-1320122101033030-0210332013213123) |
| `advertise_custom.advertise_where.virtual_site_with_vip.ip` | [advertise_custom.advertise_where.virtual_site_with_vip.ip](resources--tcp_loadbalancer--reference--group-002.md#canonical-1311000001220100-0032223213001313-1020333321030333-3010333111302320-1130002231032222-3103102000111023-1311301330201112-3023311101100022) |
| `advertise_custom.advertise_where.virtual_site_with_vip.network` | [advertise_custom.advertise_where.virtual_site_with_vip.network](resources--tcp_loadbalancer--reference--group-002.md#canonical-1200202002213230-3320332101322113-2003220320330220-0031313021101033-1202031012210013-2232332121120131-1320103011220031-0130000333232001) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](resources--tcp_loadbalancer--reference--group-002.md#canonical-1001202121302300-2231113032010323-3222121313021223-1131311012013301-3103201030112201-1010313110102030-1133101003210011-2311300302223002) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name](resources--tcp_loadbalancer--reference--group-002.md#canonical-3300122301013231-2022130100012202-1102220032121033-1313123233122100-0322113020111120-2022212030002213-1233121100013313-3133002220012230) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace](resources--tcp_loadbalancer--reference--group-002.md#canonical-0003003330211323-1131132201203213-0311320021031021-2100200001023202-1122311021320213-2010223310111032-2113030022130310-3101213313010230) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant](resources--tcp_loadbalancer--reference--group-002.md#canonical-2311000230231031-3102332121030203-0210002112013300-1312301320301332-0002133333213102-0100330301130030-2020120010303223-2213222313133332) |
| `advertise_custom.advertise_where.vk8s_service` | [advertise_custom.advertise_where.vk8s_service](resources--tcp_loadbalancer--reference--group-002.md#canonical-2032303212230003-0200223230112002-3132232213103110-2001011022332302-0332200200001332-2033111311013003-3012002113232130-3102101201313131) |
| `advertise_custom.advertise_where.vk8s_service.site` | [advertise_custom.advertise_where.vk8s_service.site](resources--tcp_loadbalancer--reference--group-002.md#canonical-3033230022332321-0231212233003132-1132003120132200-2211031231220011-1231101313123321-2310010200120010-1113023220233311-0213023330031301) |
| `advertise_custom.advertise_where.vk8s_service.site.name` | [advertise_custom.advertise_where.vk8s_service.site.name](resources--tcp_loadbalancer--reference--group-002.md#canonical-3130032322213030-3032022102012313-3313210030131300-3211013020303030-1301122220320022-2201103230332313-1201132101202320-3100221303012111) |
| `advertise_custom.advertise_where.vk8s_service.site.namespace` | [advertise_custom.advertise_where.vk8s_service.site.namespace](resources--tcp_loadbalancer--reference--group-002.md#canonical-1232133211032003-1113323021002211-2332123230032210-2120232323022132-3323200311210220-3302102123233030-0030003333103222-1132020230133100) |
| `advertise_custom.advertise_where.vk8s_service.site.tenant` | [advertise_custom.advertise_where.vk8s_service.site.tenant](resources--tcp_loadbalancer--reference--group-002.md#canonical-0000312230110203-3123331312133333-1022111132213330-2000001210002232-3120112202210003-0333003010100301-2012231031233212-3131310002000233) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site` | [advertise_custom.advertise_where.vk8s_service.virtual_site](resources--tcp_loadbalancer--reference--group-002.md#canonical-2112123100031210-3120032221323121-2033302333321130-3323200332003223-0033120310200021-2310123321010101-0023120132332210-1312203221311012) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site.name` | [advertise_custom.advertise_where.vk8s_service.virtual_site.name](resources--tcp_loadbalancer--reference--group-002.md#canonical-0111300312122122-1011232001322002-2333203111203012-0222021231213031-1312213121121211-1231012102130332-0020303021030331-2201020200031100) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site.namespace` | [advertise_custom.advertise_where.vk8s_service.virtual_site.namespace](resources--tcp_loadbalancer--reference--group-002.md#canonical-1112303020300010-1203231213133311-2033200023002001-3211030020102031-2333320332030320-3330201013133013-0003122023011321-3231201331220222) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site.tenant` | [advertise_custom.advertise_where.vk8s_service.virtual_site.tenant](resources--tcp_loadbalancer--reference--group-002.md#canonical-2200121221021111-1122220022210202-0222321303323220-1102231331110211-1301000020323311-3232300100011202-3301233232010120-2021130302312302) |
| `advertise_on_public` | [advertise_on_public](resources--tcp_loadbalancer--reference--group-002.md#canonical-0302313131313103-2320203031130213-3133301303203210-2220200131111230-0233011312332320-0230200102303112-2311223301221223-3032232213032211) |
| `advertise_on_public.public_ip` | [advertise_on_public.public_ip](resources--tcp_loadbalancer--reference--group-002.md#canonical-2322221111012111-1133332103010101-0013232322003313-2002023331303131-3211221122003301-3301212102232332-3201101333132303-2211210301321233) |
| `advertise_on_public.public_ip.name` | [advertise_on_public.public_ip.name](resources--tcp_loadbalancer--reference--group-002.md#canonical-0002200333201201-1121312313131333-3222011322012223-2011231033301000-0302211010211332-2232221113001132-1023211223120230-0300200300332131) |
| `advertise_on_public.public_ip.namespace` | [advertise_on_public.public_ip.namespace](resources--tcp_loadbalancer--reference--group-002.md#canonical-1312201323033122-3123210203001233-2320223320220233-2222201201301212-0230301203333032-3310212022000033-3000203121322123-2333333213231000) |
| `advertise_on_public.public_ip.tenant` | [advertise_on_public.public_ip.tenant](resources--tcp_loadbalancer--reference--group-002.md#canonical-1303111231321103-0130122311331131-0210132201310122-0310003201212112-1001120102120013-2222110330131000-2000310133201330-0131022101011333) |
| `advertise_on_public_default_vip` | [advertise_on_public_default_vip](resources--tcp_loadbalancer--reference--group-002.md#canonical-0321202223002312-2000122002021122-1000333311113001-0202202003021011-3030231120013213-3312123301110310-3323231213110121-0311213111231321) |
| `annotations` | [annotations](resources--tcp_loadbalancer--reference--group-001.md#canonical-2131000230031113-0213111121102030-2003021131030103-3123220303022321-0113120030212112-3201102230220033-0230300013302323-2002123223220133) |
| `default_lb_with_sni` | [default_lb_with_sni](resources--tcp_loadbalancer--reference--group-002.md#canonical-0313202033230201-1300210123230312-1001011322302321-0120211103330203-3120232100300130-1120303231202201-1102310202111313-2312111111000220) |
| `description` | [description](resources--tcp_loadbalancer--reference--group-001.md#canonical-0210300032302033-1010012212111101-2300313023100113-1033222123022332-1320011331011221-1101131323010232-3310233030122121-0122030021213023) |
| `disable` | [disable](resources--tcp_loadbalancer--reference--group-001.md#canonical-3211333200201110-3103030023122311-3210331003222030-0221231311222210-1310311121332312-3031311331201132-0221203200022123-0330120013130312) |
| `dns_volterra_managed` | [dns_volterra_managed](resources--tcp_loadbalancer--reference--group-001.md#canonical-1301013030310032-1321231213232210-3131002120333323-1213223112311112-2121332113202211-1303233101201021-0123333210323001-3013010113311202) |
| `do_not_advertise` | [do_not_advertise](resources--tcp_loadbalancer--reference--group-002.md#canonical-3132133212102201-2021300201322133-0333101230320001-0201211013310101-2132321123323201-1233011231122201-0032202131001100-1033022212130331) |
| `do_not_retract_cluster` | [do_not_retract_cluster](resources--tcp_loadbalancer--reference--group-002.md#canonical-2132121202202312-3303103110203131-0123121231302102-1022300131113230-2101310330031313-1031133323302233-2310000223033320-3130011202233233) |
| `domains` | [domains](resources--tcp_loadbalancer--reference--group-001.md#canonical-2311022230003011-3323222012332031-2330233100031023-0230001112122130-3201032212210102-0320003312010303-3032020001331300-2111232131211301) |
| `hash_policy_choice_least_active` | [hash_policy_choice_least_active](resources--tcp_loadbalancer--reference--group-002.md#canonical-2120022010132212-3200121332222303-3312131310001132-1220100002110001-3013320221312220-0031002113013130-1223201110102021-3102010323112103) |
| `hash_policy_choice_random` | [hash_policy_choice_random](resources--tcp_loadbalancer--reference--group-002.md#canonical-0320110201222002-0210223103211030-3332203031121102-3103000311221313-2002102321221001-2330032210323323-3301102223121201-1031013122301231) |
| `hash_policy_choice_round_robin` | [hash_policy_choice_round_robin](resources--tcp_loadbalancer--reference--group-002.md#canonical-3322321013232230-1300300103220333-2220321012300001-3102030202321111-2033300233313210-0220020121331312-2032311221020001-1221330012002212) |
| `hash_policy_choice_source_ip_stickiness` | [hash_policy_choice_source_ip_stickiness](resources--tcp_loadbalancer--reference--group-002.md#canonical-3101023012331020-0102212011231001-1131323230001301-1323130122000230-3202123120312213-3200111331232121-1032310313320003-1122101121103222) |
| `id` | [ID](resources--tcp_loadbalancer--reference--group-001.md#canonical-2231213110320012-1202221203123003-0013222123323322-3031302333332031-3120011310001121-1033013100102011-3321133323010033-2013320030203312) |
| `idle_timeout` | [idle_timeout](resources--tcp_loadbalancer--reference--group-001.md#canonical-0131112213231103-1132301023000303-3213220011211213-2303000030010012-2103121033122102-3000213220103223-2110220113310300-3213310301303312) |
| `labels` | [labels](resources--tcp_loadbalancer--reference--group-001.md#canonical-0332111223013101-3013220013022010-0033233202300100-1023320112113122-3120331020031032-2333222013003333-3122020112320123-1210310312210212) |
| `listen_port` | [listen_port](resources--tcp_loadbalancer--reference--group-001.md#canonical-0103102333022111-2021213011032103-3032103010200230-3202330110112210-3123320133213130-2203103322200013-2322003300231211-3113133010101313) |
| `name` | [name](resources--tcp_loadbalancer--reference--group-001.md#canonical-3220032130201011-1110211113033333-0223302122130231-1311203302233132-3303103312301022-0111301321303011-0211311021133323-1033033032033000) |
| `namespace` | [namespace](resources--tcp_loadbalancer--reference--group-001.md#canonical-0200000111122231-0232213210213212-2103020023032012-2232022333102131-3333130002021332-1012100102113301-3110330013322223-0003302022331330) |
| `no_service_policies` | [no_service_policies](resources--tcp_loadbalancer--reference--group-002.md#canonical-0102302021033021-2301111022111023-2331021210033212-0131133021322033-3000003112132020-3310331033321310-3001103300120102-2322303233312111) |
| `no_sni` | [no_sni](resources--tcp_loadbalancer--reference--group-002.md#canonical-3013302200020113-0111322303322230-0000110230220313-0232032310130333-2323102130310131-1010123111330300-2220231131123000-2211320231132021) |
| `origin_pools_weights` | [origin_pools_weights](resources--tcp_loadbalancer--reference--group-002.md#canonical-3300102231110023-1122311101122321-3133123320300330-3102222002313323-1001120223303033-3223101033122101-0032121010131101-3233033211031232) |
| `origin_pools_weights.cluster` | [origin_pools_weights.cluster](resources--tcp_loadbalancer--reference--group-002.md#canonical-2110022331031320-2323210200030022-2201303121210111-2302002332211320-1202301111330323-3122312123013131-3222123001321121-1202200010031330) |
| `origin_pools_weights.cluster.name` | [origin_pools_weights.cluster.name](resources--tcp_loadbalancer--reference--group-002.md#canonical-2221201022222220-1101232211132331-1211322311311012-0231322230031012-3120232331301301-2201122100233320-1022002012301012-1202111222101233) |
| `origin_pools_weights.cluster.namespace` | [origin_pools_weights.cluster.namespace](resources--tcp_loadbalancer--reference--group-002.md#canonical-1312321133020223-2023020303320302-1113321133221211-3310033222323121-2220011331211020-1223021233121022-2002130332202303-1101110012302002) |
| `origin_pools_weights.cluster.tenant` | [origin_pools_weights.cluster.tenant](resources--tcp_loadbalancer--reference--group-002.md#canonical-0303130010011030-3130300323203021-1212130313322132-2232300322020230-3330311013001331-3111111123120112-3333021020320230-0311201223201221) |
| `origin_pools_weights.endpoint_subsets` | [origin_pools_weights.endpoint_subsets](resources--tcp_loadbalancer--reference--group-002.md#canonical-3120233110133312-0321210100022022-3000321113133310-3130101101021221-1223301332330133-1103223110230230-2230020003011021-0233130133022230) |
| `origin_pools_weights.pool` | [origin_pools_weights.pool](resources--tcp_loadbalancer--reference--group-003.md#canonical-2211110010101203-3132220223100012-1211011020310013-3133233312232003-2030302002310022-0201203300032312-2202322100111100-2111033302210000) |
| `origin_pools_weights.pool.name` | [origin_pools_weights.pool.name](resources--tcp_loadbalancer--reference--group-003.md#canonical-3231133211120113-0210010220020230-3032001333300132-3330221200210310-2003211213221020-2022321330203001-3201322003333331-2230210321333301) |
| `origin_pools_weights.pool.namespace` | [origin_pools_weights.pool.namespace](resources--tcp_loadbalancer--reference--group-003.md#canonical-1311301302321321-1103201112111120-2303010031013133-3031313320321301-2213220113003101-2320121311011330-2210330332320103-3323320323200323) |
| `origin_pools_weights.pool.tenant` | [origin_pools_weights.pool.tenant](resources--tcp_loadbalancer--reference--group-003.md#canonical-0210031323123022-3033220210002201-2122322311203000-2131001200333233-3020321333003202-2100011023013333-1312210001212013-2132103311100202) |
| `origin_pools_weights.priority` | [origin_pools_weights.priority](resources--tcp_loadbalancer--reference--group-002.md#canonical-2023023300223021-1030013322302200-0323020103311210-3121211203123312-0132122310301111-2310333003000230-2320310032131322-3231122110302300) |
| `origin_pools_weights.weight` | [origin_pools_weights.weight](resources--tcp_loadbalancer--reference--group-002.md#canonical-1300023231131310-0000310231101120-1111103312212312-1102222100102213-1113123311303013-1121213322000130-2330000132310232-1100321300303130) |
| `port_ranges` | [port_ranges](resources--tcp_loadbalancer--reference--group-001.md#canonical-2210103003232120-0020010122112210-3121130230322133-1310230033123311-3003032123212233-0300132332322301-1122113023321201-0203320231013212) |
| `retract_cluster` | [retract_cluster](resources--tcp_loadbalancer--reference--group-003.md#canonical-0223212221221323-1003333011230323-2311103300002012-3022112031333333-0211120323003211-0323113322301032-0013033220133100-0321130210120033) |
| `service_policies_from_namespace` | [service_policies_from_namespace](resources--tcp_loadbalancer--reference--group-003.md#canonical-3001223103333333-1020031000231122-0311203031231103-3133312332031032-1302211202001202-3211013112120332-1003020031013122-3233111310220230) |
| `sni` | [sni](resources--tcp_loadbalancer--reference--group-003.md#canonical-2121331132131131-3033031100120302-0330132302022103-3000220102112000-3011311311311201-2200311121231223-3011221000201220-3201310112120322) |
| `tcp` | [tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3320220303000012-1212310221333003-2023232000001311-2003113122201032-0320213100330200-3200012310233331-1203312023032303-2033123000332201) |
| `timeouts` | [timeouts](resources--tcp_loadbalancer--reference--group-003.md#canonical-0231313301332031-0221012333210323-2202220031030022-3300010221320220-0102332011321200-0321222233312130-1133231323310012-1333333313112222) |
| `timeouts.create` | [timeouts.create](resources--tcp_loadbalancer--reference--group-003.md#canonical-1223233012111011-0320010120330323-2110100223101002-0021021131120322-3212003211111103-0020232031013220-3302212121020312-3332323131102132) |
| `timeouts.delete` | [timeouts.delete](resources--tcp_loadbalancer--reference--group-003.md#canonical-2301320011230031-3113330321132003-2132312201210201-3031201321333030-0312012330232203-0201010223302230-2233330222312123-2132112230212103) |
| `timeouts.read` | [timeouts.read](resources--tcp_loadbalancer--reference--group-003.md#canonical-2312030213132213-0021202101111112-2102221101011331-0332303033022122-0303123120222001-1001221110032001-0211033113032320-1010122102000231) |
| `timeouts.update` | [timeouts.update](resources--tcp_loadbalancer--reference--group-003.md#canonical-3323331302032012-3330110330212101-2303232133322130-0223003322013132-2302100202101301-2203312321331313-2011112230130313-1310032331303010) |
| `tls_tcp` | [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3311310101231101-0312312311123203-2112200103233023-0231013101001013-1232132002021323-0011020322221211-1311301030203010-2212103030203323) |
| `tls_tcp.tls_cert_params` | [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-003.md#canonical-2003212231231230-2122303233200032-0123102210223221-2301333103012021-0022212032211012-2002001201000110-3021200222310310-2223313022132223) |
| `tls_tcp.tls_cert_params.certificates` | [tls_tcp.tls_cert_params.certificates](resources--tcp_loadbalancer--reference--group-003.md#canonical-1023331330101033-1020020103122030-2123332131230111-2230101121013212-1023301223011112-2011303033120331-2021112200211310-0113223113132333) |
| `tls_tcp.tls_cert_params.certificates.name` | [tls_tcp.tls_cert_params.certificates.name](resources--tcp_loadbalancer--reference--group-003.md#canonical-0012113110101030-3122010132010113-1313312023003013-1023002333303203-2100120033120011-1230300102221222-0021013032200010-2031321030102233) |
| `tls_tcp.tls_cert_params.certificates.namespace` | [tls_tcp.tls_cert_params.certificates.namespace](resources--tcp_loadbalancer--reference--group-003.md#canonical-3201012123201210-1112131301001222-1133201300123010-0213203332033303-3303030130330123-0320301212130210-1233331300132103-1313122220201313) |
| `tls_tcp.tls_cert_params.certificates.tenant` | [tls_tcp.tls_cert_params.certificates.tenant](resources--tcp_loadbalancer--reference--group-003.md#canonical-3221111031212330-0122002102310230-2301001123303210-3333231002223120-3010123300203330-2320102201202030-2200331320021202-0331020202113010) |
| `tls_tcp.tls_cert_params.no_mtls` | [tls_tcp.tls_cert_params.no_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-0231032331233303-3221202323213112-2122302002313202-0322312221003231-0230200212111233-3222231300120331-0211303122102312-2001133332112332) |
| `tls_tcp.tls_cert_params.tls_config` | [tls_tcp.tls_cert_params.tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-0300232220332032-1221030130223021-3111102102101230-3303020222132332-0221110032000221-1320031310322313-2001122212000022-3332230202330021) |
| `tls_tcp.tls_cert_params.tls_config.custom_security` | [tls_tcp.tls_cert_params.tls_config.custom_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-2113323331002013-1311022023212033-3312022212313032-0122213021312121-1303211030101120-1330002303320230-0032333120113002-2323312212101221) |
| `tls_tcp.tls_cert_params.tls_config.custom_security.cipher_suites` | [tls_tcp.tls_cert_params.tls_config.custom_security.cipher_suites](resources--tcp_loadbalancer--reference--group-003.md#canonical-3121112032303131-2113333330102113-2302002222011123-1321012321213122-2301100112203022-1302000331302101-0021302310222301-1023212320320132) |
| `tls_tcp.tls_cert_params.tls_config.custom_security.max_version` | [tls_tcp.tls_cert_params.tls_config.custom_security.max_version](resources--tcp_loadbalancer--reference--group-003.md#canonical-2213033221303032-2133232233212220-0010120110122031-1131332230001231-1032003012030000-1201202020100303-2102311131322232-1312010102200300) |
| `tls_tcp.tls_cert_params.tls_config.custom_security.min_version` | [tls_tcp.tls_cert_params.tls_config.custom_security.min_version](resources--tcp_loadbalancer--reference--group-003.md#canonical-1121223320123033-2012032320221312-1303211110000233-3332230021321030-2202030111211230-0313330010310102-0221133303120323-0102111222030332) |
| `tls_tcp.tls_cert_params.tls_config.default_security` | [tls_tcp.tls_cert_params.tls_config.default_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-1312101011311300-2121123211331302-0322000312232030-1321002111111312-2211110032231032-0030213020110012-2022321120002130-3313030002300012) |
| `tls_tcp.tls_cert_params.tls_config.low_security` | [tls_tcp.tls_cert_params.tls_config.low_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-1330110313322011-0310203320203031-3121310203122013-0101231031333102-3201221010321121-2311210210011003-2000223310310022-3131021220323320) |
| `tls_tcp.tls_cert_params.tls_config.medium_security` | [tls_tcp.tls_cert_params.tls_config.medium_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-2212022333012322-1311013322202201-0332230112313031-0020113110133021-3200103330203022-1031311301012301-2010030212133332-3100012033212000) |
| `tls_tcp.tls_cert_params.use_mtls` | [tls_tcp.tls_cert_params.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-3111121020110031-3020010003330322-1122023201122301-3322201021112202-1223113202201010-1103102120111200-3022123010301001-1213101302223032) |
| `tls_tcp.tls_cert_params.use_mtls.client_certificate_optional` | [tls_tcp.tls_cert_params.use_mtls.client_certificate_optional](resources--tcp_loadbalancer--reference--group-003.md#canonical-3013320211222232-1113312010311002-2013331322303032-3122101112232213-2121333130330020-2313232131331221-1231201031202110-2000330200313310) |
| `tls_tcp.tls_cert_params.use_mtls.crl` | [tls_tcp.tls_cert_params.use_mtls.crl](resources--tcp_loadbalancer--reference--group-003.md#canonical-0201000232020212-2202311101120211-1303001110313122-1323113011333302-0023300000331133-2101223132022230-2001321002132032-1111132022101323) |
| `tls_tcp.tls_cert_params.use_mtls.crl.name` | [tls_tcp.tls_cert_params.use_mtls.crl.name](resources--tcp_loadbalancer--reference--group-003.md#canonical-0122220321213333-3032112020213112-0000131113302311-0130313203103131-0031312022112301-2313202322130323-2111333000333032-2133223233320111) |
| `tls_tcp.tls_cert_params.use_mtls.crl.namespace` | [tls_tcp.tls_cert_params.use_mtls.crl.namespace](resources--tcp_loadbalancer--reference--group-003.md#canonical-3222211200132001-3031220310033032-1310310032033320-3300220203220201-1103321122010233-1023001202031002-2032003321031202-2223222313132122) |
| `tls_tcp.tls_cert_params.use_mtls.crl.tenant` | [tls_tcp.tls_cert_params.use_mtls.crl.tenant](resources--tcp_loadbalancer--reference--group-003.md#canonical-3202100131322103-0100202301112031-2233311310302301-3323223101332211-3020310300300033-2211113331313103-2311312300112033-3302303033300002) |
| `tls_tcp.tls_cert_params.use_mtls.no_crl` | [tls_tcp.tls_cert_params.use_mtls.no_crl](resources--tcp_loadbalancer--reference--group-003.md#canonical-3321012202302322-2300022233010223-0002002320103100-2201212300331212-2131232302233030-2311321311113112-1100223113313200-2203112310201333) |
| `tls_tcp.tls_cert_params.use_mtls.trusted_ca` | [tls_tcp.tls_cert_params.use_mtls.trusted_ca](resources--tcp_loadbalancer--reference--group-003.md#canonical-2001121202320001-2202231211201110-2003012013201230-0220313302201133-2330001003222200-1012133301012132-2332233323330302-1302031300000212) |
| `tls_tcp.tls_cert_params.use_mtls.trusted_ca.name` | [tls_tcp.tls_cert_params.use_mtls.trusted_ca.name](resources--tcp_loadbalancer--reference--group-003.md#canonical-0030003332002301-3033213103121332-2202100321303020-3032001133100103-1320313301321003-1103233113131113-2120211020010311-0331221112301033) |
| `tls_tcp.tls_cert_params.use_mtls.trusted_ca.namespace` | [tls_tcp.tls_cert_params.use_mtls.trusted_ca.namespace](resources--tcp_loadbalancer--reference--group-003.md#canonical-3033303232333313-3122210031010110-1102013100233333-3221003022213232-1133233010020222-1101122133102123-0320200000011021-1313122332030222) |
| `tls_tcp.tls_cert_params.use_mtls.trusted_ca.tenant` | [tls_tcp.tls_cert_params.use_mtls.trusted_ca.tenant](resources--tcp_loadbalancer--reference--group-003.md#canonical-0133131301311303-0200121022110032-1313302030123033-2000333320301323-2023030002320232-1003302101033332-2112302300111012-0000133021222213) |
| `tls_tcp.tls_cert_params.use_mtls.trusted_ca_url` | [tls_tcp.tls_cert_params.use_mtls.trusted_ca_url](resources--tcp_loadbalancer--reference--group-003.md#canonical-1331110031120010-0133112021222233-3310123203101103-1120030233300330-2102123221100213-2032103221311323-3021201201200321-3002003101100213) |
| `tls_tcp.tls_cert_params.use_mtls.xfcc_disabled` | [tls_tcp.tls_cert_params.use_mtls.xfcc_disabled](resources--tcp_loadbalancer--reference--group-003.md#canonical-1102003200120212-0001032231120020-3010012232313310-2102131301303203-2110213321321002-0103100013232032-1012012330132120-3012111012230230) |
| `tls_tcp.tls_cert_params.use_mtls.xfcc_options` | [tls_tcp.tls_cert_params.use_mtls.xfcc_options](resources--tcp_loadbalancer--reference--group-003.md#canonical-2001133311013203-0331320222002031-1020230000001303-1230320001331322-0033131323332332-2110301023013303-0333023120112033-3333330002310022) |
| `tls_tcp.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements` | [tls_tcp.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements](resources--tcp_loadbalancer--reference--group-003.md#canonical-3110021330311023-2330101103023010-2310200210001001-3330102302120121-1231201313101111-0212112322133232-3112023211112133-0223203312230311) |
| `tls_tcp.tls_parameters` | [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-003.md#canonical-3110131221302012-0020201332212302-2232113113030133-1301112203103333-1123112003212201-2021231011313230-3030310112330201-3023113233130131) |
| `tls_tcp.tls_parameters.no_mtls` | [tls_tcp.tls_parameters.no_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-1101321203110110-2003011002113030-3132001310310233-1003003302213320-0301113232132202-1101023221222301-3222011302000021-2013123121323323) |
| `tls_tcp.tls_parameters.tls_certificates` | [tls_tcp.tls_parameters.tls_certificates](resources--tcp_loadbalancer--reference--group-003.md#canonical-1202231212120100-0230313232012210-0101220333230100-3133010011111033-3100212131003130-2032111121032021-3213120013320200-1223223221323012) |
| `tls_tcp.tls_parameters.tls_certificates.certificate_url` | [tls_tcp.tls_parameters.tls_certificates.certificate_url](resources--tcp_loadbalancer--reference--group-003.md#canonical-3323312300322303-0101303101133320-0113031230010102-0112102212033023-1100213323020213-0233230302230231-2031210032231230-3001200130112212) |
| `tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms` | [tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms](resources--tcp_loadbalancer--reference--group-003.md#canonical-3023032010200202-0123321213302310-0200011320123020-2301033322202011-3221011013213010-3000231003323200-3101222231311010-0100003033232233) |
| `tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms.hash_algorithms` | [tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms.hash_algorithms](resources--tcp_loadbalancer--reference--group-003.md#canonical-2212220120123321-2133031300330331-0111223332223332-3231033131323101-2102203211122021-1000202233110013-3331303113201222-0100103310000211) |
| `tls_tcp.tls_parameters.tls_certificates.description_spec` | [tls_tcp.tls_parameters.tls_certificates.description_spec](resources--tcp_loadbalancer--reference--group-003.md#canonical-2213300101033130-3302203302213113-0322332331210012-0201321200032210-3020333310031000-0133011302021313-1233002210320112-3113323033122330) |
| `tls_tcp.tls_parameters.tls_certificates.disable_ocsp_stapling` | [tls_tcp.tls_parameters.tls_certificates.disable_ocsp_stapling](resources--tcp_loadbalancer--reference--group-003.md#canonical-3023311011102113-3021223201133112-1323211121000232-0303210323200303-2330130013033010-3002032023100001-0212323313231232-2121232130122333) |
| `tls_tcp.tls_parameters.tls_certificates.private_key` | [tls_tcp.tls_parameters.tls_certificates.private_key](resources--tcp_loadbalancer--reference--group-003.md#canonical-1123011310203113-3021011221132332-1033103301311113-1000100031023220-1333112013232012-0000112221132213-3322033303103231-2223231110110032) |
| `tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info` | [tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info](resources--tcp_loadbalancer--reference--group-003.md#canonical-3332301012313030-0020232102210020-3000311111333222-2133112010222120-3222012301032002-0310232101232002-3323230131122013-2331232301032132) |
| `tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info.decryption_provider](resources--tcp_loadbalancer--reference--group-003.md#canonical-2012210230101302-0132203223022232-1220023022312111-2032112132230330-2301222033320111-0330020313110321-0200003003200313-0011213000113331) |
| `tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location` | [tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location](resources--tcp_loadbalancer--reference--group-003.md#canonical-3303301233102000-2312323022220232-0223230133231013-2101201011133311-3113222303012022-0013323103231121-3121230313121131-3011300303220132) |
| `tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info.store_provider` | [tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info.store_provider](resources--tcp_loadbalancer--reference--group-003.md#canonical-1023331223321223-0021203233231033-3110103223203020-0303221210102323-1232203131220130-2113000202030023-1020013013113100-2031022321022321) |
| `tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info` | [tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info](resources--tcp_loadbalancer--reference--group-003.md#canonical-1121332121001303-1330031030003332-2332023200211220-1113111101212011-0113332002311202-2312203002010102-3000023000101023-1120111233020003) |
| `tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info.provider_ref` | [tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info.provider_ref](resources--tcp_loadbalancer--reference--group-003.md#canonical-0102311301232332-0233123210030330-2203000022130033-2032202210202330-0331102020232103-3313133111031201-0203030310332121-3210013321223232) |
| `tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info.url` | [tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info.url](resources--tcp_loadbalancer--reference--group-003.md#canonical-0130202101332222-3202202301103120-0212010100131101-1120020103112031-3033201300100321-2003300301101320-1102130200313011-3131223233102310) |
| `tls_tcp.tls_parameters.tls_certificates.use_system_defaults` | [tls_tcp.tls_parameters.tls_certificates.use_system_defaults](resources--tcp_loadbalancer--reference--group-003.md#canonical-3321332112202331-0332220020203310-0303033222211031-3111032001010021-3022330202330210-0230302202001231-1203010221201321-3320220322313121) |
| `tls_tcp.tls_parameters.tls_config` | [tls_tcp.tls_parameters.tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-0212300030220000-1030102030120032-3230012030010030-0000021233232103-0220012112012103-0232213020123322-2232021021001302-0012002033001322) |
| `tls_tcp.tls_parameters.tls_config.custom_security` | [tls_tcp.tls_parameters.tls_config.custom_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-2101203230103003-0022303032202231-0131223223111012-3233012312103213-1001023021033001-2100022031011221-1203102001323321-1301100001220121) |
| `tls_tcp.tls_parameters.tls_config.custom_security.cipher_suites` | [tls_tcp.tls_parameters.tls_config.custom_security.cipher_suites](resources--tcp_loadbalancer--reference--group-003.md#canonical-2120111230032212-3301001021200020-1023133233032212-1110303201110011-2021022230221201-1131020200030133-2223232121220313-2103311220122113) |
| `tls_tcp.tls_parameters.tls_config.custom_security.max_version` | [tls_tcp.tls_parameters.tls_config.custom_security.max_version](resources--tcp_loadbalancer--reference--group-003.md#canonical-1031333132300100-0133123202123002-0000101321001333-2031122131003303-2022221001220200-0233300233000321-2332323322132110-1313021210333320) |
| `tls_tcp.tls_parameters.tls_config.custom_security.min_version` | [tls_tcp.tls_parameters.tls_config.custom_security.min_version](resources--tcp_loadbalancer--reference--group-003.md#canonical-0332013010102312-0233000311211020-2331000222113201-0120313002200032-1213122300031232-1003222000210023-3122233313232021-0310303231010021) |
| `tls_tcp.tls_parameters.tls_config.default_security` | [tls_tcp.tls_parameters.tls_config.default_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-2100100011331032-1320201311013202-1220313213333021-0132130000330331-3020102312322220-2201223031321111-3113031332123203-3321301231210332) |
| `tls_tcp.tls_parameters.tls_config.low_security` | [tls_tcp.tls_parameters.tls_config.low_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-1221332221122311-3130011130300221-1202023201232230-3013132212022033-2022330030303302-2230311211331020-3013323230032002-2121232033311301) |
| `tls_tcp.tls_parameters.tls_config.medium_security` | [tls_tcp.tls_parameters.tls_config.medium_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-1020031022132103-0201111302022203-3133231013000123-1122132022031103-2230013201200003-1031232100303201-1001212120302120-0321332000232002) |
| `tls_tcp.tls_parameters.use_mtls` | [tls_tcp.tls_parameters.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-1232221302231213-0031210310023113-0211110311201223-1120002312213120-2011321300110012-3200030013001222-2100103323222132-3302122312012030) |
| `tls_tcp.tls_parameters.use_mtls.client_certificate_optional` | [tls_tcp.tls_parameters.use_mtls.client_certificate_optional](resources--tcp_loadbalancer--reference--group-003.md#canonical-3221111013201111-3213023210020303-3131030321232001-3203102023003103-0102001010302210-2012103231122331-1203201220231100-3221211132120123) |
| `tls_tcp.tls_parameters.use_mtls.crl` | [tls_tcp.tls_parameters.use_mtls.crl](resources--tcp_loadbalancer--reference--group-003.md#canonical-2231002300022321-0310320020013321-3013231001302330-0222022201020003-1333312021222020-0123002203132031-3221023102230103-2323210312230300) |
| `tls_tcp.tls_parameters.use_mtls.crl.name` | [tls_tcp.tls_parameters.use_mtls.crl.name](resources--tcp_loadbalancer--reference--group-003.md#canonical-0230012313210303-0033010223133013-0300200112331132-0321122221223033-3132301012030211-3302203013101313-3231320322233230-3000032003213223) |
| `tls_tcp.tls_parameters.use_mtls.crl.namespace` | [tls_tcp.tls_parameters.use_mtls.crl.namespace](resources--tcp_loadbalancer--reference--group-003.md#canonical-0113221301022030-1221222321132301-0010123121103202-0300021231303212-0212320221200121-1223111003120310-3321311201112000-2313202122122223) |
| `tls_tcp.tls_parameters.use_mtls.crl.tenant` | [tls_tcp.tls_parameters.use_mtls.crl.tenant](resources--tcp_loadbalancer--reference--group-003.md#canonical-3321103023302322-1130302103302300-3232302121212231-3213320201031033-3232213100110303-2101303022312333-2103112303122232-1220320212122000) |
| `tls_tcp.tls_parameters.use_mtls.no_crl` | [tls_tcp.tls_parameters.use_mtls.no_crl](resources--tcp_loadbalancer--reference--group-003.md#canonical-3130121333221110-2303030212130233-0101213010223232-3212112032320032-3310033333332110-0330323323320023-1122000222002203-2230112223222001) |
| `tls_tcp.tls_parameters.use_mtls.trusted_ca` | [tls_tcp.tls_parameters.use_mtls.trusted_ca](resources--tcp_loadbalancer--reference--group-003.md#canonical-0223202131300212-0310312222130113-3132132013321133-1113010120111302-1112023223102000-3113203211000312-1300033232131000-2303101302210311) |
| `tls_tcp.tls_parameters.use_mtls.trusted_ca.name` | [tls_tcp.tls_parameters.use_mtls.trusted_ca.name](resources--tcp_loadbalancer--reference--group-003.md#canonical-2212212023121202-2313023330321321-0231321121110211-3123301023001322-3130220200231021-3212313322310221-3313202210101101-1232031212223231) |
| `tls_tcp.tls_parameters.use_mtls.trusted_ca.namespace` | [tls_tcp.tls_parameters.use_mtls.trusted_ca.namespace](resources--tcp_loadbalancer--reference--group-003.md#canonical-2230130301232230-3001011121023102-2302011103023023-2003121212313023-1110102030102030-2003031323222233-0120002222012302-1220033330133113) |
| `tls_tcp.tls_parameters.use_mtls.trusted_ca.tenant` | [tls_tcp.tls_parameters.use_mtls.trusted_ca.tenant](resources--tcp_loadbalancer--reference--group-003.md#canonical-0130123031211303-0212032231230031-3231110103101211-0001310013220021-3032222201300001-2001313132300123-2200003031320000-3221223311320333) |
| `tls_tcp.tls_parameters.use_mtls.trusted_ca_url` | [tls_tcp.tls_parameters.use_mtls.trusted_ca_url](resources--tcp_loadbalancer--reference--group-003.md#canonical-2003211222130021-0102013022031220-0311111202133301-1223000212300330-0120033121121310-2102321021313102-1300102312222321-2321210202130323) |
| `tls_tcp.tls_parameters.use_mtls.xfcc_disabled` | [tls_tcp.tls_parameters.use_mtls.xfcc_disabled](resources--tcp_loadbalancer--reference--group-003.md#canonical-1301301113111110-2233321110301200-3201113132333301-3123023200120333-2130030330321230-3130101213131301-3320033210001100-2302023131330223) |
| `tls_tcp.tls_parameters.use_mtls.xfcc_options` | [tls_tcp.tls_parameters.use_mtls.xfcc_options](resources--tcp_loadbalancer--reference--group-003.md#canonical-1232011003223131-1231312010203002-0320300222000333-2020333200020212-0013010110113301-0321220123313213-3331330222001311-0121200312220102) |
| `tls_tcp.tls_parameters.use_mtls.xfcc_options.xfcc_header_elements` | [tls_tcp.tls_parameters.use_mtls.xfcc_options.xfcc_header_elements](resources--tcp_loadbalancer--reference--group-003.md#canonical-1231010103323111-0133022222121131-3031321113032032-2113331320323301-0010101132333022-1312231130002100-1331200300323310-3001332033133021) |
| `tls_tcp_auto_cert` | [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-003.md#canonical-1021002212100302-1311013322023113-0033102003303321-1102223200321033-3111102121313232-0020131233022011-3121111200002222-0320203122133203) |
| `tls_tcp_auto_cert.no_mtls` | [tls_tcp_auto_cert.no_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-2313120012003131-0311100200113110-0123033033123200-3003200132300220-1021000011003221-3200000101330201-2110012131301202-3132133220230013) |
| `tls_tcp_auto_cert.tls_config` | [tls_tcp_auto_cert.tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-3321310201213312-1102103322302112-2213212202200032-3302231101100330-0213221333111211-1023311110133123-0020313120313112-1300130202230102) |
| `tls_tcp_auto_cert.tls_config.custom_security` | [tls_tcp_auto_cert.tls_config.custom_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-3122320210213120-3132310300203321-0023203133132221-2012132122221032-1132021010222230-3110333322211003-1200110210233200-3202113110210333) |
| `tls_tcp_auto_cert.tls_config.custom_security.cipher_suites` | [tls_tcp_auto_cert.tls_config.custom_security.cipher_suites](resources--tcp_loadbalancer--reference--group-003.md#canonical-3110230210310221-1011323300032322-0113000130312010-2022231123113023-1022302023122221-3111300133033221-3101321133322100-3233221210211200) |
| `tls_tcp_auto_cert.tls_config.custom_security.max_version` | [tls_tcp_auto_cert.tls_config.custom_security.max_version](resources--tcp_loadbalancer--reference--group-003.md#canonical-3302230203202020-1020012130110030-2020113300330120-3012211330212032-2212103232133231-1230220210132333-3301023232002323-1023002132323012) |
| `tls_tcp_auto_cert.tls_config.custom_security.min_version` | [tls_tcp_auto_cert.tls_config.custom_security.min_version](resources--tcp_loadbalancer--reference--group-003.md#canonical-2331012331102330-2222311331122201-3102103312030101-0320011333220032-0103132130020020-1210103003231030-0300130312113113-3122112322122300) |
| `tls_tcp_auto_cert.tls_config.default_security` | [tls_tcp_auto_cert.tls_config.default_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-2010010302012332-2230323211012202-0312230201203102-3300001130302322-1330311131021002-0113302200111113-2122213232321110-2211120031233012) |
| `tls_tcp_auto_cert.tls_config.low_security` | [tls_tcp_auto_cert.tls_config.low_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-0012331301303113-0122013031120210-3111133330002022-1223232130010220-0213102121003312-3333303311320231-2330320030021323-1132012003211232) |
| `tls_tcp_auto_cert.tls_config.medium_security` | [tls_tcp_auto_cert.tls_config.medium_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-2203331123101331-3200220000230021-0112033122002013-0333232221023131-1003012123322012-1001322221220201-2000000113100103-2111033323100312) |
| `tls_tcp_auto_cert.use_mtls` | [tls_tcp_auto_cert.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-2122123310312001-1332302032133120-0112220022001113-0332021122030022-0331230110000112-3330102331101112-1023032300031022-0133033333131032) |
| `tls_tcp_auto_cert.use_mtls.client_certificate_optional` | [tls_tcp_auto_cert.use_mtls.client_certificate_optional](resources--tcp_loadbalancer--reference--group-003.md#canonical-1021332202221031-0011322033110003-0000322201110132-2312203201221221-1310102321101111-2323221212331321-3000212001013023-3303203103330332) |
| `tls_tcp_auto_cert.use_mtls.crl` | [tls_tcp_auto_cert.use_mtls.crl](resources--tcp_loadbalancer--reference--group-003.md#canonical-1020103222230313-0322312213133322-0121301323231031-0302030102111222-0300002203120120-0010103310112202-3130101220001322-3000220213230313) |
| `tls_tcp_auto_cert.use_mtls.crl.name` | [tls_tcp_auto_cert.use_mtls.crl.name](resources--tcp_loadbalancer--reference--group-003.md#canonical-2022111101313010-0313322233121013-1331001202300120-1223110330130112-1030131213103012-2132110103021201-3013121013320201-0010010332220303) |
| `tls_tcp_auto_cert.use_mtls.crl.namespace` | [tls_tcp_auto_cert.use_mtls.crl.namespace](resources--tcp_loadbalancer--reference--group-003.md#canonical-2123213131033113-0101303223000211-1112011200112121-3110332222321100-3213022001232231-1300122121202230-0012311230210033-3333321322223111) |
| `tls_tcp_auto_cert.use_mtls.crl.tenant` | [tls_tcp_auto_cert.use_mtls.crl.tenant](resources--tcp_loadbalancer--reference--group-003.md#canonical-3023300012323011-3112212113020112-3203202020332121-3202100302232031-3000323012331313-1211002031231033-1323023311223000-3001221322120300) |
| `tls_tcp_auto_cert.use_mtls.no_crl` | [tls_tcp_auto_cert.use_mtls.no_crl](resources--tcp_loadbalancer--reference--group-003.md#canonical-1331203211100310-2320012102132103-0030010000223320-3300333111012011-2010310131300000-1010323123033112-1010210100203101-2310320223223302) |
| `tls_tcp_auto_cert.use_mtls.trusted_ca` | [tls_tcp_auto_cert.use_mtls.trusted_ca](resources--tcp_loadbalancer--reference--group-003.md#canonical-0310300202211220-2130103020100302-0102210122321330-2030310221102210-3211111311311132-3303103211003233-1323202302101323-1103012012131011) |
| `tls_tcp_auto_cert.use_mtls.trusted_ca.name` | [tls_tcp_auto_cert.use_mtls.trusted_ca.name](resources--tcp_loadbalancer--reference--group-003.md#canonical-0101112133310323-1100032200311133-2011111120122012-0120033102303020-1103302330232031-0323101331133123-0210322302301331-0113113013302222) |
| `tls_tcp_auto_cert.use_mtls.trusted_ca.namespace` | [tls_tcp_auto_cert.use_mtls.trusted_ca.namespace](resources--tcp_loadbalancer--reference--group-003.md#canonical-3100333102033203-1230122020112301-3010023213303000-1120310003030230-3131012112101221-3333023301131013-0101220110332130-0130312003203320) |
| `tls_tcp_auto_cert.use_mtls.trusted_ca.tenant` | [tls_tcp_auto_cert.use_mtls.trusted_ca.tenant](resources--tcp_loadbalancer--reference--group-003.md#canonical-3132002233303301-3102011330212212-1201132302202201-0213033231101332-0111010120002031-0330003031010021-3213110011101303-1011031101020301) |
| `tls_tcp_auto_cert.use_mtls.trusted_ca_url` | [tls_tcp_auto_cert.use_mtls.trusted_ca_url](resources--tcp_loadbalancer--reference--group-003.md#canonical-1301233321023301-3221301123322013-3120111200100111-3213131121132220-1322131002130010-1021133020131331-2222210230130311-1313302230033120) |
| `tls_tcp_auto_cert.use_mtls.xfcc_disabled` | [tls_tcp_auto_cert.use_mtls.xfcc_disabled](resources--tcp_loadbalancer--reference--group-003.md#canonical-0013103230120133-2032003020333120-1113103020030311-3300020222213333-1202230333110320-1210013302101223-1323322211113010-3321320103021301) |
| `tls_tcp_auto_cert.use_mtls.xfcc_options` | [tls_tcp_auto_cert.use_mtls.xfcc_options](resources--tcp_loadbalancer--reference--group-003.md#canonical-2213032130123000-1133101312211231-0130100322231020-3133332321101210-2202113113101213-2133221323233300-3123203031201133-1322320303110102) |
| `tls_tcp_auto_cert.use_mtls.xfcc_options.xfcc_header_elements` | [tls_tcp_auto_cert.use_mtls.xfcc_options.xfcc_header_elements](resources--tcp_loadbalancer--reference--group-003.md#canonical-1122302311011201-3233002131113212-0111022023223311-3002001100333310-2132111310012313-0201110102321230-2332103233021031-1231211232113132) |

<a id="canonical-2123332010300020-1023231311030232-1300031122101011-1032023203210220-1200333010322303-0323133123222003-1020130133223133-1112131030203102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `active_service_policies` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- active_service_policies

<a id="canonical-1211213311100203-3013023320103303-3223003322203320-0331213010022010-0001333301100210-0022302033003002-2302133212310230-3112000003210013"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: active\_service\_policies, no\_service\_policies, service\_policies\_from\_namespace;
Default: no\_service\_policies\] Configuration parameter for active service policies.

Additional upstream details:

List of service policies.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("policies")}
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

OneOf alternatives in this subsection:

- [active_service_policies](resources--tcp_loadbalancer--reference--group-001.md#canonical-1211213311100203-3013023320103303-3223003322203320-0331213010022010-0001333301100210-0022302033003002-2302133212310230-3112000003210013)
- [no_service_policies](resources--tcp_loadbalancer--reference--group-002.md#canonical-0102302021033021-2301111022111023-2331021210033212-0131133021322033-3000003112132020-3310331033321310-3001103300120102-2322303233312111)
- [service_policies_from_namespace](resources--tcp_loadbalancer--reference--group-003.md#canonical-3001223103333333-1020031000231122-0311203031231103-3133312332031032-1302211202001202-3211013112120332-1003020031013122-3233111310220230)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
active_service_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-1131121310130023-2021302123133302-0101221021011013-0322021001000121-2112012221201101-0200303011200100-2102311132131132-3310100213310201"></a>

### Direct properties for `active_service_policies`

- [policies](resources--tcp_loadbalancer--reference--group-001.md#canonical-3131001221313221-0211331023022001-0110003203230011-1312001122213133-2111103232033312-0120333101021323-2203101020232232-3010122220001001): complete subsection reference.

<a id="canonical-3131001221313221-0211331023022001-0110003203230011-1312001122213133-2111103232033312-0120333101021323-2203101020232232-3010122220001001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `active_service_policies.policies` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [active_service_policies](resources--tcp_loadbalancer--reference--group-001.md#canonical-2123332010300020-1023231311030232-1300031122101011-1032023203210220-1200333010322303-0323133123222003-1020130133223133-1112131030203102)
- active_service_policies.policies

<a id="canonical-1021220010111222-0012031311020213-0212322031031001-2302322202223003-0023101312211013-1323013313031122-1021102000213020-3032233133313022"></a>

Type: `"object"`. list nested block, Optional.

Service Policies is a sequential engine where policies (and rules within the policy) are evaluated
one after the other. It's important to define the correct order (policies evaluated from top to
bottom in the list) for service policies, to GET the intended result. For each request, its
characteristics are evaluated based on the match criteria in each service policy starting at the
top. If there is a match in the current policy, then the policy takes effect, and no more policies
are evaluated. Otherwise, the next policy is evaluated. If all policies are evaluated and none
match, then the request will be denied by default.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
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

Terraform syntax:

```terraform
policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-1012201013233032-3003130030021001-3001330111233113-1220013021100101-2231201102102011-1323122301230112-2230211201130332-2012211201020313"></a>

### Direct properties for `active_service_policies.policies`

<a id="canonical-3302102322212223-0220012013110110-3111012133113323-3202233121111113-3313103212223113-2310001323123201-3311030311110132-1131020303313020"></a>

#### `active_service_policies.policies.name` property

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

<a id="canonical-1003233022033031-1030322201031222-3033132222330123-2233223202130130-1301012022301010-0303322113300322-1200112210010010-3212333223002110"></a>

<a id="canonical-2311111121320201-2103000002110330-0311103020210023-2320213000102303-3232302003011200-0321111011022333-1311313101322203-2033023033030300"></a>

#### `active_service_policies.policies.namespace` property

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

<a id="canonical-2111310323203122-2030220313001320-2100222123223211-3320102230102102-1312313123002322-3033210322301002-0012221213302131-3000012301202300"></a>

<a id="canonical-0102312131221231-0221110213030030-0301222211031113-2211333022020010-3231100030221322-2020223321103202-3203103231322212-2011220231133131"></a>

#### `active_service_policies.policies.tenant` property

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

<a id="canonical-0101013130323102-0111020001013232-1223003210233021-3112202320110020-1310022023311310-1003000113120300-2312202022202013-0222311212101130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- advertise_custom

<a id="canonical-1301323022031230-1003003001020123-1132201332123232-2311223310032312-1132021010021312-1231021033101001-1131013213121211-2103230021312203"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: advertise\_custom, advertise\_on\_public, advertise\_on\_public\_default\_vip,
do\_not\_advertise; Default: advertise\_on\_public\_default\_vip\] Defines a way to advertise a VIP
on specific sites.

Additional upstream details:

This defines a way to advertise a VIP on specific sites.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("advertise_where")}
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

OneOf alternatives in this subsection:

- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-1301323022031230-1003003001020123-1132201332123232-2311223310032312-1132021010021312-1231021033101001-1131013213121211-2103230021312203)
- [advertise_on_public](resources--tcp_loadbalancer--reference--group-002.md#canonical-0302313131313103-2320203031130213-3133301303203210-2220200131111230-0233011312332320-0230200102303112-2311223301221223-3032232213032211)
- [advertise_on_public_default_vip](resources--tcp_loadbalancer--reference--group-002.md#canonical-0321202223002312-2000122002021122-1000333311113001-0202202003021011-3030231120013213-3312123301110310-3323231213110121-0311213111231321)
- [do_not_advertise](resources--tcp_loadbalancer--reference--group-002.md#canonical-3132133212102201-2021300201322133-0333101230320001-0201211013310101-2132321123323201-1233011231122201-0032202131001100-1033022212130331)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
advertise_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-1102332303000001-3222020323221233-3023203310211300-2012030013112002-3023332321212220-2232100302210003-3110310131321003-0023103103213133"></a>

### Direct properties for `advertise_custom`

- [advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-2220031130203102-2210011003013002-3311212001302200-0330220102220213-0312033302222203-3301311320102322-2002200311110023-0013020221203332): complete subsection reference.

<a id="canonical-2220031130203102-2210011003013002-3311212001302200-0330220102220213-0312033302222203-3301311320102322-2002200311110023-0013020221203332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-0101013130323102-0111020001013232-1223003210233021-3112202320110020-1310022023311310-1003000113120300-2312202022202013-0222311212101130)
- advertise_custom.advertise_where

<a id="canonical-3131003031113023-0220010230102313-0300110222112030-3300113232132233-3331103203103011-0320002201322122-1212023321021000-3111233322220223"></a>

Type: `"object"`. list nested block, Optional.

Where should this load balancer be available.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "advertise_on_public"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "advertise_v6_on_public"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "site"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "advertise_v6_on_public"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "site"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "site"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("port",
    "port_ranges"),
  validators.ConflictingListObjectAttributes("port",
    "use_default_port"),
  validators.ConflictingListObjectAttributes("port_ranges",
    "use_default_port"),
  validators.ConflictingListObjectAttributes("site",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("site",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("site",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("site",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("virtual_network",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("virtual_network",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("virtual_network",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("virtual_site",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("virtual_site",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("virtual_site_with_vip",
    "vk8s_service")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
advertise_where {
  # Configure direct properties listed below.
}
```

<a id="canonical-3322010112302121-1132013211222322-3212311312331332-0200113021121130-0131233320130001-1112100332011333-2101013123101301-0032131120303023"></a>

### Direct properties for `advertise_custom.advertise_where`

- [advertise_dualstack_on_public](resources--tcp_loadbalancer--reference--group-001.md#canonical-3021201332333301-1221103113230301-2222130122300123-1303203313031312-2113012101013220-3333202101332112-3112113102220123-3003103203310001): complete subsection reference.

- [advertise_on_public](resources--tcp_loadbalancer--reference--group-001.md#canonical-3313033122011113-0012321000120221-3030301223212102-0322311032301001-1312203032330232-1233032011232103-2302012030122212-0020323220220130): complete subsection reference.

- [advertise_v6_on_public](resources--tcp_loadbalancer--reference--group-001.md#canonical-0020010320102332-2030221110332020-0302003031113211-1211232110302323-2133310300202132-3131311232123211-3120121031321210-1201000300021230): complete subsection reference.

<a id="canonical-2002133223111111-1332301233021023-2023320112110320-0132101131320300-1222322222313100-3222021031112003-3122113122212320-0301310021002200"></a>

<a id="canonical-3021023303012010-0331122133322200-1212230330221002-1203301233223010-3112211201211131-2022312110103020-3031131010312022-1132333122231333"></a>

#### `advertise_custom.advertise_where.port` property

Type: `"number"`. Optional.

Exclusive with \[port\_ranges use\_default\_port\] Port to Listen.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2020233021133322-1223231322123313-0013211212333111-3122032101023223-0122112102132110-1011000101200201-3101130220022100-3122220301122023"></a>

<a id="canonical-0103202121011220-2003111231313303-3220123331333002-1120022203221222-2102011233211230-3130330130112012-0221330212102213-3222023112102001"></a>

#### `advertise_custom.advertise_where.port_ranges` property

Type: `"string"`. Optional.

Exclusive with \[port use\_default\_port\] A string containing a comma separated list of port
ranges. Each port range consists of a single port or two ports separated by '-'.

Additional upstream details:

Each port range consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

- [site](resources--tcp_loadbalancer--reference--group-001.md#canonical-0212220231301321-1302012333301320-1033203121232320-2312031111320321-1011312223000220-2301211003310223-0322223221300313-1111000033023321): complete subsection reference.

- [use_default_port](resources--tcp_loadbalancer--reference--group-001.md#canonical-1101203011303313-3221210002023210-0302033021031331-2121122302332230-3323030202321000-3001331232322320-3330220021022113-1030311210112112): complete subsection reference.

- [virtual_network](resources--tcp_loadbalancer--reference--group-001.md#canonical-1003112313011332-2020131311011001-0030223322103102-1202300132320101-1111022212122021-1030301120012112-0331021223030333-0330020330031020): complete subsection reference.

- [virtual_site](resources--tcp_loadbalancer--reference--group-002.md#canonical-1122231030213111-1302221130230031-0331311330131011-0113323220002101-1200331300203123-1110330313212100-3233313332101212-2310321002101013): complete subsection reference.

- [virtual_site_with_vip](resources--tcp_loadbalancer--reference--group-002.md#canonical-1121201333003001-3211103001030232-1120020013233002-1201212112111322-2313032201232300-2313132333210321-3020122003000202-2203220100301301): complete subsection reference.

- [vk8s_service](resources--tcp_loadbalancer--reference--group-002.md#canonical-2310032102311332-1320010320220300-3000131113132003-3031300330001303-3011312212131020-2233012333120212-1031220012331120-3102210330032022): complete subsection reference.

<a id="canonical-3021201332333301-1221103113230301-2222130122300123-1303203313031312-2113012101013220-3333202101332112-3112113102220123-3003103203310001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.advertise_dualstack_on_public` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-0101013130323102-0111020001013232-1223003210233021-3112202320110020-1310022023311310-1003000113120300-2312202022202013-0222311212101130)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-2220031130203102-2210011003013002-3311212001302200-0330220102220213-0312033302222203-3301311320102322-2002200311110023-0013020221203332)
- advertise_custom.advertise_where.advertise_dualstack_on_public

<a id="canonical-0221013131212033-0013120323321102-3323220031131213-0222130200222221-0313231022200022-1301131202233120-2032213103011301-1210211312102100"></a>

Type: `"object"`. single nested block, Optional.

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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
advertise_dualstack_on_public {
  # Configure direct properties listed below.
}
```

<a id="canonical-3301321123020033-1200133023321331-1030300021212023-2100313120103202-3200031311110020-3002130001001011-3202100032322102-3330320202311222"></a>

### Direct properties for `advertise_custom.advertise_where.advertise_dualstack_on_public`

- [public_ip](resources--tcp_loadbalancer--reference--group-001.md#canonical-2302312312303212-3133301300110213-2212211221023002-2212112003010020-2322330212112320-0032013303322111-1102103121032303-0000120110201313): complete subsection reference.

<a id="canonical-2302312312303212-3133301300110213-2212211221023002-2212112003010020-2322330212112320-0032013303322111-1102103121032303-0000120110201313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-0101013130323102-0111020001013232-1223003210233021-3112202320110020-1310022023311310-1003000113120300-2312202022202013-0222311212101130)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-2220031130203102-2210011003013002-3311212001302200-0330220102220213-0312033302222203-3301311320102322-2002200311110023-0013020221203332)
- [advertise_custom.advertise_where.advertise_dualstack_on_public](resources--tcp_loadbalancer--reference--group-001.md#canonical-3021201332333301-1221103113230301-2222130122300123-1303203313031312-2113012101013220-3333202101332112-3112113102220123-3003103203310001)
- advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip

<a id="canonical-3332001223333303-3233000200022303-2321323123102312-0022303021030000-1021310312230322-0322113012112313-1330103032002000-1333133110121111"></a>

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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-1202010001103223-2302103232221131-3011202003222332-0220032033200013-0311022300023331-1002021231120010-3303133211230100-0300120210223130"></a>

### Direct properties for `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip`

<a id="canonical-0221322202223201-1123100322330313-3231133311311312-0221021122230232-1123221302020202-0011011332210010-0202331311303202-3033311202212310"></a>

#### `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name` property

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

<a id="canonical-3323102333211120-0101210012201003-1032323201330203-1221331133232232-3322103100033303-3032103302303030-3033200031331220-0003310000200202"></a>

<a id="canonical-2000211013333132-1323321213002222-3312111233002130-2020120023030213-1332001101222120-2201110032311033-2023003031332333-0133113012330230"></a>

#### `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace` property

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

<a id="canonical-0313202121202003-1200320220033121-1132201322020102-2023101031023232-0123102323011122-0121233030020222-1000310330221213-2321023232233002"></a>

<a id="canonical-3122102321000322-3112030302200022-3230101002322011-0012323032332112-0112022103300323-0302322022200323-1213000230133313-2103200130121211"></a>

#### `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant` property

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

<a id="canonical-3313033122011113-0012321000120221-3030301223212102-0322311032301001-1312203032330232-1233032011232103-2302012030122212-0020323220220130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.advertise_on_public` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-0101013130323102-0111020001013232-1223003210233021-3112202320110020-1310022023311310-1003000113120300-2312202022202013-0222311212101130)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-2220031130203102-2210011003013002-3311212001302200-0330220102220213-0312033302222203-3301311320102322-2002200311110023-0013020221203332)
- advertise_custom.advertise_where.advertise_on_public

<a id="canonical-1201200023021201-0022030201023003-0103320033131331-0000301102200231-3001022032020300-3322212200133032-1321112001233010-0323231311113011"></a>

Type: `"object"`. single nested block, Optional.

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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
advertise_on_public {
  # Configure direct properties listed below.
}
```

<a id="canonical-3132232311320030-0230022123332103-2221330031111303-3133201223012102-3101012333030133-3202032120320213-2321023200333333-3200220101231000"></a>

### Direct properties for `advertise_custom.advertise_where.advertise_on_public`

- [public_ip](resources--tcp_loadbalancer--reference--group-001.md#canonical-0121132110213323-3121330201102232-2300130311101220-3121032120210331-2122312032032232-3331211232221233-1213102221212133-3203022130100213): complete subsection reference.

<a id="canonical-0121132110213323-3121330201102232-2300130311101220-3121032120210331-2122312032032232-3331211232221233-1213102221212133-3203022130100213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.advertise_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-0101013130323102-0111020001013232-1223003210233021-3112202320110020-1310022023311310-1003000113120300-2312202022202013-0222311212101130)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-2220031130203102-2210011003013002-3311212001302200-0330220102220213-0312033302222203-3301311320102322-2002200311110023-0013020221203332)
- [advertise_custom.advertise_where.advertise_on_public](resources--tcp_loadbalancer--reference--group-001.md#canonical-3313033122011113-0012321000120221-3030301223212102-0322311032301001-1312203032330232-1233032011232103-2302012030122212-0020323220220130)
- advertise_custom.advertise_where.advertise_on_public.public_ip

<a id="canonical-0332233033123312-0102121003221012-3222221121020122-0022020210311233-3232202002303010-2130110133003003-3030030103200003-2231121111132312"></a>

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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-2120222122032020-0131002201120323-3203322010112210-0223210222022123-2230001111000112-3322320002312220-1020111213132201-2022011023323302"></a>

### Direct properties for `advertise_custom.advertise_where.advertise_on_public.public_ip`

<a id="canonical-3300313210230301-0332223103002122-0311202001021020-0202020033202022-1010202003133022-1003203121122110-2110101113323103-2323321320231003"></a>

#### `advertise_custom.advertise_where.advertise_on_public.public_ip.name` property

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

<a id="canonical-2231213000202020-3002220300131232-3011131123013301-1202102012211022-1332312311233223-1233322230233110-2323023033020120-3123221101013032"></a>

<a id="canonical-1122333132022032-1011112013313022-2113233202002300-3022101323233013-0233203033203111-2020301330322223-0321102220211230-1123130333122300"></a>

#### `advertise_custom.advertise_where.advertise_on_public.public_ip.namespace` property

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

<a id="canonical-1203002101333331-0311103033111033-0223230232212203-2333012003102112-3200110333121003-2132232223320131-0313222321002102-3202022220331121"></a>

<a id="canonical-1332021030031220-0302120002230333-1202232101301031-0313333313033220-1111203002312211-2121113112112021-3222020101013221-3123133233120100"></a>

#### `advertise_custom.advertise_where.advertise_on_public.public_ip.tenant` property

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

<a id="canonical-0020010320102332-2030221110332020-0302003031113211-1211232110302323-2133310300202132-3131311232123211-3120121031321210-1201000300021230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.advertise_v6_on_public` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-0101013130323102-0111020001013232-1223003210233021-3112202320110020-1310022023311310-1003000113120300-2312202022202013-0222311212101130)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-2220031130203102-2210011003013002-3311212001302200-0330220102220213-0312033302222203-3301311320102322-2002200311110023-0013020221203332)
- advertise_custom.advertise_where.advertise_v6_on_public

<a id="canonical-2303333110033321-0232231022302200-2111310303132201-2102232021330023-2200131322010331-0101020022101002-1133232131120010-1223330011311310"></a>

Type: `"object"`. single nested block, Optional.

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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
advertise_v6_on_public {
  # Configure direct properties listed below.
}
```

<a id="canonical-0031330221122003-1211132303110011-2332303020120332-1021033030320230-0231231133331012-0311212013333220-1233311211330332-1112111323133001"></a>

### Direct properties for `advertise_custom.advertise_where.advertise_v6_on_public`

- [public_ip](resources--tcp_loadbalancer--reference--group-001.md#canonical-0013123222000011-2102200221113321-2212131130321231-1103213322233220-1212121101221010-2300021213130011-1000113000130132-1011112310302222): complete subsection reference.

<a id="canonical-0013123222000011-2102200221113321-2212131130321231-1103213322233220-1212121101221010-2300021213130011-1000113000130132-1011112310302222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.advertise_v6_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-0101013130323102-0111020001013232-1223003210233021-3112202320110020-1310022023311310-1003000113120300-2312202022202013-0222311212101130)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-2220031130203102-2210011003013002-3311212001302200-0330220102220213-0312033302222203-3301311320102322-2002200311110023-0013020221203332)
- [advertise_custom.advertise_where.advertise_v6_on_public](resources--tcp_loadbalancer--reference--group-001.md#canonical-0020010320102332-2030221110332020-0302003031113211-1211232110302323-2133310300202132-3131311232123211-3120121031321210-1201000300021230)
- advertise_custom.advertise_where.advertise_v6_on_public.public_ip

<a id="canonical-0012012132323332-1022112231133122-1200032003331230-1113113012222130-2211203331020323-0000321232333311-1003220212031020-0021121120312310"></a>

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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-3122302230200001-1112113132130330-1022220032001311-2000122203112010-1031211130212010-3303023330003021-3201232010220223-3213313121320010"></a>

### Direct properties for `advertise_custom.advertise_where.advertise_v6_on_public.public_ip`

<a id="canonical-1313133112303321-3313300322202013-3030231212010212-0110101300113302-2133303021100311-2323031022033111-2023112313111301-0020022022320211"></a>

#### `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name` property

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

<a id="canonical-2333012203131021-2231333212012003-2220103302332302-1023123111320211-2133332130133321-1010310201023112-1210022322132320-0331023111131201"></a>

<a id="canonical-2003320212033221-2302132122120112-2031303313131112-3012311320322303-0312320021212221-2103233202133220-3120132100002030-1223123303321321"></a>

#### `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace` property

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

<a id="canonical-0110320130203131-3103330310223131-0021011322231100-1232131013013012-0223100001111230-3121320321200320-3100033130133321-1311321302213131"></a>

<a id="canonical-0232011000221310-2123303000230013-3000111013210021-2133323223203220-3232230232033032-0102203211023030-3320101311213100-0211332222032100"></a>

#### `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant` property

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

<a id="canonical-0212220231301321-1302012333301320-1033203121232320-2312031111320321-1011312223000220-2301211003310223-0322223221300313-1111000033023321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.site` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-0101013130323102-0111020001013232-1223003210233021-3112202320110020-1310022023311310-1003000113120300-2312202022202013-0222311212101130)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-2220031130203102-2210011003013002-3311212001302200-0330220102220213-0312033302222203-3301311320102322-2002200311110023-0013020221203332)
- advertise_custom.advertise_where.site

<a id="canonical-0233002212112233-3111330321132031-3121222333110322-3021330031300232-1303220232320133-0011323303112302-3032101121311222-1110223323033002"></a>

Type: `"object"`. single nested block, Optional.

This defines a reference to a CE site along with network type and an optional IP address where a
load balancer could be advertised.

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-1302303103200302-0332211330233210-2031301223323213-1011233101013011-3320131031100000-2223132103130021-2121200230103131-3302033022302212"></a>

### Direct properties for `advertise_custom.advertise_where.site`

<a id="canonical-1113230300201002-1332212231203013-3010303033032200-3233112012302130-3032011322120312-1222100231330122-0210122212300312-2310030320331022"></a>

#### `advertise_custom.advertise_where.site.ip` property

Type: `"string"`. Optional.

Use given IP address as VIP on the site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-3011111001110212-1023211021302000-2211212223201003-1103100111331111-2000310332023020-3121110310133200-2023320121220011-2230231312223101"></a>

<a id="canonical-0031303011312220-0220231333102021-3102322021333120-0120321121111211-2230033101001310-2000320202112012-3111230220100333-3301020220213032"></a>

#### `advertise_custom.advertise_where.site.network` property

Type: `"string"`. Optional.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Additional upstream details:

This defines network types to be used on site

All inside and outside networks. All outside networks. All outside networks with internet VIP
support. VK8s service network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the
site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_INSIDE_AND_OUTSIDE",
  "enum": [
    "SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [site](resources--tcp_loadbalancer--reference--group-001.md#canonical-2000212032131230-3031223031320310-3321022200321323-0112030200221113-2110233200133223-2101230233220331-2111000321222102-1031313203321122): complete subsection reference.

<a id="canonical-2000212032131230-3031223031320310-3321022200321323-0112030200221113-2110233200133223-2101230233220331-2111000321222102-1031313203321122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.site.site` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-0101013130323102-0111020001013232-1223003210233021-3112202320110020-1310022023311310-1003000113120300-2312202022202013-0222311212101130)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-2220031130203102-2210011003013002-3311212001302200-0330220102220213-0312033302222203-3301311320102322-2002200311110023-0013020221203332)
- [advertise_custom.advertise_where.site](resources--tcp_loadbalancer--reference--group-001.md#canonical-0212220231301321-1302012333301320-1033203121232320-2312031111320321-1011312223000220-2301211003310223-0322223221300313-1111000033023321)
- advertise_custom.advertise_where.site.site

<a id="canonical-0333313201001022-2313031213321231-0132021101002003-3230023031113000-3120220323200230-0231030202030323-1113111301131021-1333031320113202"></a>

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-2210103023310030-1020303332102103-3312302222233223-3103232210213110-1312201022121010-1023023033221102-0312300223013332-0110030110320221"></a>

### Direct properties for `advertise_custom.advertise_where.site.site`

<a id="canonical-1210232312302010-1302132201332010-1312211222003320-3100311111101220-2311132320321202-2313233201230211-2022110212303031-0322300213200133"></a>

#### `advertise_custom.advertise_where.site.site.name` property

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

<a id="canonical-2033122031331020-0100100121330230-3210002110302121-2013122231122233-2231223131310330-1003313112033312-1300213020002000-1223010232332120"></a>

<a id="canonical-2232233020012221-3111133333010313-0032032102020231-0111231033002110-2110313320200221-2003322322001001-0322131302203221-2303322100231123"></a>

#### `advertise_custom.advertise_where.site.site.namespace` property

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

<a id="canonical-3113113123012313-0313130312223311-1012221013103101-1122211320033332-3000303220103300-1001301130011202-3323100001120222-2102302122302230"></a>

<a id="canonical-0032110211210302-1332203311323310-1301120033130301-1200131113131330-2213230110023230-1020300112112110-2220222232120333-1200020303202013"></a>

#### `advertise_custom.advertise_where.site.site.tenant` property

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

<a id="canonical-1101203011303313-3221210002023210-0302033021031331-2121122302332230-3323030202321000-3001331232322320-3330220021022113-1030311210112112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.use_default_port` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-0101013130323102-0111020001013232-1223003210233021-3112202320110020-1310022023311310-1003000113120300-2312202022202013-0222311212101130)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-2220031130203102-2210011003013002-3311212001302200-0330220102220213-0312033302222203-3301311320102322-2002200311110023-0013020221203332)
- advertise_custom.advertise_where.use_default_port

<a id="canonical-2120120222220122-3323012000120222-2202102233331233-2101022230221031-0201321302033102-3300022313330233-1200131130212300-2311310103212210"></a>

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
use_default_port = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1003112313011332-2020131311011001-0030223322103102-1202300132320101-1111022212122021-1030301120012112-0331021223030333-0330020330031020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.virtual_network` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-0101013130323102-0111020001013232-1223003210233021-3112202320110020-1310022023311310-1003000113120300-2312202022202013-0222311212101130)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-2220031130203102-2210011003013002-3311212001302200-0330220102220213-0312033302222203-3301311320102322-2002200311110023-0013020221203332)
- advertise_custom.advertise_where.virtual_network

<a id="canonical-2203310231222103-1020202323302322-3300011300002121-2100232333130220-2302211013332332-2321310111100322-3013310130031021-2023131333112201"></a>

Type: `"object"`. single nested block, Optional.

Parameters to advertise on a given virtual network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_v6_vip",
    "specific_v6_vip"),
  validators.ConflictingObjectAttributes("default_vip",
    "specific_vip")}
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
  "x-ves-oneof-field-v6_vip_choice": "[\"default_v6_vip\",\"specific_v6_vip\"]",
  "x-ves-oneof-field-vip_choice": "[\"default_vip\",\"specific_vip\"]"
}
```

Terraform syntax:

```terraform
virtual_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-1232313233223330-1023113310120113-3030232333333001-0030322313121302-0032123212133023-1000311031032310-0133220012101322-2110330030332032"></a>

### Direct properties for `advertise_custom.advertise_where.virtual_network`

- [default_v6_vip](resources--tcp_loadbalancer--reference--group-002.md#canonical-3022220111330023-2033301113121011-0103101011332303-2132012011023202-2013032210001032-3030011120200301-1033300322022213-2033331112303313): complete subsection reference.

- [default_vip](resources--tcp_loadbalancer--reference--group-002.md#canonical-2032130230213230-3303113002020202-3030213303311100-1301120021221311-2201300123303031-2332211032000302-1000323111133012-2212221120012222): complete subsection reference.

<a id="canonical-0021222232011222-1023122310013311-1120212012210322-3212002123023303-1000220300330333-0022100330301332-2330301010203333-3022032031232230"></a>
