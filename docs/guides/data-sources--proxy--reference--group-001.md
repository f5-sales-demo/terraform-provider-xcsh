---
page_title: "xcsh_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_proxy reference."
---

# xcsh_proxy reference

<a id="canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- Property reference

<a id="canonical-1330313230123033-2110103332221031-3323232132122022-2230033213330030-2300021221012011-0301321112001230-3203310313203020-0111001111032311"></a>

### Direct properties for `xcsh_proxy`

- [active_forward_proxy_policies](data-sources--proxy--reference--group-001.md#canonical-2213113013300313-0020111122011302-2231233220020000-0031023111301110-2222323311112113-1032001030003100-3301333013321011-0123202302110231): complete subsection reference.

<a id="canonical-1003302031101000-0100000130100220-1103323331003101-0330000333301003-1231210320013021-3011302210122010-2223300220031302-2321132213101030"></a>

<a id="canonical-0220201233310010-0302122003000320-0203233211222132-1303210111003300-0102102313200303-0322213010300120-3022012011131022-0021313113002003"></a>

#### `annotations` property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Additional upstream details:

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

<a id="canonical-2300331113031003-2012023230322023-1122310003203301-1232202132022011-1300131312031002-1020111123222232-1001030201230121-0012112131001332"></a>

<a id="canonical-0030333330311333-2222120013312132-3122122331223111-0203011202000232-3323101130321320-1323131022321230-3101003031210102-3202200311332032"></a>

#### `connection_timeout` property

Type: `"number"`. Computed.

The timeout for new network connections to upstream server. This is specified in milliseconds. The
(2 seconds). Defaults to \`2000\`.

Additional upstream details:

The default value is 2000 (2 seconds)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1800000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1800000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1800000"
  }
}
```

<a id="canonical-3120233003021023-0230132001302011-0100331223311131-1100002120312221-1122332003303020-1301310323232332-1123123201221333-2310320233120021"></a>

<a id="canonical-2110213312323020-1221330031201011-2232322120322230-2303323102131222-1000221211322212-1111031031330003-1132031013330331-3322210033103222"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the Proxy.

Additional upstream details:

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [do_not_advertise](data-sources--proxy--reference--group-001.md#canonical-1222310002221222-1321223032311031-2020310323203003-1210330021211320-0013020220203303-3203312313122011-3000002100223313-0101333020130003): complete subsection reference.

- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313): complete subsection reference.

- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302): complete subsection reference.

<a id="canonical-1130221001032033-1030030112333213-1123022030330001-1311301012302002-2120010301130031-3133033212133020-0130201113230032-1303202110011102"></a>

<a id="canonical-2313320221033130-3021133313120323-2122311132020331-2210330121011113-2302120033122122-2022112312312303-1320332021102203-2000012132232032"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1213130031000122-3022010201211232-0323311333201333-1210322012230330-3132323121100221-0203132212232103-3033013311102232-0110021230113013"></a>

<a id="canonical-2323030332323300-1213000222210121-2123101123220111-0302331030333201-2130222210332303-3113022203332033-3233213123112111-1022032202232031"></a>

#### `labels` property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

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

<a id="canonical-3103233300113031-0132101231131312-0031021021302033-0310201110201033-0001301210003222-3012303230123133-1133003313010131-1210312013301311"></a>

<a id="canonical-3321221232022332-3202300010111021-2110213111201321-0033200231332210-1130130210023231-3222111321321023-0230332012322321-0033221200301100"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Proxy.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1323121121310331-1300013312012021-0001100133232222-2033223032200103-2331112110110111-1332000313033131-2320330301301013-1223312101031122"></a>

<a id="canonical-0300201312332312-0300031320130333-0130303201101012-3023223210001120-0023201330320100-0003232002132321-3022113210021301-2330031232013320"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Proxy exists.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [no_forward_proxy_policy](data-sources--proxy--reference--group-004.md#canonical-0311103232311212-0122021311132121-3312130333133131-0323020121313100-3121300110200320-2101201223212331-0111013021211303-2001132132123021): complete subsection reference.

- [no_interception](data-sources--proxy--reference--group-004.md#canonical-1030230220130100-0131133033331303-0222210130321320-0031101211030302-3010333010202023-1003003201301120-2030130231203121-0133333310330322): complete subsection reference.

- [site_local_inside_network](data-sources--proxy--reference--group-005.md#canonical-1101012223103130-2201131330221131-1331022102330232-3330000112313011-2223113210330001-3012032001012200-0303100022020121-1102012131121322): complete subsection reference.

- [site_local_network](data-sources--proxy--reference--group-005.md#canonical-1313130232332212-2111301322113102-2223321320232221-1220210033111003-1130311211020203-2230312001121313-0110231202002231-1133102331220011): complete subsection reference.

- [site_virtual_sites](data-sources--proxy--reference--group-005.md#canonical-1012201322333030-2201110212121013-2012102101130230-0200303100213333-0202312022013300-2020101021021321-1320201313130100-3300013202221330): complete subsection reference.

- [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312): complete subsection reference.

<a id="canonical-0203011232302311-3211301130120310-2000323123132003-0223312012233033-1330101033310032-0123123221211031-3003303223301000-1333332023210021"></a>

### All schema paths for `xcsh_proxy`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `active_forward_proxy_policies` | [active_forward_proxy_policies](data-sources--proxy--reference--group-001.md#canonical-0113112101100330-3012113212033210-2300302031303022-0103202033010002-2001100310212330-3321202003130323-1203132322032202-2300022112133033) |
| `active_forward_proxy_policies.forward_proxy_policies` | [active_forward_proxy_policies.forward_proxy_policies](data-sources--proxy--reference--group-001.md#canonical-3223202303233120-3303312123211202-3230113310312121-1202000222113312-0220031320133220-1202132101012222-1223330032222033-3020232113130020) |
| `active_forward_proxy_policies.forward_proxy_policies.name` | [active_forward_proxy_policies.forward_proxy_policies.name](data-sources--proxy--reference--group-001.md#canonical-0033222202020033-0222320011332203-0301111221233023-3313203333112103-2010300223312221-0300012033032113-3003102112322220-0010110301213120) |
| `active_forward_proxy_policies.forward_proxy_policies.namespace` | [active_forward_proxy_policies.forward_proxy_policies.namespace](data-sources--proxy--reference--group-001.md#canonical-0103000020310112-2213133132031210-2100321012310113-3131100323330323-2010032111102200-1231113131113000-3033302303000000-2233102120210110) |
| `active_forward_proxy_policies.forward_proxy_policies.tenant` | [active_forward_proxy_policies.forward_proxy_policies.tenant](data-sources--proxy--reference--group-001.md#canonical-1311100302230330-0032213030132301-2312233330311202-3232121001012230-1021320212330202-0033200032030223-0113221013123310-0130021032011232) |
| `annotations` | [annotations](data-sources--proxy--reference--group-001.md#canonical-1003302031101000-0100000130100220-1103323331003101-0330000333301003-1231210320013021-3011302210122010-2223300220031302-2321132213101030) |
| `connection_timeout` | [connection_timeout](data-sources--proxy--reference--group-001.md#canonical-2300331113031003-2012023230322023-1122310003203301-1232202132022011-1300131312031002-1020111123222232-1001030201230121-0012112131001332) |
| `description` | [description](data-sources--proxy--reference--group-001.md#canonical-3120233003021023-0230132001302011-0100331223311131-1100002120312221-1122332003303020-1301310323232332-1123123201221333-2310320233120021) |
| `do_not_advertise` | [do_not_advertise](data-sources--proxy--reference--group-001.md#canonical-1111013301201012-2330232112312001-2103313002312231-0111011120033031-0013111131103130-3330111233202133-2120132303021132-2203322221101333) |
| `dynamic_proxy` | [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3300012110333323-0313303121121121-1331003000322330-3313312302103220-3210020033323300-0212112321033131-0003101301121210-3311300223113220) |
| `dynamic_proxy.disable_dns_masquerade` | [dynamic_proxy.disable_dns_masquerade](data-sources--proxy--reference--group-001.md#canonical-3030020103222030-0001110010222311-0220032123002311-3222122011000332-3031313220303211-1021032221320220-3332132130301120-3230332303122111) |
| `dynamic_proxy.domains` | [dynamic_proxy.domains](data-sources--proxy--reference--group-001.md#canonical-2102010211002230-1012113212113010-0230200110221020-0322311033302313-0111002122110103-1110032011212313-0210101211333210-3201030101120223) |
| `dynamic_proxy.enable_dns_masquerade` | [dynamic_proxy.enable_dns_masquerade](data-sources--proxy--reference--group-001.md#canonical-2202230330010222-1201110002013011-1212113012210011-2200032131101320-3011012201022202-2033103202322312-0113100331230101-1333012213003120) |
| `dynamic_proxy.http_proxy` | [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-1300103332011102-3110211000323203-3312121232322222-2123330223011231-2120001322320111-3200231200233213-0220313023303020-2021202211232301) |
| `dynamic_proxy.http_proxy.more_option` | [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-3101311213023200-2211010131221101-1300310131110022-2101131101100121-1303203223302122-0130123012333100-3011103221300133-3322122013013011) |
| `dynamic_proxy.http_proxy.more_option.buffer_policy` | [dynamic_proxy.http_proxy.more_option.buffer_policy](data-sources--proxy--reference--group-001.md#canonical-0222223111112122-0311010333221231-1330112033312323-1110032213002322-1233012130211221-3202122112303132-0233112330220011-2232221003320031) |
| `dynamic_proxy.http_proxy.more_option.buffer_policy.disabled` | [dynamic_proxy.http_proxy.more_option.buffer_policy.disabled](data-sources--proxy--reference--group-001.md#canonical-2302003030211222-3302133203311203-2111030211113223-1220221232130320-1201202101211023-3121231322010003-1332011030330010-1130102122300220) |
| `dynamic_proxy.http_proxy.more_option.buffer_policy.max_request_bytes` | [dynamic_proxy.http_proxy.more_option.buffer_policy.max_request_bytes](data-sources--proxy--reference--group-001.md#canonical-1322201233103223-3212321023013011-0201020010201022-3010300320111222-0301332113200312-0320011022031213-1331131110320202-1122130021321003) |
| `dynamic_proxy.http_proxy.more_option.compression_params` | [dynamic_proxy.http_proxy.more_option.compression_params](data-sources--proxy--reference--group-001.md#canonical-3131231122122120-1311012232200231-0013223201213033-3132203112113013-2121102012312220-3122200200021102-1011311212113201-3221302133332203) |
| `dynamic_proxy.http_proxy.more_option.compression_params.content_length` | [dynamic_proxy.http_proxy.more_option.compression_params.content_length](data-sources--proxy--reference--group-001.md#canonical-0120222323323200-0002230030123330-3102212131121033-0303221031221031-1010312321220102-1212121320312021-3022222131323313-1100031021021230) |
| `dynamic_proxy.http_proxy.more_option.compression_params.content_type` | [dynamic_proxy.http_proxy.more_option.compression_params.content_type](data-sources--proxy--reference--group-001.md#canonical-1203112102021003-1103012013331100-0212122112121301-2220302130233030-1122120203332323-3013311121202100-2300213011230321-1010301221303231) |
| `dynamic_proxy.http_proxy.more_option.compression_params.disable_on_etag_header` | [dynamic_proxy.http_proxy.more_option.compression_params.disable_on_etag_header](data-sources--proxy--reference--group-001.md#canonical-1302031213300313-1013113230021212-3331032202031310-2001023212300000-2310210233111123-0312010313110220-0331322010220030-3122202102103033) |
| `dynamic_proxy.http_proxy.more_option.compression_params.remove_accept_encoding_header` | [dynamic_proxy.http_proxy.more_option.compression_params.remove_accept_encoding_header](data-sources--proxy--reference--group-001.md#canonical-2111101123033202-0011333333221022-0030311000103111-3203303233133113-2331220323230102-1121303311210103-1201223120331231-2110002012303222) |
| `dynamic_proxy.http_proxy.more_option.custom_errors` | [dynamic_proxy.http_proxy.more_option.custom_errors](data-sources--proxy--reference--group-001.md#canonical-1022133221122202-1212331330033211-1312211000001103-3013202201300323-1110300122111311-2030301113012030-0200301033331310-2333013322312331) |
| `dynamic_proxy.http_proxy.more_option.disable_default_error_pages` | [dynamic_proxy.http_proxy.more_option.disable_default_error_pages](data-sources--proxy--reference--group-001.md#canonical-2232323331212222-1112021203322001-2302002102312111-3313213131032032-1102030103113102-3013222012221001-3102012323313230-3333323232113210) |
| `dynamic_proxy.http_proxy.more_option.disable_path_normalize` | [dynamic_proxy.http_proxy.more_option.disable_path_normalize](data-sources--proxy--reference--group-001.md#canonical-1113233131312122-3130102222100323-3231320311313003-1310221021330023-3030222122111001-2101022103200131-0013322303312130-3131013210211301) |
| `dynamic_proxy.http_proxy.more_option.enable_path_normalize` | [dynamic_proxy.http_proxy.more_option.enable_path_normalize](data-sources--proxy--reference--group-001.md#canonical-1013200221322220-2303203122231100-2220211203312111-3131012311313011-3032010320121222-3032020223110100-1022222233312332-1212000331030330) |
| `dynamic_proxy.http_proxy.more_option.idle_timeout` | [dynamic_proxy.http_proxy.more_option.idle_timeout](data-sources--proxy--reference--group-001.md#canonical-3211123313011203-1321001221312312-1121232100011220-1220210030002232-2201201232001122-2103232002030230-2030023013232331-0201012332003023) |
| `dynamic_proxy.http_proxy.more_option.max_request_header_size` | [dynamic_proxy.http_proxy.more_option.max_request_header_size](data-sources--proxy--reference--group-001.md#canonical-0331032133232121-2233003012131223-0032012131212123-1200312020112211-3303310202321202-1331132102323103-0120020122310320-3030223202221113) |
| `dynamic_proxy.http_proxy.more_option.max_requests_per_connection` | [dynamic_proxy.http_proxy.more_option.max_requests_per_connection](data-sources--proxy--reference--group-001.md#canonical-2301223333323001-0022112210023021-3222012003101100-3300223111212312-0033032103001232-2011200212020211-3211111301210223-2302231002232221) |
| `dynamic_proxy.http_proxy.more_option.no_request_limit_per_connection` | [dynamic_proxy.http_proxy.more_option.no_request_limit_per_connection](data-sources--proxy--reference--group-001.md#canonical-2000323201012320-3100103022020310-3123303320001201-1030021200210023-1030001230021201-2030100323302131-3122122130100233-0310212333023031) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add](data-sources--proxy--reference--group-001.md#canonical-2130320311301030-1232110232001001-2303031323322313-3101213313002222-2320230031302332-1233131000130212-3312300322001333-2131311311101133) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.name` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.name](data-sources--proxy--reference--group-001.md#canonical-1232112232231320-2103131202200113-1020211021103210-0201230001230332-1032322021233103-2300222032120032-3230322313032102-0001123012122012) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.overwrite` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.overwrite](data-sources--proxy--reference--group-001.md#canonical-3311101103220112-1002132233203333-1203021102013101-0202222313221000-0232120033101211-3030330310102031-3233121212313222-3333023122011032) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value](data-sources--proxy--reference--group-001.md#canonical-0333230303230001-1233300210233331-3020210231321132-1031311122212001-1220232330130111-2000032233002001-1213012030001221-1033331012303011) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info](data-sources--proxy--reference--group-002.md#canonical-1211303312220022-1312022323120323-3011211303333311-3132111010121232-0113300330031023-3223321133003011-2110301210112110-1110313202320011) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--proxy--reference--group-002.md#canonical-3303311232111323-3222113001012120-3032223032120130-3212102210010203-0220203220211303-0131100132233333-0212100021322100-1312302103002131) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location](data-sources--proxy--reference--group-002.md#canonical-1223030023101231-3112022210032323-0102122112113111-0200221231203023-0320230011231121-3100300030120021-2031302313331323-2200130002221321) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--proxy--reference--group-002.md#canonical-3210113310033300-2322002110333102-0022212201033331-0102133331203003-1010023310233312-2303210113322233-0131101332023020-2031030211000200) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info](data-sources--proxy--reference--group-002.md#canonical-2310323101130232-1133301231011132-2032103022123020-2022112230133001-3003200011223031-3212120120233020-2132021223301100-2011033121212132) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref](data-sources--proxy--reference--group-002.md#canonical-0312331223033230-0211003030101323-3131213131002022-2222333230102330-3313231300230201-3103333121201112-1301202210332330-2131103103310111) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.url](data-sources--proxy--reference--group-002.md#canonical-2323300032213010-3121211031212221-3120330123120222-1100002232200201-3002331033103313-2013110012232003-1122023222312003-1001002320200312) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.value` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.value](data-sources--proxy--reference--group-001.md#canonical-3023030233021220-0202311202300332-1033012333012113-1103322301333322-3320000023031230-0202323203200020-3323023101023000-2112333212010202) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_remove` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_remove](data-sources--proxy--reference--group-001.md#canonical-2010112220311330-3221310033202300-3333321003312022-2203033030000021-2222201312101102-1023101233003131-3013231111030233-3210221101332331) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add](data-sources--proxy--reference--group-002.md#canonical-3033212330002301-2223112230232320-3303200103112333-1112230022013013-0313213311013121-0302131013331302-3222000332122222-3101001032101222) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.append` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.append](data-sources--proxy--reference--group-002.md#canonical-2321122302230302-2013133233200101-1133331021120212-2230323313212122-0031213222121201-1231321303101310-3003003030102231-0113021320311011) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.name` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.name](data-sources--proxy--reference--group-002.md#canonical-2320132322211121-2102333302210233-3323020111200302-2221023103003132-0223110202030021-0310133113131321-0003022223030302-3333323203303331) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-2130230302102212-1132001110323302-0203313122101123-3030122312133130-0132330203001310-1201103130313330-0002021201121002-3313230321123000) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info](data-sources--proxy--reference--group-002.md#canonical-0230110211222232-0300211231021233-1333102222010201-2313310000200202-1003333101033102-3313230113030221-0101313331212221-3001112321010023) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--proxy--reference--group-002.md#canonical-3001023131031010-2231330200312133-2322101230312331-3011320223022223-0000030033333020-0120001031010132-3320333303121212-0132312230200200) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.location](data-sources--proxy--reference--group-002.md#canonical-1311311131030100-1210002313201333-0112333221213020-0210112201211020-3321312032333331-2012130112011321-2033102121330320-1100202011331101) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--proxy--reference--group-002.md#canonical-0202003202012110-0213223010031310-2213300200100313-3310302010322121-1231232130010303-3201110010122131-0002121023213222-0003210021231212) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info](data-sources--proxy--reference--group-002.md#canonical-1330030230210000-3221210120301102-2103012100032121-2020030333101001-0230203232011213-1222330311320303-3112130302111012-3303322031123201) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref](data-sources--proxy--reference--group-002.md#canonical-1312023111303332-0313231320121023-0031221020031331-0332001120210012-1312122113220000-0013030221311233-3103203123002232-3112200113321200) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.url](data-sources--proxy--reference--group-002.md#canonical-3330131200223123-0333131311021013-0111121223230110-2200201312111333-1113331120203120-2033302312131222-0311321303202233-2033320011212321) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.value` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.value](data-sources--proxy--reference--group-002.md#canonical-1223303110101311-2313013223303000-1121030211303130-1203110311322321-3003021220110333-3110301222221223-0131331131332302-0021122331201302) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_remove` | [dynamic_proxy.http_proxy.more_option.request_headers_to_remove](data-sources--proxy--reference--group-001.md#canonical-1233012331102201-0203001131001223-1313011002111222-0113121222210203-2200102223231332-1002011031121333-0011211233313133-1301100330020101) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-2303221113023010-3131012220121131-0330120301220120-3100110331223003-0303233132033021-0112220111212022-1232323201333112-0213302133110011) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_domain` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_domain](data-sources--proxy--reference--group-002.md#canonical-3023013221131033-3202201310301300-2332010032102102-1233031200300201-1023132003100011-1100022102322232-2121201331101322-2132221033322223) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_expiry` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_expiry](data-sources--proxy--reference--group-002.md#canonical-1033211023122220-0131120210001303-2322021230222232-0303111203013313-1231321031332130-3131313023012223-1203312000221322-3100331313110031) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_httponly` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_httponly](data-sources--proxy--reference--group-002.md#canonical-0310122130232112-2112303020323010-3031023021201200-1130110131122123-2201132012111011-3113013103333323-0001203003201012-3021230300123112) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_partitioned` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_partitioned](data-sources--proxy--reference--group-002.md#canonical-0110302311211213-3001303232113032-0003032011123132-0133302231331331-1310310101301212-0000213221332023-1032033122213310-2313230122302223) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_path` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_path](data-sources--proxy--reference--group-002.md#canonical-2003222133030332-1123332022110010-3112233213021133-0123312131313012-3121212120131112-3002201333103213-2232000001100111-0020003212303110) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_secure` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_secure](data-sources--proxy--reference--group-002.md#canonical-0220131211131002-0332312121332013-1330233133101112-3220330332333000-3313302101130201-2300300120112221-3100012203003200-2331120132030221) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_domain` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_domain](data-sources--proxy--reference--group-002.md#canonical-2101021113223202-0123220102122120-0123221031303212-1312220102021220-1202223020210220-0002020000313002-3131103022312331-2012233001131102) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_expiry` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_expiry](data-sources--proxy--reference--group-002.md#canonical-0032021002112131-0003113112020131-3120301302320101-3213313323300300-2112223001122211-0221323123012330-0011020113213020-1120030031203032) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_httponly` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_httponly](data-sources--proxy--reference--group-002.md#canonical-2313311020230312-2220012210312121-3313322021232321-3200112231012013-1301123033111102-3232203120130100-2332301213231102-0301232300021000) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_max_age` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_max_age](data-sources--proxy--reference--group-002.md#canonical-1321012033301332-3302212312232303-0122211230023002-0121111130121222-2101103012213111-1110031012203111-3213200211133000-3120021013232121) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_partitioned` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_partitioned](data-sources--proxy--reference--group-002.md#canonical-2230013130210020-0003131033111010-2310230123223133-0231213112133200-2022021213120021-1330003213230231-2122120301321300-0012032221232201) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_path` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_path](data-sources--proxy--reference--group-002.md#canonical-0201110230001312-3212013101310213-2320110011202212-1101100113212332-2211121302011023-1021120111000321-3212010100200332-2031301200300102) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_samesite` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_samesite](data-sources--proxy--reference--group-002.md#canonical-2330110303223023-3313302133211022-0233030220312310-1103112330203223-1221312131230132-1221100210112131-1003123023100332-0312111312203230) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_secure` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_secure](data-sources--proxy--reference--group-002.md#canonical-3303201032311210-2020021000200200-2220232002301232-0211113322233021-0230201133112211-1010131022001332-2132330011231011-1010101312013003) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_value` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_value](data-sources--proxy--reference--group-002.md#canonical-2301101020300023-0111213200000021-1323000000033301-3310331102012233-0030213212133212-1232010331201013-3131013233000003-2131332310003122) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.max_age_value` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.max_age_value](data-sources--proxy--reference--group-002.md#canonical-2003031311121321-1223131020202121-0110213100303031-1331131133103000-1301122311021202-1103101100111030-3001022331020023-2331302301222022) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.name` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.name](data-sources--proxy--reference--group-002.md#canonical-3313201223020301-0203303212322000-1113201022203120-1130012021020013-0123312312002312-0130012321230022-0231121311001222-1100023122333101) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.overwrite` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.overwrite](data-sources--proxy--reference--group-002.md#canonical-0110222213003303-2020220123133310-2210311110223312-2013032112032001-1002110300330130-3323031320230111-2320202113021320-3300222311203012) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_lax` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_lax](data-sources--proxy--reference--group-002.md#canonical-2202132320013103-3010012212231032-2101223300130022-1001213131103112-2231202333201030-3121003213311320-2320300021323001-0303013133230313) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_none` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_none](data-sources--proxy--reference--group-002.md#canonical-1023311230313012-0120313221202230-3332212131221231-0112320210311300-3300003001002203-0023023123010100-3001311022130311-1301220013212001) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_strict` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_strict](data-sources--proxy--reference--group-002.md#canonical-2322131320020103-2300130103020322-3113133131332023-0031120230120110-3223233320103111-3231103232021121-3203332032130100-2133221212003303) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-3012002322003201-0231300023000001-0302103313033300-3213113012213023-1223311310222202-1201302023030212-2202102002320333-2111121210023111) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info](data-sources--proxy--reference--group-002.md#canonical-0100220301033201-3112201120003031-3212023010131120-1011112212222101-2010320023022232-2223030002300112-3131303121213203-3031131200011312) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--proxy--reference--group-002.md#canonical-0321222212300231-2212221100131023-1233200021032232-0102231033011212-0113122232020030-2320020003123132-3022200320233231-0011312231213003) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.location](data-sources--proxy--reference--group-002.md#canonical-1000033231102332-3331101220100303-2231220202321201-2312310233111211-0002023333111332-1133101321323022-3312313133310231-1201112101020100) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--proxy--reference--group-002.md#canonical-3222223232013300-1011122202303010-1113102100300303-2010201110333222-1223221013011323-3220023221321200-2123131310033301-0302233330000111) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info](data-sources--proxy--reference--group-002.md#canonical-3211102222212302-0033302310130320-1031323022131132-0121131100222101-1222031021211330-1312201131101231-2211000200302110-1221000122110021) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.provider_ref](data-sources--proxy--reference--group-002.md#canonical-3303201130223220-0203303223302211-3001021132200222-1222203131300213-1222111000010232-1103003211020300-2132223000223002-2121332230303121) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.url](data-sources--proxy--reference--group-002.md#canonical-0331003312223320-1111131030123311-0200321322003221-1321210120132220-2131022001331130-3003000113011312-2211001203111333-3223133323310012) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.value` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.value](data-sources--proxy--reference--group-002.md#canonical-1230232323110101-1102031022210022-1021312321323213-0300001323213303-2112213330203121-2302321323030303-0220131203011102-1121012011203323) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_remove` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_remove](data-sources--proxy--reference--group-001.md#canonical-1333123000020021-2122322132331113-3302010131303300-0320032313303003-2012201131301132-1200200203003010-1103222201312003-0133312102013233) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add](data-sources--proxy--reference--group-002.md#canonical-3312212330130212-1201120222102111-0201011303032130-2220131322201202-1311213021322313-1302101033302113-3111130332210110-3103131033031213) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.append` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.append](data-sources--proxy--reference--group-002.md#canonical-1323013212323222-0133321011100133-2333213202320102-1223333010032010-2223003322201213-3032033112221002-2303013121220003-2300013210310010) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.name` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.name](data-sources--proxy--reference--group-002.md#canonical-1012103112123222-1301212121322131-2311113300033002-1002302330012130-3233130120211102-2020022212321211-1302220202233010-3010312200232002) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-0202312230102310-0202111122000122-2000221003011320-3123021103110030-2213030221113132-1013111103002002-1320332303203212-3110312302301201) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info](data-sources--proxy--reference--group-002.md#canonical-1121220313233030-1120102232203323-1332201110330110-2032322310121320-1230120122031222-2301011113313200-0212100131332320-3322222011332332) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--proxy--reference--group-002.md#canonical-2203122130223323-3003210121230023-1021220133033131-3320003312111331-3330030010033020-0113112013231021-3232011211001010-3133000233233300) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.location](data-sources--proxy--reference--group-002.md#canonical-0101123010030322-0033102211302131-2000200321033131-1322131201111022-3032102003220133-3011213101231032-2002232201130011-3212312321220301) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--proxy--reference--group-002.md#canonical-3133303311122322-1310223103032302-2013010233303330-1312322111100110-3223200101223111-1020300112103021-0210103310023021-0121221010030221) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info](data-sources--proxy--reference--group-002.md#canonical-2321312023321100-1110303211202212-1330033333322010-2022310312130231-0322110112100033-1300203012310301-2131310320122331-3230100003121011) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.provider_ref](data-sources--proxy--reference--group-002.md#canonical-3102201322001013-3203201013031202-3003333023302012-3131310322031333-2103002121302113-3322232313022210-0300102131110032-1133032300032321) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.url](data-sources--proxy--reference--group-002.md#canonical-1202110312310200-1031310233010010-3110311022210110-2133120010113003-0132021131000010-2320003022322121-1013121223031113-0120221031331221) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.value` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.value](data-sources--proxy--reference--group-002.md#canonical-1230313323202122-2020220200321131-3322123220303122-1232202302011313-2301030330311020-3233312120032133-0111213333132301-1022121123202203) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_remove` | [dynamic_proxy.http_proxy.more_option.response_headers_to_remove](data-sources--proxy--reference--group-001.md#canonical-1312001010100030-2312110320210232-3023322332332132-1220131202220111-3230330222213212-0011031032133123-3320310331220020-2221232102023022) |
| `dynamic_proxy.https_proxy` | [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-1222311003032122-1302001312111013-2110112322002123-3320333003001322-0310220221331213-3032032120010010-1010200111113120-0300001233302320) |
| `dynamic_proxy.https_proxy.more_option` | [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-0330330032102302-0232000313232301-3130033233130020-3331220030003003-0300111231313010-0220033130201110-3321033000213100-3331323221103331) |
| `dynamic_proxy.https_proxy.more_option.buffer_policy` | [dynamic_proxy.https_proxy.more_option.buffer_policy](data-sources--proxy--reference--group-002.md#canonical-1221233103020020-0003012303203331-0212230222222101-2003133212022001-0121012210022112-0103213300230112-1011033033220303-1003103313210123) |
| `dynamic_proxy.https_proxy.more_option.buffer_policy.disabled` | [dynamic_proxy.https_proxy.more_option.buffer_policy.disabled](data-sources--proxy--reference--group-002.md#canonical-1203121332032200-1223221101111131-1310112222102202-3133233200132133-2313110030313103-3011123213332230-2130123013213113-0010002233313120) |
| `dynamic_proxy.https_proxy.more_option.buffer_policy.max_request_bytes` | [dynamic_proxy.https_proxy.more_option.buffer_policy.max_request_bytes](data-sources--proxy--reference--group-002.md#canonical-1021131103002220-0300022023301210-3200010011320302-3120120332300130-0211122032211102-0113003022022230-1001322131012303-2131230101131102) |
| `dynamic_proxy.https_proxy.more_option.compression_params` | [dynamic_proxy.https_proxy.more_option.compression_params](data-sources--proxy--reference--group-002.md#canonical-3131232203221303-1301131211000030-0000300231122333-3320212213012330-2103033113211120-1023230022101131-3010100320233032-1203112331212132) |
| `dynamic_proxy.https_proxy.more_option.compression_params.content_length` | [dynamic_proxy.https_proxy.more_option.compression_params.content_length](data-sources--proxy--reference--group-002.md#canonical-1311303312130301-1023002133002212-3010220210302301-2103223022001132-0022032231332033-3112010132021102-0031312001202101-1232331232003220) |
| `dynamic_proxy.https_proxy.more_option.compression_params.content_type` | [dynamic_proxy.https_proxy.more_option.compression_params.content_type](data-sources--proxy--reference--group-002.md#canonical-0202323112211111-1310033312203211-1000031212200310-2313003310113303-0022301010333100-3013000101013112-0123330211231033-2331123222300032) |
| `dynamic_proxy.https_proxy.more_option.compression_params.disable_on_etag_header` | [dynamic_proxy.https_proxy.more_option.compression_params.disable_on_etag_header](data-sources--proxy--reference--group-002.md#canonical-0203013101331223-3022223011232200-0101313112222101-1011230231132233-0312201121110103-3322130232202100-1301001333012223-1310233011221310) |
| `dynamic_proxy.https_proxy.more_option.compression_params.remove_accept_encoding_header` | [dynamic_proxy.https_proxy.more_option.compression_params.remove_accept_encoding_header](data-sources--proxy--reference--group-002.md#canonical-1010221101013213-2032320220222230-1102100011002203-3311200031013012-2230101231012133-1033130330112113-1202231112020231-1000111302331112) |
| `dynamic_proxy.https_proxy.more_option.custom_errors` | [dynamic_proxy.https_proxy.more_option.custom_errors](data-sources--proxy--reference--group-002.md#canonical-1302320102022003-1333333000110020-1012331230132103-1222010333210330-2202321101220122-3232101221031000-0102132312131302-2111301003330301) |
| `dynamic_proxy.https_proxy.more_option.disable_default_error_pages` | [dynamic_proxy.https_proxy.more_option.disable_default_error_pages](data-sources--proxy--reference--group-002.md#canonical-0223222221302010-0121102012211230-0121112100120322-0123230212103030-0002101113013302-0030330321232012-3001300102033103-1133102033002212) |
| `dynamic_proxy.https_proxy.more_option.disable_path_normalize` | [dynamic_proxy.https_proxy.more_option.disable_path_normalize](data-sources--proxy--reference--group-002.md#canonical-3103332100222101-1112002010100030-1022130110103320-0300221131200130-0222300223101232-0223103123020310-2123021211020331-3213302213300322) |
| `dynamic_proxy.https_proxy.more_option.enable_path_normalize` | [dynamic_proxy.https_proxy.more_option.enable_path_normalize](data-sources--proxy--reference--group-002.md#canonical-2303121322131302-0230000310301101-2320103130300030-0030232112012122-1302210012013113-2012231111301211-0121003020113310-3320323211012003) |
| `dynamic_proxy.https_proxy.more_option.idle_timeout` | [dynamic_proxy.https_proxy.more_option.idle_timeout](data-sources--proxy--reference--group-002.md#canonical-0312330002201332-3100003203032203-0101221121320130-1221121131122020-2012003110103213-3312121132100001-0321022011323020-0223310333331300) |
| `dynamic_proxy.https_proxy.more_option.max_request_header_size` | [dynamic_proxy.https_proxy.more_option.max_request_header_size](data-sources--proxy--reference--group-002.md#canonical-0002213032113133-1200220100003330-0231031223001001-0330113303132221-0202112113332311-1112120120223211-1213301131311022-2111121131311332) |
| `dynamic_proxy.https_proxy.more_option.max_requests_per_connection` | [dynamic_proxy.https_proxy.more_option.max_requests_per_connection](data-sources--proxy--reference--group-002.md#canonical-2332110313100232-1131222210031023-0303012112310203-0323102202300320-3303200030222001-1001023231031120-3011000223203020-1032200300223022) |
| `dynamic_proxy.https_proxy.more_option.no_request_limit_per_connection` | [dynamic_proxy.https_proxy.more_option.no_request_limit_per_connection](data-sources--proxy--reference--group-002.md#canonical-0311321200323311-3132122110012022-3221232200013330-3110022100232103-3012002031333200-1312122310023310-3112313221121133-2333323311033011) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-1120003130332102-3121303130133111-0023133331300321-0002012231203222-1101000221120303-2323012301233003-1123031022032000-1310010103221012) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.name` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.name](data-sources--proxy--reference--group-002.md#canonical-1100023031122013-1331003032000011-2011322302303301-3203312230121030-2233200223203200-0300320203032202-1222001122113000-2020031023003013) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.overwrite` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.overwrite](data-sources--proxy--reference--group-002.md#canonical-1130010121100103-3011233012203111-2110212110322131-0010312120010233-2321110312001212-3333001021001011-1020331130031320-1300002110101210) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-0133113203220032-3033213123102000-1013000023302223-3103331001200322-0100311102203213-3321203303331111-2132103233303122-1210221212212101) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info](data-sources--proxy--reference--group-002.md#canonical-1231001223112233-1321000010203011-1032331011123031-2031100232232332-1110311110000232-1313233231000023-1030103103223213-1032030231000311) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--proxy--reference--group-002.md#canonical-3103002003010201-3021132302230011-3300010130223332-2101111230202210-0111103311123100-0111311311020023-2103010130023331-1220221010320302) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location](data-sources--proxy--reference--group-002.md#canonical-3132202111311010-3311010320310033-1122322303312122-1203012001322312-0210313310102133-3130203121113012-0013131301013223-2022311320301310) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--proxy--reference--group-002.md#canonical-3001003132301223-1232013310232003-0133100211221012-1131300331221312-1110222122002303-2011332001312321-3233001331213330-0323313322212331) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info](data-sources--proxy--reference--group-002.md#canonical-0113020303211220-1332132312111122-1301012222223012-2211200003001330-0122331332012230-1223322331220311-1123011031012021-3200331220121021) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref](data-sources--proxy--reference--group-002.md#canonical-3111300220203102-3031202013232321-2312111032100332-3001100302231212-1101020300111320-0302322211132302-1203010011220012-3301030131300131) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.url](data-sources--proxy--reference--group-002.md#canonical-0021121303003233-3200122100110013-1133113020133311-0001003012301202-1102113310312132-2313320320203010-0233100132023220-2331203032301301) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.value` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.value](data-sources--proxy--reference--group-002.md#canonical-3021100312103311-2102033330200311-0013322122301123-0122110333211210-3323112320311102-3310332223011132-1000001312233221-0231020102213213) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_remove` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_remove](data-sources--proxy--reference--group-002.md#canonical-0013021232203110-0121011220101202-0003133232002101-1303313031223013-0032001011132200-2033312112211033-2131120213333130-1011020010110310) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add](data-sources--proxy--reference--group-002.md#canonical-3312031003003121-2202113112200322-0311030122313202-3130310302231112-0031002230332131-1120122303100123-1331331323031201-2212133220310223) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.append` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.append](data-sources--proxy--reference--group-002.md#canonical-2212000203223010-2011131120233121-3202112123223212-1202111112133130-2231013131032122-0310131023332023-3010222311220312-2102212020231131) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.name` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.name](data-sources--proxy--reference--group-002.md#canonical-2210312100202211-2013002111300220-0020103200032232-2200022311110202-1210002020220311-1012010313120200-2002231012021120-3013201032032113) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-1201101302100032-2312131313221330-0211230131000012-2121132320121313-2133303012002113-0222332313332133-1212332302103120-1002132021133021) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info](data-sources--proxy--reference--group-002.md#canonical-1002102230330012-3122222012010311-0222330200330320-3210312211000010-0010202023332122-3330033020220023-0010301020033212-2303031023232022) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--proxy--reference--group-002.md#canonical-0322132111223111-2202133330222211-1221230231001330-3231222132300210-3010112231231102-2123131010101110-2202002103210011-3222132012300221) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.location](data-sources--proxy--reference--group-002.md#canonical-2320313302110102-2210321030231001-2101100211020111-2230323322223010-0310213100130211-1000210020132333-0322013331220231-3120313220311312) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--proxy--reference--group-002.md#canonical-1123323133232202-3211333100130330-0020002032203111-1322101201000131-1220020331212222-3032331030312213-0321111332111123-2331200003212000) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info](data-sources--proxy--reference--group-002.md#canonical-1023202223030322-0313112312020121-2212013320301123-2000311231112222-0221023111221332-1020202100200000-3021231222311020-3111212131220202) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref](data-sources--proxy--reference--group-002.md#canonical-3123032203120132-1120012321001000-0003211232211102-3312131213333323-3032311131330303-2023032022123231-0003031300032202-1123200033130021) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.url](data-sources--proxy--reference--group-002.md#canonical-2122312013300023-1221200303212211-1102101332313232-1332332111322121-2203113320303313-2012103022033212-1020333120132211-1133323213322121) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.value` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.value](data-sources--proxy--reference--group-002.md#canonical-3101202010010102-2013122031202021-0320233111012113-2202211220032300-3021322310300301-0213202323222031-1011311031121332-1321233313221013) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_remove` | [dynamic_proxy.https_proxy.more_option.request_headers_to_remove](data-sources--proxy--reference--group-002.md#canonical-1212333223303331-3022333132132112-1203003331023323-3032033020311210-2331301132021322-3211312001213031-3303322313130032-0303311302102031) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-0203120121330102-2310212113121002-1120030322022220-0021020203203213-0212332201212022-2331232103202013-0313320031023130-1313210203310003) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_domain` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_domain](data-sources--proxy--reference--group-002.md#canonical-2022112002030110-1222131032112301-1233132032302020-0120322100003302-2312102032222000-1332312211010230-3323001133230011-2232010333322202) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_expiry` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_expiry](data-sources--proxy--reference--group-002.md#canonical-1212032033300313-1301132213333302-3030222031003203-1322202300112311-2022221101100221-1110322100031020-3322033223130021-0122232100120202) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_httponly` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_httponly](data-sources--proxy--reference--group-002.md#canonical-1030022032310100-2200333123033030-1020121311301202-3232002113321223-3210202201221232-1012003202213320-3233013010301320-1210031031113033) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_partitioned` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_partitioned](data-sources--proxy--reference--group-002.md#canonical-2011313122003311-0223031331331003-2301310011200011-3112222330233212-1131120132013220-3123003001333333-3132230233110333-0311330131203110) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_path` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_path](data-sources--proxy--reference--group-002.md#canonical-1210103000201213-3220112232320100-3330002123110310-2300011203210321-2210123033311210-3303213100123121-3000000213330323-1112320221323100) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_secure` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_secure](data-sources--proxy--reference--group-002.md#canonical-0321121222223030-1311333233112223-2103133110020100-0201112112220123-2223311032212113-2120131010310131-1012112301223313-3212002002312110) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_domain` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_domain](data-sources--proxy--reference--group-002.md#canonical-1231030112203133-1010332211102021-3210120233303311-3320121112233023-2123300013011131-2002103011020201-1232121131310112-1320221000011130) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_expiry` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_expiry](data-sources--proxy--reference--group-002.md#canonical-0233113013210303-2013112211312213-3113312111130111-2202032233310023-3213131133321033-2113332020222301-1331220220232222-0033110133120232) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_httponly` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_httponly](data-sources--proxy--reference--group-002.md#canonical-0230123113032310-2322121012122312-3031113000122300-2221222231011322-1311103002101030-1013101312023002-1230013310330002-1331021203130223) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_max_age` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_max_age](data-sources--proxy--reference--group-002.md#canonical-0131313331203112-2100101212131003-2020203023333111-0310100020310331-0310210320213022-2313111203333013-3003102002231012-0231303203003233) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_partitioned` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_partitioned](data-sources--proxy--reference--group-002.md#canonical-0231110023133002-3110330031210010-0320233231333020-0302020122033101-0330232221310331-0103321113031111-1131131223313020-2310010123333320) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_path` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_path](data-sources--proxy--reference--group-002.md#canonical-2030032232231213-3130313302001010-2300312000022300-1132120112323211-0231113112103320-0302023001012301-0322012212132303-3013222333100010) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_samesite` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_samesite](data-sources--proxy--reference--group-002.md#canonical-2203303222021322-0201021023033333-2123103311320302-1331302322030200-3301022300222132-0103101123022023-1001321330113020-3321332123013221) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_secure` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_secure](data-sources--proxy--reference--group-002.md#canonical-2200022321221103-2123330123332320-2311233131232032-2023203213210230-2101311320111103-2332003031300032-2100012103012022-2203222011323222) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_value` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_value](data-sources--proxy--reference--group-002.md#canonical-3020130010203313-2210332222101212-1322222213313300-2101011023020002-2202022313232223-2223301103300232-2232102331233001-1020331111212232) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.max_age_value` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.max_age_value](data-sources--proxy--reference--group-002.md#canonical-1212020221200133-2323012123101123-3213132102310110-2233230111231022-3102103021301310-2002221321223101-2021020021202002-1110002211310103) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.name` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.name](data-sources--proxy--reference--group-002.md#canonical-0110202313010001-1121133213333113-0101222132333212-3200022023312210-2310131213301221-1331200000123030-0323333323032232-3223210130201322) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.overwrite` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.overwrite](data-sources--proxy--reference--group-002.md#canonical-3000122130200330-3231321103321023-3331223223203131-1120233000131120-3333013101301112-0210113230232133-3200330101230123-0323223310210211) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_lax` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_lax](data-sources--proxy--reference--group-002.md#canonical-2222110232022130-1222222010113003-0200310333010113-1220212012200300-0123132111000303-1001223013233230-1210320030233102-1211002312322133) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_none` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_none](data-sources--proxy--reference--group-002.md#canonical-0312001230223020-1002031033023122-1100130220123213-1310030333202222-1003211123331333-0331112021100212-1002312132332031-3203220321233001) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_strict` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_strict](data-sources--proxy--reference--group-002.md#canonical-1320223211213201-1223210201030313-1102123301201132-2100323030100020-0232022333320012-3200033230230033-0032330000330111-3213232300121022) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-3303231130222320-3223020332320100-0313213221120110-0122233002133301-2332120003203031-2332103202211133-2311010202321121-1202030103302313) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info](data-sources--proxy--reference--group-002.md#canonical-3301333212312201-3212102303333011-1300233102130222-2032333312031302-1323322320032323-1032203333321020-3111021223230103-0133010300100113) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--proxy--reference--group-002.md#canonical-0111120030232022-1323011110111102-3033202332200303-2212102123010221-1211132233012111-0203330322001021-0302132221231323-3030113031020103) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.location](data-sources--proxy--reference--group-002.md#canonical-3020112322000122-1122021322330203-1320311223332311-0331220232120013-2213313102123132-2000202211132200-1120321122023313-0013101230310330) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--proxy--reference--group-003.md#canonical-3232222122232311-1303120010233130-2230232120103321-2023120003321320-1200211100322123-1133121230202310-3211203200201212-2010131310302021) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info](data-sources--proxy--reference--group-003.md#canonical-2200302302110211-1030300101310031-3322320121003133-0213223013303100-2221221310320301-2322100123021320-2320301032302133-0031200210212013) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.provider_ref](data-sources--proxy--reference--group-003.md#canonical-0033313300123220-0303310302311121-3233300133011331-2303023003323011-2313313003332201-1000333331122312-1303033221011200-0321333230020320) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.url](data-sources--proxy--reference--group-003.md#canonical-2021332301232030-3332130002023033-0213332130212210-3323202321013330-1003020310333300-0203131021231232-1330130101230111-2132011002112003) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.value` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.value](data-sources--proxy--reference--group-002.md#canonical-3032110323023033-2002113132223120-1200213300220020-0023233022202212-0012023221111023-0130313101113113-3001233302302130-2331033323302301) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_remove` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_remove](data-sources--proxy--reference--group-002.md#canonical-0313202333230232-0103223103102030-2202010300303120-1313132013333101-2332322122312001-1122100232002210-1311122231003100-3313300213211212) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add](data-sources--proxy--reference--group-003.md#canonical-2123101123332311-3031202301323130-0110300312210222-2331121320202100-1032103031232223-1221221122131211-3001213003030130-0201323233223000) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.append` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.append](data-sources--proxy--reference--group-003.md#canonical-1203032111112212-1020332110123231-3130231133112021-0131313031133202-1321301012112002-3212120301020021-3231122210222203-1212212233102003) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.name` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.name](data-sources--proxy--reference--group-003.md#canonical-3011221200301132-0212202310111112-0303101213203330-2211100330222323-1201320103233113-0310121302112332-1220031203302303-2032112113121230) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value](data-sources--proxy--reference--group-003.md#canonical-0331200100122020-0301121230130102-2003231012100132-0112200302302112-0323233221121232-2003030300322010-3302211031220210-0311310312003221) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info](data-sources--proxy--reference--group-003.md#canonical-3033000212132313-1011020111322012-0131230311002101-0221211232220222-0200312013302312-0302011011221003-3111030312100301-3121130012003023) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--proxy--reference--group-003.md#canonical-1333311221313012-0030120122211311-1133310333000323-1330123010022211-1113302230030213-1130310111220103-3212022001202221-0002211203133121) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.location](data-sources--proxy--reference--group-003.md#canonical-3310030111011022-3321011130112133-3113132022130011-3233200231011133-1220312210201110-0032212231001233-2112111300221030-3110021021200320) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--proxy--reference--group-003.md#canonical-0222300231300331-0132020303031330-3131022120033021-2013110032131232-1332220232013012-1211113200120330-0010022120030132-0001313023101200) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info](data-sources--proxy--reference--group-003.md#canonical-0012100123010332-3102003203112311-2012322311223210-0021133121103213-1221210203230022-1023023030322011-0121323120312313-1300132303101320) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.provider_ref](data-sources--proxy--reference--group-003.md#canonical-1333200201101002-1033310131011013-0321333022133222-0010322133221113-1202013000131213-1333203313123211-0333322113023222-1221302001133332) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.url](data-sources--proxy--reference--group-003.md#canonical-2032211333012200-0220231303113202-2311000022011121-1102123302100200-1000332131033000-1231001213202310-0332003312123230-2013230032302311) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.value` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.value](data-sources--proxy--reference--group-003.md#canonical-1012013020301023-2200130313023231-0212323002033231-3023001233120132-1120300321332100-3011223232231003-3001301013222003-1022111320101103) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_remove` | [dynamic_proxy.https_proxy.more_option.response_headers_to_remove](data-sources--proxy--reference--group-002.md#canonical-2232102300210212-3233132000220001-2013323301012302-0321031023311030-2000212133332333-2103031122121122-1330211330231213-3230132013332311) |
| `dynamic_proxy.https_proxy.tls_params` | [dynamic_proxy.https_proxy.tls_params](data-sources--proxy--reference--group-003.md#canonical-2300121300020131-3132302333231323-3132222033011130-0022032313222002-2010322122031022-0022320002210033-3111311330003131-1101302211212013) |
| `dynamic_proxy.https_proxy.tls_params.no_mtls` | [dynamic_proxy.https_proxy.tls_params.no_mtls](data-sources--proxy--reference--group-003.md#canonical-1002223300332130-1230101312003133-3303313330101023-2021123122231123-3101112032202332-3300220121122133-2231212301330201-2333313211300021) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates` | [dynamic_proxy.https_proxy.tls_params.tls_certificates](data-sources--proxy--reference--group-003.md#canonical-1002312003201232-2023232030333123-2231313101130321-2323101203222030-2113231220210200-3033303331023323-1112300223332130-0231230210233312) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.certificate_url` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.certificate_url](data-sources--proxy--reference--group-003.md#canonical-1300132222232231-1223212123003021-0111212022130003-3323320311212010-3300001320121301-3100232323012031-1010023213011301-1113232112020030) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms](data-sources--proxy--reference--group-003.md#canonical-3133222303131002-2122230320100212-1012211110031022-3310322230231332-2203323012000233-3113332121322211-0201233033003031-0111031300323231) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms.hash_algorithms` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms.hash_algorithms](data-sources--proxy--reference--group-003.md#canonical-2023110202032210-1030213310201112-3122032123133210-2011232302131231-1023103201122233-1311100200330100-3131203213111303-2310332331120323) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.description_spec` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.description_spec](data-sources--proxy--reference--group-003.md#canonical-2230333312130010-3022222130132212-0231222010012012-0021310101330113-2121333113203120-0132322321331000-1220303031313233-3221122132103232) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.disable_ocsp_stapling` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.disable_ocsp_stapling](data-sources--proxy--reference--group-003.md#canonical-3112022122101323-0202021220102300-2301030102212030-1322202212310302-0232120133103111-0223301011113111-1013322303020021-0122230123032013) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key](data-sources--proxy--reference--group-003.md#canonical-2201122231301012-0023331312313012-3111012303001312-3313003131030003-2202110322321101-1013333300130223-1002122210212113-3323121010100111) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info](data-sources--proxy--reference--group-003.md#canonical-0220201112112132-0230303230010213-2201101010300111-0110301120022122-1013012213023133-2132122232312123-1032211112121132-3022210313110030) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider](data-sources--proxy--reference--group-003.md#canonical-0113332310103101-2201013210101201-0322100201210010-0130031001220132-1101031000312111-0101223011302321-0100312211232230-1120120103220211) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info.location` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info.location](data-sources--proxy--reference--group-003.md#canonical-3311220132230002-2320322001203333-1233002013210001-2132331211212302-2000010110112303-0021123101023301-1103200010131302-1112001132333203) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info.store_provider` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info.store_provider](data-sources--proxy--reference--group-003.md#canonical-1211132311312320-3000011013212332-2031301121321132-0331222330001310-1321220321232222-0333220001221023-1301002123310200-2311231002100201) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info](data-sources--proxy--reference--group-003.md#canonical-2232203133202201-3333221331213311-3021233210331122-3300223213000300-0221232032101112-1030222132330303-3320132222023033-1210122333011033) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info.provider_ref` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info.provider_ref](data-sources--proxy--reference--group-003.md#canonical-3320300320001002-3130031302021133-1313031321123203-2322022303002203-2332123220023202-0222101023033101-2213230201032202-1112112133200033) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info.url` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info.url](data-sources--proxy--reference--group-003.md#canonical-1020100020033023-1120223332312312-3100031003333000-1023013321312232-1023031311213011-2030022002003200-3132232223321303-0311003013111311) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.use_system_defaults` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.use_system_defaults](data-sources--proxy--reference--group-003.md#canonical-2232201200302112-0320111133320302-1310300220032100-1203112203332311-1200031222102001-3102333300013303-2123120102021220-3231013120212222) |
| `dynamic_proxy.https_proxy.tls_params.tls_config` | [dynamic_proxy.https_proxy.tls_params.tls_config](data-sources--proxy--reference--group-003.md#canonical-2233112202122010-3010112132002021-3210133102201233-0102331302303331-1203210212030220-2213011200032022-2112200013302032-2331033311323313) |
| `dynamic_proxy.https_proxy.tls_params.tls_config.custom_security` | [dynamic_proxy.https_proxy.tls_params.tls_config.custom_security](data-sources--proxy--reference--group-003.md#canonical-0101030312021031-2020030211033233-0322120320230323-2111320223003101-1311201213323310-2221133022000013-1211312111311020-1130111332220323) |
| `dynamic_proxy.https_proxy.tls_params.tls_config.custom_security.cipher_suites` | [dynamic_proxy.https_proxy.tls_params.tls_config.custom_security.cipher_suites](data-sources--proxy--reference--group-003.md#canonical-2003332020112002-3210312220301022-3001002301303232-0131013312213112-1302033021123101-1000322131321013-1123301100302321-0113312212001110) |
| `dynamic_proxy.https_proxy.tls_params.tls_config.custom_security.max_version` | [dynamic_proxy.https_proxy.tls_params.tls_config.custom_security.max_version](data-sources--proxy--reference--group-003.md#canonical-1210233313320332-1121130020313022-2022331102103000-3301030303210330-3222112231001010-0010233033112312-2121020022111110-3013031130330231) |
| `dynamic_proxy.https_proxy.tls_params.tls_config.custom_security.min_version` | [dynamic_proxy.https_proxy.tls_params.tls_config.custom_security.min_version](data-sources--proxy--reference--group-003.md#canonical-1022111030033021-0030322110011203-2121133310121210-0321323120330203-0212333011320102-1201123131033233-3113303223023031-0231300120112233) |
| `dynamic_proxy.https_proxy.tls_params.tls_config.default_security` | [dynamic_proxy.https_proxy.tls_params.tls_config.default_security](data-sources--proxy--reference--group-003.md#canonical-3223030231100320-2012111120031111-2301211201031211-3332323200303003-0230100011031211-2233131202203130-2220231033020220-2013213212202322) |
| `dynamic_proxy.https_proxy.tls_params.tls_config.low_security` | [dynamic_proxy.https_proxy.tls_params.tls_config.low_security](data-sources--proxy--reference--group-003.md#canonical-1220303220311330-1101122010200201-3000233111013203-3201023330031113-0301031121020101-1213213110223333-3322012113303323-3021021202113121) |
| `dynamic_proxy.https_proxy.tls_params.tls_config.medium_security` | [dynamic_proxy.https_proxy.tls_params.tls_config.medium_security](data-sources--proxy--reference--group-003.md#canonical-2131330032020130-3311112022321122-3311023213102130-0232202032303231-0032030210120103-0030021122101131-3122300233323213-1303300121302012) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls` | [dynamic_proxy.https_proxy.tls_params.use_mtls](data-sources--proxy--reference--group-003.md#canonical-2000211023303131-0012230310021300-0022112001322322-2332011321010212-2100130113212330-1020212222020300-0021120300200211-3333001223320132) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.client_certificate_optional` | [dynamic_proxy.https_proxy.tls_params.use_mtls.client_certificate_optional](data-sources--proxy--reference--group-003.md#canonical-1103113010020312-3300002120333031-2211001101000133-0312213313003010-2011223103131322-0111321133213011-3021310013032100-0321210202313300) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.crl` | [dynamic_proxy.https_proxy.tls_params.use_mtls.crl](data-sources--proxy--reference--group-003.md#canonical-1303123232010222-2232220211003013-2023320332332330-3123120313001031-3332300223103232-2311121201101030-1020002203002022-2103022013322123) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.crl.name` | [dynamic_proxy.https_proxy.tls_params.use_mtls.crl.name](data-sources--proxy--reference--group-003.md#canonical-0133030130112123-1212102000133321-0311213121323301-2323033000122222-3211023210313102-0203200033231223-1231313331030133-0210213233223112) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.crl.namespace` | [dynamic_proxy.https_proxy.tls_params.use_mtls.crl.namespace](data-sources--proxy--reference--group-003.md#canonical-3223210211221312-1120002302130021-2320323020311011-0132201003303232-2313133110133002-2001030003300213-2213202113230030-0330122322000131) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.crl.tenant` | [dynamic_proxy.https_proxy.tls_params.use_mtls.crl.tenant](data-sources--proxy--reference--group-003.md#canonical-0000323110331110-3232231213201033-3223100323220300-2313023310031130-0210100303322330-3100133100233321-0212122000312301-3123020001112103) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.no_crl` | [dynamic_proxy.https_proxy.tls_params.use_mtls.no_crl](data-sources--proxy--reference--group-003.md#canonical-0131113120331002-0010031001001200-0110003003011223-2120211010221030-1103213121000003-1333012311113002-1220312222011012-1300101230203200) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca` | [dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca](data-sources--proxy--reference--group-003.md#canonical-1200001031022130-1320001220323302-0323102312131312-2330301113213313-0320302332203300-0223300300320000-1200110221322102-2031032201210113) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca.name` | [dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca.name](data-sources--proxy--reference--group-003.md#canonical-1103202323220033-2101211321111112-2121113333300023-2110222103032012-3203100213120033-3011113322000003-3020302310021231-0021210332111221) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca.namespace` | [dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca.namespace](data-sources--proxy--reference--group-003.md#canonical-1303332203210303-2023211021000011-2331111123020320-0030312303133223-3221030101122311-1003211021323331-2003121100232102-3233200110013313) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca.tenant` | [dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca.tenant](data-sources--proxy--reference--group-003.md#canonical-2013220022101112-0002121100302233-0112301331332211-2022211231011220-3303123211320132-2232012112110332-1302111223120003-2112111123321012) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca_url` | [dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca_url](data-sources--proxy--reference--group-003.md#canonical-3210023202232112-0212033313200000-3021111130211213-3022111022021333-0321132231022130-2330322320200102-0101321223003003-0013330323121332) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_disabled` | [dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_disabled](data-sources--proxy--reference--group-003.md#canonical-1202022302133120-0300000122233013-3220023003210312-1111200222033232-3121000202130220-0122123310002213-1001121220103201-2022203130321011) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options` | [dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options](data-sources--proxy--reference--group-003.md#canonical-2310130021133233-1332332030232213-1300031130301012-1331213332102230-1033302320022221-0031213302121323-1321231031322020-1112233103330111) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options.xfcc_header_elements` | [dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options.xfcc_header_elements](data-sources--proxy--reference--group-003.md#canonical-2203001132031022-2000203311031220-1322123101332333-3313300201210223-3133303120300111-2022100200122011-3330333231123333-3202301102212033) |
| `dynamic_proxy.sni_proxy` | [dynamic_proxy.sni_proxy](data-sources--proxy--reference--group-003.md#canonical-3032311002003003-0031320000113303-0231310022213220-1000120202113332-0322013233223033-3222133013023011-3303303230201203-2301013113121233) |
| `dynamic_proxy.sni_proxy.idle_timeout` | [dynamic_proxy.sni_proxy.idle_timeout](data-sources--proxy--reference--group-003.md#canonical-1011010201103311-1212012221033122-0130321032331332-3001112123001233-3300202030322131-1303013320000301-2312033331002021-3223312233001103) |
| `http_proxy` | [http_proxy](data-sources--proxy--reference--group-003.md#canonical-2121010022301312-1300233121020222-1323121211031111-0131133120001130-0000032031101232-1000113012222203-2213302130100231-2210233233221031) |
| `http_proxy.enable_http` | [http_proxy.enable_http](data-sources--proxy--reference--group-003.md#canonical-0230111231103311-1312220320210302-2330212201112022-3310011330222110-3132131100012230-2000310112111311-0030002102130202-2303211213010213) |
| `http_proxy.more_option` | [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-1201330330310212-2100230231110200-3132112310002331-0010303310030303-1330321022302012-2110220102020001-0331300200111120-0131133122200222) |
| `http_proxy.more_option.buffer_policy` | [http_proxy.more_option.buffer_policy](data-sources--proxy--reference--group-003.md#canonical-1321132332021200-2120310123310102-0103322000100121-3000030121130002-2223320110121200-3112131000011300-2302222120113103-0113223023320013) |
| `http_proxy.more_option.buffer_policy.disabled` | [http_proxy.more_option.buffer_policy.disabled](data-sources--proxy--reference--group-003.md#canonical-0313221210102222-3210003110231013-3300331310032013-3133022030230113-3130303231202203-0230032231133111-3002212000013123-0212133113301220) |
| `http_proxy.more_option.buffer_policy.max_request_bytes` | [http_proxy.more_option.buffer_policy.max_request_bytes](data-sources--proxy--reference--group-003.md#canonical-1022132332103130-2100113231102003-2233023212231200-1330201321202132-2301302100131211-0101312020230220-1230012310100323-1313000000313122) |
| `http_proxy.more_option.compression_params` | [http_proxy.more_option.compression_params](data-sources--proxy--reference--group-004.md#canonical-0313033112003111-2131131313202310-0310031032122100-2303211320132223-3331103232121323-0320331222212100-2023100122002211-0131301333121303) |
| `http_proxy.more_option.compression_params.content_length` | [http_proxy.more_option.compression_params.content_length](data-sources--proxy--reference--group-004.md#canonical-1322130002332112-0230232123110323-0232012131331333-0100113123000233-3223220121133023-3221332000303330-0222120232200220-0010211030001202) |
| `http_proxy.more_option.compression_params.content_type` | [http_proxy.more_option.compression_params.content_type](data-sources--proxy--reference--group-004.md#canonical-1031020310231023-2322010012221120-0122303310130022-1030010331301120-1133032310101000-1023311202121200-0111110210231122-1010130120110311) |
| `http_proxy.more_option.compression_params.disable_on_etag_header` | [http_proxy.more_option.compression_params.disable_on_etag_header](data-sources--proxy--reference--group-004.md#canonical-3032003032003313-1021313033000233-3022022031131121-1130121201010010-2120001210120200-1103332211201001-3122211111330320-3021010222233131) |
| `http_proxy.more_option.compression_params.remove_accept_encoding_header` | [http_proxy.more_option.compression_params.remove_accept_encoding_header](data-sources--proxy--reference--group-004.md#canonical-2232202130012220-3333231323331001-2021221313101122-3323331312111021-0303031013333310-1011320002012113-1122202301333201-2202113313321213) |
| `http_proxy.more_option.custom_errors` | [http_proxy.more_option.custom_errors](data-sources--proxy--reference--group-003.md#canonical-2333002231021331-1223330222233313-0101223222300120-2201313302230100-2033113301212220-0111120302312132-1111211110003010-2001023212110120) |
| `http_proxy.more_option.disable_default_error_pages` | [http_proxy.more_option.disable_default_error_pages](data-sources--proxy--reference--group-003.md#canonical-0322331033031011-3133000320003201-3122310323213210-0230212123310233-1022330210020122-1132200321001012-2301221000221230-0222232003013103) |
| `http_proxy.more_option.disable_path_normalize` | [http_proxy.more_option.disable_path_normalize](data-sources--proxy--reference--group-004.md#canonical-1212020001203000-1002120131000330-1130002202233320-0220111300223101-3020222232023222-0120311112213310-2332113122313223-3112102232123113) |
| `http_proxy.more_option.enable_path_normalize` | [http_proxy.more_option.enable_path_normalize](data-sources--proxy--reference--group-004.md#canonical-3020131031130323-1003333333013013-0031312111203230-0310332113001133-2202023111001233-2121033023033110-3300311321320120-1112233200102232) |
| `http_proxy.more_option.idle_timeout` | [http_proxy.more_option.idle_timeout](data-sources--proxy--reference--group-003.md#canonical-3230100301210323-2331103232032032-2312021301210102-2303231322000030-3013231113010120-0333120203020233-0310002123112021-2132021030330003) |
| `http_proxy.more_option.max_request_header_size` | [http_proxy.more_option.max_request_header_size](data-sources--proxy--reference--group-003.md#canonical-1212311023133021-0210021031311231-3300310003332302-2210023032233010-0333113122010121-3223223022301331-3113223303333233-2122131203321233) |
| `http_proxy.more_option.max_requests_per_connection` | [http_proxy.more_option.max_requests_per_connection](data-sources--proxy--reference--group-003.md#canonical-0031233121333221-3320313202222003-1232130012101101-2210201002133311-1303010102031032-3201233201123011-3010003302100032-1231030102331231) |
| `http_proxy.more_option.no_request_limit_per_connection` | [http_proxy.more_option.no_request_limit_per_connection](data-sources--proxy--reference--group-004.md#canonical-3110332221023003-0010012021211113-2023301200130303-3222120211012021-0331212112013313-3113232233010100-2210230223231130-0210331222132133) |
| `http_proxy.more_option.request_cookies_to_add` | [http_proxy.more_option.request_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-1113011033320100-3100121032123121-3233031121013102-1031130132232213-3202123333300222-2301201301130112-2323212312033203-3233220231201101) |
| `http_proxy.more_option.request_cookies_to_add.name` | [http_proxy.more_option.request_cookies_to_add.name](data-sources--proxy--reference--group-004.md#canonical-3033000322330000-1213232200133320-1211122301003003-2012113300231112-0010211301200130-2020310332110233-0000222021201331-1021102222002211) |
| `http_proxy.more_option.request_cookies_to_add.overwrite` | [http_proxy.more_option.request_cookies_to_add.overwrite](data-sources--proxy--reference--group-004.md#canonical-1010232220302003-3220202200031323-2333002213123010-1111103100323102-0023211032023033-3032300020132311-1321032231030323-0311103111213130) |
| `http_proxy.more_option.request_cookies_to_add.secret_value` | [http_proxy.more_option.request_cookies_to_add.secret_value](data-sources--proxy--reference--group-004.md#canonical-0203033032303303-0123120312203130-2013232002310203-3222023000210323-1131022313020101-0121200002133200-0231130033132022-0330102113212001) |
| `http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info` | [http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info](data-sources--proxy--reference--group-004.md#canonical-0200033012003323-0013322130300100-2011011013101202-2030222302131113-1322231221003333-2331001230300201-0300301033133003-0210032112021113) |
| `http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--proxy--reference--group-004.md#canonical-0130001331303131-3333310333203120-1322212022213130-0303102001333301-1031330002211220-0111003230132022-0301301111000313-2022102201300002) |
| `http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location` | [http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location](data-sources--proxy--reference--group-004.md#canonical-1011332010303003-1213312121211321-2213202322201313-1303123330021003-3212003320002323-2323232112221330-2100213033212203-2003231333001123) |
| `http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--proxy--reference--group-004.md#canonical-1022310330123320-3103020123313330-0022110212010001-3013022123332200-1202230030102300-1200112313001113-3303022131122333-0233003220000301) |
| `http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info` | [http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info](data-sources--proxy--reference--group-004.md#canonical-0000210223011120-1130233221012022-0103123200122123-0120213303032301-1023223003121300-1112333313013233-3021322210032112-2301231022302133) |
| `http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref](data-sources--proxy--reference--group-004.md#canonical-0001030300322200-0313002133010203-2231111202212031-3232101310222221-2333220111102221-3310310123131130-0133012201133201-1001220332131100) |
| `http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.url` | [http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.url](data-sources--proxy--reference--group-004.md#canonical-3030303211211211-3310101212230233-1323330200201221-0110023032201002-2110303130301030-3121300313010021-1302203002233002-1231000133023210) |
| `http_proxy.more_option.request_cookies_to_add.value` | [http_proxy.more_option.request_cookies_to_add.value](data-sources--proxy--reference--group-004.md#canonical-3330013233111301-1202200232133333-3200333322223231-2103012312102013-0033022330200313-1303221313023312-1311330233111200-3022222221010100) |
| `http_proxy.more_option.request_cookies_to_remove` | [http_proxy.more_option.request_cookies_to_remove](data-sources--proxy--reference--group-003.md#canonical-1333112201131100-0013301330011210-1210310230113332-1132000232012332-2320103102023022-2133201100202132-1013213203301113-1301200322132320) |
| `http_proxy.more_option.request_headers_to_add` | [http_proxy.more_option.request_headers_to_add](data-sources--proxy--reference--group-004.md#canonical-0012103111211323-0333112200113131-1032002301331130-2023212131232012-1232300131123303-1311330302210322-0321330323130231-3131131131011032) |
| `http_proxy.more_option.request_headers_to_add.append` | [http_proxy.more_option.request_headers_to_add.append](data-sources--proxy--reference--group-004.md#canonical-1323003012311023-0210113211212233-0222203202210333-0112122001121002-2100000100321232-1302112301112200-2322323132321312-3212313202332000) |
| `http_proxy.more_option.request_headers_to_add.name` | [http_proxy.more_option.request_headers_to_add.name](data-sources--proxy--reference--group-004.md#canonical-1131320000020230-0000123322100013-0320000230000302-2230133212233030-1322301012111220-3033132320232201-0132110100021333-0223102033200231) |
| `http_proxy.more_option.request_headers_to_add.secret_value` | [http_proxy.more_option.request_headers_to_add.secret_value](data-sources--proxy--reference--group-004.md#canonical-0012133331210231-1033120113223323-1020300211202120-1313201130120200-2212211100022100-1201223311031002-1100100013211013-3032120110200030) |
| `http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info` | [http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info](data-sources--proxy--reference--group-004.md#canonical-0100003012311202-0013222012310310-3020211233233001-0033110201312022-3120010111013032-1321122101132033-2112321030011213-0222013311131203) |
| `http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--proxy--reference--group-004.md#canonical-0011311022333030-1202323122211012-0320220320201000-0011121233233031-2213112221312031-0212311102201333-3320230013102301-0323221020330010) |
| `http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.location` | [http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.location](data-sources--proxy--reference--group-004.md#canonical-0300011331133010-3202023310003323-0322221003210101-0313121030321020-2033333332223003-1230010332301202-1231320203121001-2220031321221101) |
| `http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--proxy--reference--group-004.md#canonical-3320023322012332-1111321211222023-0003202323023233-2211230333100020-1111021033023222-2002331032302212-2110331002313000-1033330223131223) |
| `http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info` | [http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info](data-sources--proxy--reference--group-004.md#canonical-0201211033211200-3200310322011212-3011133230330231-2230323013132020-2303221112211222-1020302302121231-1003103031301023-2322313032302302) |
| `http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref` | [http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref](data-sources--proxy--reference--group-004.md#canonical-0010212122022320-2132131302032203-2013122030213213-3023132233231111-3112211022120030-2211233031122122-3112212300100202-3001310133111102) |
| `http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.url` | [http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.url](data-sources--proxy--reference--group-004.md#canonical-1233320022020013-3202233331231313-0301313303133223-0130320311003332-3312302023222312-2232023011020130-0311011320030333-0121131223101220) |
| `http_proxy.more_option.request_headers_to_add.value` | [http_proxy.more_option.request_headers_to_add.value](data-sources--proxy--reference--group-004.md#canonical-3213312202203033-3012220023103210-1231122133122220-2010023103100002-3312303002000020-2332023010202030-0213132213022211-3022211123321212) |
| `http_proxy.more_option.request_headers_to_remove` | [http_proxy.more_option.request_headers_to_remove](data-sources--proxy--reference--group-003.md#canonical-3002330133023103-1221321012010332-2102210101113202-1003001322212101-0303003133302003-1213301311010022-2213031032010312-3133101201020200) |
| `http_proxy.more_option.response_cookies_to_add` | [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-1200123032030133-1311012231221230-0103320121233100-3311001301023020-2102110110301231-1013312113122003-3203323300313112-1012323010231313) |
| `http_proxy.more_option.response_cookies_to_add.add_domain` | [http_proxy.more_option.response_cookies_to_add.add_domain](data-sources--proxy--reference--group-004.md#canonical-2113231012001031-1130301311113232-0122110331213320-3232210320202200-2223321022132333-2100313030023203-3333231020002011-3230332320333203) |
| `http_proxy.more_option.response_cookies_to_add.add_expiry` | [http_proxy.more_option.response_cookies_to_add.add_expiry](data-sources--proxy--reference--group-004.md#canonical-0210030111020122-3202120112000020-2010131310221301-1102231220321010-1332102121213102-0210123110211011-3111111122001211-2121032230033322) |
| `http_proxy.more_option.response_cookies_to_add.add_httponly` | [http_proxy.more_option.response_cookies_to_add.add_httponly](data-sources--proxy--reference--group-004.md#canonical-1331002022203331-1033333310320033-2001201311210010-0321032103111211-1233220311111223-1300102223123111-1020030100011002-1201021123003332) |
| `http_proxy.more_option.response_cookies_to_add.add_partitioned` | [http_proxy.more_option.response_cookies_to_add.add_partitioned](data-sources--proxy--reference--group-004.md#canonical-0020020021322301-2323112330110333-3132031320131301-3030112203220102-0031201010023003-3110233333200022-0221330323311020-2120023213311032) |
| `http_proxy.more_option.response_cookies_to_add.add_path` | [http_proxy.more_option.response_cookies_to_add.add_path](data-sources--proxy--reference--group-004.md#canonical-0313033301121032-3003322310301321-2223112230312313-1132203121031112-0233302012132303-3230233300310011-0111130133003333-0320330211301323) |
| `http_proxy.more_option.response_cookies_to_add.add_secure` | [http_proxy.more_option.response_cookies_to_add.add_secure](data-sources--proxy--reference--group-004.md#canonical-3011221200103301-1222210003201331-0311102332120120-0332010031121332-2330031030002322-2102302233131003-1131101110021221-0003032211030222) |
| `http_proxy.more_option.response_cookies_to_add.ignore_domain` | [http_proxy.more_option.response_cookies_to_add.ignore_domain](data-sources--proxy--reference--group-004.md#canonical-1111131023013203-0112033121001130-1301212212221212-1302103223012233-2220332321020130-1233203031130330-2201310022101321-1101322122120332) |
| `http_proxy.more_option.response_cookies_to_add.ignore_expiry` | [http_proxy.more_option.response_cookies_to_add.ignore_expiry](data-sources--proxy--reference--group-004.md#canonical-2112200201133031-0033032031113300-0113111013203012-2111302312221220-2330102030030021-1223330101202320-0331322222200213-2011011320112132) |
| `http_proxy.more_option.response_cookies_to_add.ignore_httponly` | [http_proxy.more_option.response_cookies_to_add.ignore_httponly](data-sources--proxy--reference--group-004.md#canonical-2320132221330003-0011011232132110-3031301320032313-3333311000203302-0000320111201111-1010301020322231-1111232332231212-2302032133120021) |
| `http_proxy.more_option.response_cookies_to_add.ignore_max_age` | [http_proxy.more_option.response_cookies_to_add.ignore_max_age](data-sources--proxy--reference--group-004.md#canonical-1211321020230120-2322013003030213-1322230202122333-2213203210220112-1330333021111023-2202032130213133-3002103000002213-1221302302203103) |
| `http_proxy.more_option.response_cookies_to_add.ignore_partitioned` | [http_proxy.more_option.response_cookies_to_add.ignore_partitioned](data-sources--proxy--reference--group-004.md#canonical-0103010131332030-3110223133113203-3011011322120231-0011122321333322-2300033132222003-2030311020323002-2011233321303232-2203320330121123) |
| `http_proxy.more_option.response_cookies_to_add.ignore_path` | [http_proxy.more_option.response_cookies_to_add.ignore_path](data-sources--proxy--reference--group-004.md#canonical-1320001222320213-0310131312002323-2222203332202312-0013003110011120-3321122321101011-0100100131303023-2331223323102011-2011132232032322) |
| `http_proxy.more_option.response_cookies_to_add.ignore_samesite` | [http_proxy.more_option.response_cookies_to_add.ignore_samesite](data-sources--proxy--reference--group-004.md#canonical-3330101012320310-0132221021103101-3103312213311320-3222131131203223-1102033211003220-1033122103300230-3010010113330131-0030133201101112) |
| `http_proxy.more_option.response_cookies_to_add.ignore_secure` | [http_proxy.more_option.response_cookies_to_add.ignore_secure](data-sources--proxy--reference--group-004.md#canonical-0300001231320332-2132013332301022-3033233003030100-0022222002331003-2012322020211110-0211333113301122-2331010211103100-1210011332222202) |
| `http_proxy.more_option.response_cookies_to_add.ignore_value` | [http_proxy.more_option.response_cookies_to_add.ignore_value](data-sources--proxy--reference--group-004.md#canonical-3021032121332123-1023233222033123-2100321020233211-2102231102321303-1312022012013301-0320031202300021-1320331030223123-0332011332013112) |
| `http_proxy.more_option.response_cookies_to_add.max_age_value` | [http_proxy.more_option.response_cookies_to_add.max_age_value](data-sources--proxy--reference--group-004.md#canonical-3230100330101132-1031010002311301-3103300123310201-2123313031101131-3230002331032101-0223003130132233-1321000112101333-0330003010220032) |
| `http_proxy.more_option.response_cookies_to_add.name` | [http_proxy.more_option.response_cookies_to_add.name](data-sources--proxy--reference--group-004.md#canonical-2203223121231210-1220102232310303-3213302321022123-0033132330132322-2020230210132332-3110302032101121-0011022201331223-2001233103123101) |
| `http_proxy.more_option.response_cookies_to_add.overwrite` | [http_proxy.more_option.response_cookies_to_add.overwrite](data-sources--proxy--reference--group-004.md#canonical-3032333202033300-2131313322330000-1230030000203213-2131013303231223-0223130010002332-1303032330112011-3002131301120320-2232211231202103) |
| `http_proxy.more_option.response_cookies_to_add.samesite_lax` | [http_proxy.more_option.response_cookies_to_add.samesite_lax](data-sources--proxy--reference--group-004.md#canonical-0100331312131203-1211220132311333-2102020011232021-1102302030201201-3012203021230301-3133320302123232-0211213222332223-2111310332320202) |
| `http_proxy.more_option.response_cookies_to_add.samesite_none` | [http_proxy.more_option.response_cookies_to_add.samesite_none](data-sources--proxy--reference--group-004.md#canonical-3211112320230211-3202201132110202-3021223013012010-1210232013220030-0331113120020233-1233200131210313-2132121033122230-0030010120303301) |
| `http_proxy.more_option.response_cookies_to_add.samesite_strict` | [http_proxy.more_option.response_cookies_to_add.samesite_strict](data-sources--proxy--reference--group-004.md#canonical-3310301322003302-0020103110330300-0002202332212232-0301012102233002-0031202130313030-1132323320033023-3201020211003101-0020202033232210) |
| `http_proxy.more_option.response_cookies_to_add.secret_value` | [http_proxy.more_option.response_cookies_to_add.secret_value](data-sources--proxy--reference--group-004.md#canonical-0330122331133332-1132112222233101-2120123103120200-2220202012320331-1300301020223333-3101320200202103-0110232323000202-0203212201200230) |
| `http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info` | [http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info](data-sources--proxy--reference--group-004.md#canonical-1112302211001001-0313120311033130-1202210231033323-0032230231010330-3320321031310221-2233113231020202-0103222123001200-2303202121232020) |
| `http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--proxy--reference--group-004.md#canonical-1000213030013100-3000111232000301-2320321212023200-3023312202232030-1132021012120203-1101301011123211-3303331331201231-1000133210232110) |
| `http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.location` | [http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.location](data-sources--proxy--reference--group-004.md#canonical-0010102132120333-1231311102311111-0313200000120000-2103013111100103-0320222022101032-3010120112113002-3310312011012300-2332333210130032) |
| `http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--proxy--reference--group-004.md#canonical-3212230120220110-3110221231111303-0212023231133103-0231010023213301-2212220011022113-2102110112023203-2302030201132033-1212033101201310) |
| `http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info` | [http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info](data-sources--proxy--reference--group-004.md#canonical-2330320131133003-2203130001113011-3331301101120330-1222133301311212-2120110300213032-2333113003012100-3110101212103220-3102123322033130) |
| `http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.provider_ref](data-sources--proxy--reference--group-004.md#canonical-1032001023333022-3103110001101321-2131211302311222-1333111211031212-2312012122212122-0100010302221220-0300300122000212-3120201130313012) |
| `http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.url` | [http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.url](data-sources--proxy--reference--group-004.md#canonical-2033131202313003-3203211110032230-1333221332333201-0010332323103122-2101012111301311-2200300112100133-0221223020311032-2023131020023100) |
| `http_proxy.more_option.response_cookies_to_add.value` | [http_proxy.more_option.response_cookies_to_add.value](data-sources--proxy--reference--group-004.md#canonical-0222030122213021-2312121333100133-2030013121322131-0323120211101332-0122310123110303-1000031130120031-1230211233002210-1312112102030211) |
| `http_proxy.more_option.response_cookies_to_remove` | [http_proxy.more_option.response_cookies_to_remove](data-sources--proxy--reference--group-003.md#canonical-2113122013300121-1023223313203123-3103111120130023-3313031223211330-2133223003000001-3331301000020003-1332321203031002-2110103013230132) |
| `http_proxy.more_option.response_headers_to_add` | [http_proxy.more_option.response_headers_to_add](data-sources--proxy--reference--group-004.md#canonical-2300103121122010-3001301033010231-0133213333231212-3101121230301122-1330123230021300-1310022212311120-2031131303111210-2132110220313001) |
| `http_proxy.more_option.response_headers_to_add.append` | [http_proxy.more_option.response_headers_to_add.append](data-sources--proxy--reference--group-004.md#canonical-3313012113311301-3213110112302230-3010210021003001-2221030330301032-2102223031312100-1013101222022111-2321001031223111-3133330130001210) |
| `http_proxy.more_option.response_headers_to_add.name` | [http_proxy.more_option.response_headers_to_add.name](data-sources--proxy--reference--group-004.md#canonical-0123322303020310-3320022311222001-0301022222110323-3030223320110021-0300321323303231-1220331100010012-3013113023031210-2232202200010201) |
| `http_proxy.more_option.response_headers_to_add.secret_value` | [http_proxy.more_option.response_headers_to_add.secret_value](data-sources--proxy--reference--group-004.md#canonical-3001022123232121-1120011210200220-3202200010120130-2323123310200100-1121112020111222-1113031101111323-1113113003020333-0301323320033313) |
| `http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info` | [http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info](data-sources--proxy--reference--group-004.md#canonical-3100211230013112-2013021001132102-3111023312320020-3301231130112110-0321010130310212-3213222101220330-3021200332103211-3301023011023211) |
| `http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--proxy--reference--group-004.md#canonical-2312233123033323-3110311302111022-0112101010001302-3211202123311121-1121110332111011-0232020112011131-2321102331031210-1103313233302102) |
| `http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.location` | [http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.location](data-sources--proxy--reference--group-004.md#canonical-1012312321132231-2011330231303332-0302311023202302-2003103101131310-0211123003223113-3031001031102033-3333101021122012-0132323023223230) |
| `http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--proxy--reference--group-004.md#canonical-1122031120132133-2032103321300312-0000022001221321-0230212332312020-3021322010011012-2311130302321231-3330113233011133-2003013221131123) |
| `http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info` | [http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info](data-sources--proxy--reference--group-004.md#canonical-0103002101130030-1201031220223010-0123311110333220-3220331011023301-2132123320310303-1012222322223112-0230210210102112-0123121210100211) |
| `http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.provider_ref` | [http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.provider_ref](data-sources--proxy--reference--group-004.md#canonical-0000111202001022-2113022311110011-1322333012230023-2100231011330231-2232011131100112-2113330202232212-3111203302121030-1013211202320030) |
| `http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.url` | [http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.url](data-sources--proxy--reference--group-004.md#canonical-0220200130102331-2332133020310220-2101021320201012-1311202132313001-3201110130032221-1231321132130231-3033100233131003-0233331130232231) |
| `http_proxy.more_option.response_headers_to_add.value` | [http_proxy.more_option.response_headers_to_add.value](data-sources--proxy--reference--group-004.md#canonical-0303232221131002-0213001331122131-0021203013102120-0120320123102102-0003111033132013-1020023101323313-0221200300313330-1310122120101130) |
| `http_proxy.more_option.response_headers_to_remove` | [http_proxy.more_option.response_headers_to_remove](data-sources--proxy--reference--group-003.md#canonical-0202123213213210-2301333312111103-3013031003323121-3301331221010203-1022120101001100-0003032011000001-1321102223223012-0301231030211322) |
| `id` | [ID](data-sources--proxy--reference--group-001.md#canonical-1130221001032033-1030030112333213-1123022030330001-1311301012302002-2120010301130031-3133033212133020-0130201113230032-1303202110011102) |
| `labels` | [labels](data-sources--proxy--reference--group-001.md#canonical-1213130031000122-3022010201211232-0323311333201333-1210322012230330-3132323121100221-0203132212232103-3033013311102232-0110021230113013) |
| `name` | [name](data-sources--proxy--reference--group-001.md#canonical-3103233300113031-0132101231131312-0031021021302033-0310201110201033-0001301210003222-3012303230123133-1133003313010131-1210312013301311) |
| `namespace` | [namespace](data-sources--proxy--reference--group-001.md#canonical-1323121121310331-1300013312012021-0001100133232222-2033223032200103-2331112110110111-1332000313033131-2320330301301013-1223312101031122) |
| `no_forward_proxy_policy` | [no_forward_proxy_policy](data-sources--proxy--reference--group-004.md#canonical-2310011101002013-3322323123332321-0121203101211110-3230111130002230-1130201221322013-0333330001310331-3020301030030112-3222302010330102) |
| `no_interception` | [no_interception](data-sources--proxy--reference--group-004.md#canonical-1013023110032033-3302100211223003-1013000321132320-3030301322333203-1220320210021310-3300313321011210-2333330211031221-1231213103320203) |
| `site_local_inside_network` | [site_local_inside_network](data-sources--proxy--reference--group-005.md#canonical-3011333013231301-3311000112331023-1233112001130020-2301120100033013-0202330003110123-3112302112010301-0101132030321312-1221012000013202) |
| `site_local_network` | [site_local_network](data-sources--proxy--reference--group-005.md#canonical-2120312123230111-3221222202321003-3323023300110113-2301113023020122-0303322012003223-0223312123130333-1302023103013221-2212023311032121) |
| `site_virtual_sites` | [site_virtual_sites](data-sources--proxy--reference--group-005.md#canonical-0112331330121211-0322102123023021-1032313031312220-3300101133311202-1313122111333003-0221311103130312-2122111322222100-2203321301210323) |
| `site_virtual_sites.advertise_where` | [site_virtual_sites.advertise_where](data-sources--proxy--reference--group-005.md#canonical-1200222300023123-0302000223123311-2123221321233322-1010333233212212-2321202302300322-2012321212102331-3303103110300100-3303320120320112) |
| `site_virtual_sites.advertise_where.port` | [site_virtual_sites.advertise_where.port](data-sources--proxy--reference--group-005.md#canonical-1003300213011310-1131002231013312-2001022031013312-0121210110123100-1210120023012300-2113300211020220-1300003101302022-0013221322100331) |
| `site_virtual_sites.advertise_where.site` | [site_virtual_sites.advertise_where.site](data-sources--proxy--reference--group-005.md#canonical-2012232123210032-2130001211033133-0303112231020220-1031303102323210-1122110133122322-2220130102120022-2012222330012033-2011300322132030) |
| `site_virtual_sites.advertise_where.site.ip` | [site_virtual_sites.advertise_where.site.ip](data-sources--proxy--reference--group-005.md#canonical-1302131332333120-1102321233132020-2302312223133223-0230201112113333-2023022101000223-3120201021121003-2120112020133113-3032022223022133) |
| `site_virtual_sites.advertise_where.site.network` | [site_virtual_sites.advertise_where.site.network](data-sources--proxy--reference--group-005.md#canonical-1122332323303310-3312233030010333-2221100001312001-0222200032012321-1313232000200032-1200333331021300-1002133123002101-0021322201133333) |
| `site_virtual_sites.advertise_where.site.site` | [site_virtual_sites.advertise_where.site.site](data-sources--proxy--reference--group-005.md#canonical-2012113200013210-2311222310130020-1222230031123000-0012231310010203-1112223133301330-1330022320000303-1131213031302031-1300112130230223) |
| `site_virtual_sites.advertise_where.site.site.name` | [site_virtual_sites.advertise_where.site.site.name](data-sources--proxy--reference--group-005.md#canonical-0012223120102133-3213112123330320-2213332311232003-2331113031020023-3200310320320223-1113221300330310-3221103311103200-2333121012203323) |
| `site_virtual_sites.advertise_where.site.site.namespace` | [site_virtual_sites.advertise_where.site.site.namespace](data-sources--proxy--reference--group-005.md#canonical-2033122221300231-0112332232003212-2221120000202000-0021303332001113-2101122323322311-2213310323313321-1023312001221300-2133320313010302) |
| `site_virtual_sites.advertise_where.site.site.tenant` | [site_virtual_sites.advertise_where.site.site.tenant](data-sources--proxy--reference--group-005.md#canonical-1331132003120312-2310221101112021-3100333113333103-3313200330003332-2133322123333000-1312312121331031-2131222322020212-1001303131321030) |
| `site_virtual_sites.advertise_where.use_default_port` | [site_virtual_sites.advertise_where.use_default_port](data-sources--proxy--reference--group-005.md#canonical-3113233331220322-3002231203202102-3112012102020323-3002010222031113-1020101312111023-3020322300132123-2232233003113231-1133211003031123) |
| `site_virtual_sites.advertise_where.virtual_site` | [site_virtual_sites.advertise_where.virtual_site](data-sources--proxy--reference--group-005.md#canonical-0122302112031312-2233031012313202-2222333322110221-1032100112002110-1223131202102313-2003322332112302-2011203233121300-0131010111020301) |
| `site_virtual_sites.advertise_where.virtual_site.network` | [site_virtual_sites.advertise_where.virtual_site.network](data-sources--proxy--reference--group-005.md#canonical-0222120033110300-0222120010010000-3112100331310022-0123123022321231-1022120003330213-2001010322213310-2212220103223230-3302330231222233) |
| `site_virtual_sites.advertise_where.virtual_site.virtual_site` | [site_virtual_sites.advertise_where.virtual_site.virtual_site](data-sources--proxy--reference--group-005.md#canonical-3201133000030101-2111222220311312-1223130313233220-0003221131111102-3001333133233332-1111313322322230-2032222020330202-3202301320120331) |
| `site_virtual_sites.advertise_where.virtual_site.virtual_site.name` | [site_virtual_sites.advertise_where.virtual_site.virtual_site.name](data-sources--proxy--reference--group-005.md#canonical-1201303031032001-1001122221023331-3320003123203113-1132120311222212-0322333312112103-0100123030300021-1300222133300112-3030230332331030) |
| `site_virtual_sites.advertise_where.virtual_site.virtual_site.namespace` | [site_virtual_sites.advertise_where.virtual_site.virtual_site.namespace](data-sources--proxy--reference--group-005.md#canonical-2212033210031212-3310110313112022-3023000013120231-2133003010321322-0311122100133321-3321322031031220-2011330021220100-0011310223203122) |
| `site_virtual_sites.advertise_where.virtual_site.virtual_site.tenant` | [site_virtual_sites.advertise_where.virtual_site.virtual_site.tenant](data-sources--proxy--reference--group-005.md#canonical-0332021221300200-0222232011123103-0233323231032102-1212312113123021-1233032122322330-0332003210200322-1333012322002222-3032021030131001) |
| `tls_intercept` | [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-2121232312132030-1021002130200310-0121113113303233-3010003311301312-3031223213223011-2223320220020023-0033302022021110-1101201020202012) |
| `tls_intercept.custom_certificate` | [tls_intercept.custom_certificate](data-sources--proxy--reference--group-005.md#canonical-3232201211233031-3302012031311123-3001022011213013-3030010132310122-1012011300203110-2120201011302001-2312333002221232-2013321321332302) |
| `tls_intercept.custom_certificate.certificate_url` | [tls_intercept.custom_certificate.certificate_url](data-sources--proxy--reference--group-005.md#canonical-3032113001001113-0020202300333322-2021210112230303-0303003102320003-0302033300021212-0320103321201300-0232333012323230-3310330020032210) |
| `tls_intercept.custom_certificate.custom_hash_algorithms` | [tls_intercept.custom_certificate.custom_hash_algorithms](data-sources--proxy--reference--group-005.md#canonical-0130300112022033-0030312103101302-2022332303012121-2223121110330330-1012220323322231-3210232222201113-0223302013033130-0122303212333021) |
| `tls_intercept.custom_certificate.custom_hash_algorithms.hash_algorithms` | [tls_intercept.custom_certificate.custom_hash_algorithms.hash_algorithms](data-sources--proxy--reference--group-005.md#canonical-3200323301112120-3101303120100111-2021213132221023-3212121111103233-1022010200113323-2130201330103311-2202010233201012-2033110311302111) |
| `tls_intercept.custom_certificate.description_spec` | [tls_intercept.custom_certificate.description_spec](data-sources--proxy--reference--group-005.md#canonical-2020220001230310-1310333302301032-0112320213131013-2131311131001203-1233030201122032-0221031232311101-2331332332010021-2022033103331101) |
| `tls_intercept.custom_certificate.disable_ocsp_stapling` | [tls_intercept.custom_certificate.disable_ocsp_stapling](data-sources--proxy--reference--group-005.md#canonical-3201322323312302-3101312131312220-2022201233311300-0021122010211331-3201230333320103-1221301010202003-1100011032121030-3221131332121201) |
| `tls_intercept.custom_certificate.private_key` | [tls_intercept.custom_certificate.private_key](data-sources--proxy--reference--group-005.md#canonical-1321110132331221-3323111110320330-3201123022231032-1020320300200320-0200110113213103-1200221012112103-0011002102322123-2311311312011132) |
| `tls_intercept.custom_certificate.private_key.blindfold_secret_info` | [tls_intercept.custom_certificate.private_key.blindfold_secret_info](data-sources--proxy--reference--group-005.md#canonical-2130120001211200-2023110233223223-0222210032032110-1233200000232212-2311301002222312-0120313103023031-1000020000022222-1123030132123131) |
| `tls_intercept.custom_certificate.private_key.blindfold_secret_info.decryption_provider` | [tls_intercept.custom_certificate.private_key.blindfold_secret_info.decryption_provider](data-sources--proxy--reference--group-005.md#canonical-3133020322322213-2333332200311132-0130310312200321-2122211023321002-1202113031210112-0203031120110211-3301223211131021-1310203132122120) |
| `tls_intercept.custom_certificate.private_key.blindfold_secret_info.location` | [tls_intercept.custom_certificate.private_key.blindfold_secret_info.location](data-sources--proxy--reference--group-005.md#canonical-3300030203302003-3021022002320200-2212330120322031-3023010201301310-2123012002331131-2133310031211122-2231200320122121-1002313301032322) |
| `tls_intercept.custom_certificate.private_key.blindfold_secret_info.store_provider` | [tls_intercept.custom_certificate.private_key.blindfold_secret_info.store_provider](data-sources--proxy--reference--group-005.md#canonical-3323231132013021-1201331320010211-0210103100312033-3011001303330321-0132133201011233-1303003000210203-2301001312231223-2130133020133313) |
| `tls_intercept.custom_certificate.private_key.clear_secret_info` | [tls_intercept.custom_certificate.private_key.clear_secret_info](data-sources--proxy--reference--group-005.md#canonical-1220233030332323-0010003022111130-1221031230130102-3021122123302321-2221222223133323-0133001022232301-2021221101210130-2333211301320002) |
| `tls_intercept.custom_certificate.private_key.clear_secret_info.provider_ref` | [tls_intercept.custom_certificate.private_key.clear_secret_info.provider_ref](data-sources--proxy--reference--group-005.md#canonical-2023012032300020-3130130122203230-0200111221020000-3021113231330001-2233110010301332-0203331020321031-3300300113232133-1203013223223222) |
| `tls_intercept.custom_certificate.private_key.clear_secret_info.url` | [tls_intercept.custom_certificate.private_key.clear_secret_info.url](data-sources--proxy--reference--group-005.md#canonical-3003332223312231-3301323201013300-1201200303001212-0030101103331331-1302011320113303-1203213132132320-2223023212111003-1321313220320020) |
| `tls_intercept.custom_certificate.use_system_defaults` | [tls_intercept.custom_certificate.use_system_defaults](data-sources--proxy--reference--group-005.md#canonical-3210101010210211-0003110102111231-0313232033230131-1320322121023121-0211233123211003-0103112203310321-3320111020221230-0113222113120201) |
| `tls_intercept.enable_for_all_domains` | [tls_intercept.enable_for_all_domains](data-sources--proxy--reference--group-005.md#canonical-1330022123231330-2322122032123301-0312121302133213-0322313230020103-0033201221231313-1133301231300323-0012010313013202-1302211022331311) |
| `tls_intercept.policy` | [tls_intercept.policy](data-sources--proxy--reference--group-005.md#canonical-2101031213001101-1213003123113223-1320233222033313-3033132022011020-1230122203212100-3230122133133232-3210213110311030-2201310011021210) |
| `tls_intercept.policy.interception_rules` | [tls_intercept.policy.interception_rules](data-sources--proxy--reference--group-005.md#canonical-1113032002213112-3322213023132212-3103021213030223-0312300312212233-0303223123321130-3231332133031300-3212332320202121-1301323320211021) |
| `tls_intercept.policy.interception_rules.disable_interception` | [tls_intercept.policy.interception_rules.disable_interception](data-sources--proxy--reference--group-005.md#canonical-3230310311131200-2202011103231103-1133330322213212-3011112320010120-2313232303213000-1102322223010202-1222012233231310-0213032033000332) |
| `tls_intercept.policy.interception_rules.domain_match` | [tls_intercept.policy.interception_rules.domain_match](data-sources--proxy--reference--group-005.md#canonical-2000211212020220-0323011121211120-3022030001001331-2333103121200000-1100332222323320-2303213003233321-3100201022110311-0132232200311121) |
| `tls_intercept.policy.interception_rules.domain_match.exact_value` | [tls_intercept.policy.interception_rules.domain_match.exact_value](data-sources--proxy--reference--group-005.md#canonical-3121132113013333-2011011013232313-3020113110312310-0012302303230203-0302013022233011-3210033020331303-2102200101103230-1203301202033122) |
| `tls_intercept.policy.interception_rules.domain_match.regex_value` | [tls_intercept.policy.interception_rules.domain_match.regex_value](data-sources--proxy--reference--group-005.md#canonical-1210110210120023-1103002303201111-1103113101110233-3132303330011010-0203300113020133-3222233222001221-2230310212311113-2100312032012320) |
| `tls_intercept.policy.interception_rules.domain_match.suffix_value` | [tls_intercept.policy.interception_rules.domain_match.suffix_value](data-sources--proxy--reference--group-005.md#canonical-1010202121322323-3021203112223202-2113023231112103-2020231202201303-2102203201030300-1013032132130231-1300331012111300-2303021121300211) |
| `tls_intercept.policy.interception_rules.enable_interception` | [tls_intercept.policy.interception_rules.enable_interception](data-sources--proxy--reference--group-005.md#canonical-1100020223102300-2322222300101130-2112310310233020-1220231110200301-3020010330133230-1301022323201233-1320122113213301-2310222122330113) |
| `tls_intercept.trusted_ca_url` | [tls_intercept.trusted_ca_url](data-sources--proxy--reference--group-005.md#canonical-0102133020010102-2213012022123223-3310030013031013-0203321130301312-2202121010022122-3010200322231312-1221222212020132-2033120022201111) |
| `tls_intercept.volterra_certificate` | [tls_intercept.volterra_certificate](data-sources--proxy--reference--group-005.md#canonical-3232030101211322-0013110331131102-1232303132112113-0233211103033200-1033002221322313-2220003312130223-3301221323232332-3331032222313221) |
| `tls_intercept.volterra_trusted_ca` | [tls_intercept.volterra_trusted_ca](data-sources--proxy--reference--group-005.md#canonical-3310300332010012-0031300302033223-2003332123223232-2201023223220121-2313332013010022-0102102113300211-3023130302230021-3310013030133301) |

<a id="canonical-2213113013300313-0020111122011302-2231233220020000-0031023111301110-2222323311112113-1032001030003100-3301333013321011-0123202302110231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `active_forward_proxy_policies` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- active_forward_proxy_policies

<a id="canonical-0113112101100330-3012113212033210-2300302031303022-0103202033010002-2001100310212330-3321202003130323-1203132322032202-2300022112133033"></a>

Type: `"single"`. Computed.

\[OneOf: active\_forward\_proxy\_policies, no\_forward\_proxy\_policy; Default:
no\_forward\_proxy\_policy\] Ordered List of Forward Proxy Policies active.

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

- [active_forward_proxy_policies](data-sources--proxy--reference--group-001.md#canonical-0113112101100330-3012113212033210-2300302031303022-0103202033010002-2001100310212330-3321202003130323-1203132322032202-2300022112133033)
- [no_forward_proxy_policy](data-sources--proxy--reference--group-004.md#canonical-2310011101002013-3322323123332321-0121203101211110-3230111130002230-1130201221322013-0333330001310331-3020301030030112-3222302010330102)

Select alternatives according to the provider validators above.

<a id="canonical-2132222223133002-0302233321121230-1303010202221021-3113011310203211-0313122311030202-0201222232202103-1001031302300323-1311001001211312"></a>

### Direct properties for `active_forward_proxy_policies`

- [forward_proxy_policies](data-sources--proxy--reference--group-001.md#canonical-2002320100212011-1320223330220032-3221313013230231-3210120122003112-1003233202313032-3323202300031111-2301110212321233-3233133110013332): complete subsection reference.

<a id="canonical-2002320100212011-1320223330220032-3221313013230231-3210120122003112-1003233202313032-3323202300031111-2301110212321233-3233133110013332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `active_forward_proxy_policies.forward_proxy_policies` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [active_forward_proxy_policies](data-sources--proxy--reference--group-001.md#canonical-2213113013300313-0020111122011302-2231233220020000-0031023111301110-2222323311112113-1032001030003100-3301333013321011-0123202302110231)
- active_forward_proxy_policies.forward_proxy_policies

<a id="canonical-3223202303233120-3303312123211202-3230113310312121-1202000222113312-0220031320133220-1202132101012222-1223330032222033-3020232113130020"></a>

Type: `"list"`. Computed.

Ordered List of Forward Proxy Policies active.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-0133323111010211-2213030002331302-1223301131133111-2021031103112313-3221331300020100-2103232212210202-2011113202111230-1002210002110323"></a>

### Direct properties for `active_forward_proxy_policies.forward_proxy_policies`

<a id="canonical-0033222202020033-0222320011332203-0301111221233023-3313203333112103-2010300223312221-0300012033032113-3003102112322220-0010110301213120"></a>

#### `active_forward_proxy_policies.forward_proxy_policies.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0103000020310112-2213133132031210-2100321012310113-3131100323330323-2010032111102200-1231113131113000-3033302303000000-2233102120210110"></a>

<a id="canonical-0100021202130002-3003123000020201-2233313333133122-3302301120112100-0301212203321120-2133321203022203-3321212201120200-1203013211121121"></a>

#### `active_forward_proxy_policies.forward_proxy_policies.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1311100302230330-0032213030132301-2312233330311202-3232121001012230-1021320212330202-0033200032030223-0113221013123310-0130021032011232"></a>

<a id="canonical-0013222221011300-2030030320211303-0322221330121121-2112113311021131-0033000023313010-1001320300330312-2110031013321131-3012230020221023"></a>

#### `active_forward_proxy_policies.forward_proxy_policies.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1222310002221222-1321223032311031-2020310323203003-1210330021211320-0013020220203303-3203312313122011-3000002100223313-0101333020130003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `do_not_advertise` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- do_not_advertise

<a id="canonical-1111013301201012-2330232112312001-2103313002312231-0111011120033031-0013111131103130-3330111233202133-2120132303021132-2203322221101333"></a>

Type: `["object", {}]`. Computed.

\[OneOf: do\_not\_advertise, site\_virtual\_sites\] Configuration parameter for do not advertise.

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

OneOf alternatives in this subsection:

- [do_not_advertise](data-sources--proxy--reference--group-001.md#canonical-1111013301201012-2330232112312001-2103313002312231-0111011120033031-0013111131103130-3330111233202133-2120132303021132-2203322221101333)
- [site_virtual_sites](data-sources--proxy--reference--group-005.md#canonical-0112331330121211-0322102123023021-1032313031312220-3300101133311202-1313122111333003-0221311103130312-2122111322222100-2203321301210323)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- dynamic_proxy

<a id="canonical-3300012110333323-0313303121121121-1331003000322330-3313312302103220-3210020033323300-0212112321033131-0003101301121210-3311300223113220"></a>

Type: `"single"`. Computed.

\[OneOf: dynamic\_proxy, http\_proxy\] Configuration parameter for dynamic proxy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dns_choice": "[\"disable_dns_masquerade\",\"enable_dns_masquerade\"]",
  "x-ves-oneof-field-proxy_choice": "[\"http_proxy\",\"https_proxy\",\"sni_proxy\"]"
}
```

OneOf alternatives in this subsection:

- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3300012110333323-0313303121121121-1331003000322330-3313312302103220-3210020033323300-0212112321033131-0003101301121210-3311300223113220)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-2121010022301312-1300233121020222-1323121211031111-0131133120001130-0000032031101232-1000113012222203-2213302130100231-2210233233221031)

Select alternatives according to the provider validators above.

<a id="canonical-3033323232000201-0231300002330213-2300211321200303-1223131023013233-3122213121232113-2320031321303210-3300233123212021-2113211200113230"></a>

### Direct properties for `dynamic_proxy`

- [disable_dns_masquerade](data-sources--proxy--reference--group-001.md#canonical-0001101103333020-0011211230303113-2323233010210223-3023203313031222-2101103123231330-2300102331123310-1120301300210210-3120220210001130): complete subsection reference.

<a id="canonical-2102010211002230-1012113212113010-0230200110221020-0322311033302313-0111002122110103-1110032011212313-0210101211333210-3201030101120223"></a>

<a id="canonical-0322003111133323-1320020232000222-0331011002202103-0021033203121010-3203231333301330-3101033121232300-3310303330110133-3032330003322222"></a>

#### `dynamic_proxy.domains` property

Type: `["list", "string"]`. Computed.

A list of Domains to be proxied. Wildcard hosts are supported in the suffix or prefix form

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [enable_dns_masquerade](data-sources--proxy--reference--group-001.md#canonical-2120100023112023-2201321122213302-1113233313232010-0001112322301303-1133001303322133-3312111310100030-0133303310331022-1211300202000102): complete subsection reference.

- [http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000): complete subsection reference.

- [https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220): complete subsection reference.

- [sni_proxy](data-sources--proxy--reference--group-003.md#canonical-3201203013323203-2030201011013320-2110012213112021-0333333012303323-1113131112130002-1333330220013232-0212100100120203-2000200313330302): complete subsection reference.

<a id="canonical-0001101103333020-0011211230303113-2323233010210223-3023203313031222-2101103123231330-2300102331123310-1120301300210210-3120220210001130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.disable_dns_masquerade` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- dynamic_proxy.disable_dns_masquerade

<a id="canonical-3030020103222030-0001110010222311-0220032123002311-3222122011000332-3031313220303211-1021032221320220-3332132130301120-3230332303122111"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable DNS masquerade.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2120100023112023-2201321122213302-1113233313232010-0001112322301303-1133001303322133-3312111310100030-0133303310331022-1211300202000102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.enable_dns_masquerade` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- dynamic_proxy.enable_dns_masquerade

<a id="canonical-2202230330010222-1201110002013011-1212113012210011-2200032131101320-3011012201022202-2033103202322312-0113100331230101-1333012213003120"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable DNS masquerade.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- dynamic_proxy.http_proxy

<a id="canonical-1300103332011102-3110211000323203-3312121232322222-2123330223011231-2120001322320111-3200231200233213-0220313023303020-2021202211232301"></a>

Type: `"single"`. Computed.

Dynamic HTTP Proxy Type. Parameters for dynamic HTTP proxy.

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

<a id="canonical-0301220311223222-0311333132011103-0130023323212323-2022223103210033-2103220330201300-2030302001312002-2011233031320323-1320213010322232"></a>

### Direct properties for `dynamic_proxy.http_proxy`

- [more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201): complete subsection reference.

<a id="canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- dynamic_proxy.http_proxy.more_option

<a id="canonical-3101311213023200-2211010131221101-1300310131110022-2101131101100121-1303203223302122-0130123012333100-3011103221300133-3322122013013011"></a>

Type: `"single"`. Computed.

This defines various OPTIONS to define a route.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-max_requests_per_connection_choice": "[\"max_requests_per_connection\",\"no_request_limit_per_connection\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-strict_sni_host_header_check_choice": "[]"
}
```

<a id="canonical-2321322103012323-0321123302021130-2132122303123102-3133123131021012-2211213000301230-1020020033013113-0011300022322122-1133122113323222"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option`

- [buffer_policy](data-sources--proxy--reference--group-001.md#canonical-2112113001020132-0110332202131113-0021200011312123-0010031002212123-0223022001100023-2010132023033232-1232330330122203-3002210012211100): complete subsection reference.

- [compression_params](data-sources--proxy--reference--group-001.md#canonical-1220222323113313-2333213122220121-2121213311112002-3123211111201221-3221111323303203-3233001210121203-0310322131301123-1112223000211322): complete subsection reference.

<a id="canonical-1022133221122202-1212331330033211-1312211000001103-3013202201300323-1110300122111311-2030301113012030-0200301033331310-2333013322312331"></a>

<a id="canonical-3211223331203200-2202111323113101-0303222112112310-3212232122231010-2201222221303130-1221011102223030-3211102131132000-3333011312011311"></a>

#### `dynamic_proxy.http_proxy.more_option.custom_errors` property

Type: `["map", "string"]`. Computed.

Map of integer error codes as keys and string values that can be used to provide custom HTTP pages
for each error code. Key of the map can be either response code class or HTTP Error code. Response
code classes for key is configured as follows 3 -- for 3xx response code class 4 -- for 4xx response
code class 5 -- for 5xx response code class Value of the map is string which represents custom HTTP
responses. Specific response code takes preference when both response code and response code class
matches for a request.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 16
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "ranges": [
        [
          3,
          3
        ],
        [
          4,
          4
        ],
        [
          5,
          5
        ],
        [
          300,
          599
        ]
      ],
      "type": "uint32-string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.uint32.ranges": "3,4,5,300-599",
      "ves.io.schema.rules.map.max_pairs": "16",
      "ves.io.schema.rules.map.values.string.max_len": "65536",
      "ves.io.schema.rules.map.values.string.uri_ref": "true"
    },
    "values": {
      "format": "uri-reference",
      "maxLength": 65536,
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
    "ves.io.schema.rules.map.keys.uint32.ranges": "3,4,5,300-599",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "65536",
    "ves.io.schema.rules.map.values.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.uint32.ranges": "3,4,5,300-599",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "65536",
    "ves.io.schema.rules.map.values.string.uri_ref": "true"
  }
}
```

<a id="canonical-2232323331212222-1112021203322001-2302002102312111-3313213131032032-1102030103113102-3013222012221001-3102012323313230-3333323232113210"></a>

<a id="canonical-0331221322022000-2301302112211021-1322102123212112-2022213001130203-3123313221331023-3133330011333000-2102033311002103-2122002130013330"></a>

#### `dynamic_proxy.http_proxy.more_option.disable_default_error_pages` property

Type: `"bool"`. Computed.

Disable the use of default F5XC error pages.

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

- [disable_path_normalize](data-sources--proxy--reference--group-001.md#canonical-3203323212312201-1231032110031320-2020103200220313-3212233112202321-1030133022313212-3212101120003222-0012012010330012-2000103003213133): complete subsection reference.

- [enable_path_normalize](data-sources--proxy--reference--group-001.md#canonical-3200333303133021-3231310312203010-2233202233001200-0113213233012032-0020231311133133-1130212020030322-2121010020313133-2111331200210323): complete subsection reference.

<a id="canonical-3211123313011203-1321001221312312-1121232100011220-1220210030002232-2201201232001122-2103232002030230-2030023013232331-0201012332003023"></a>

<a id="canonical-3221001111233011-1302330310321002-1331102012332322-1211222110002033-2333021313002331-3302322133330231-2303100201321322-0102101302211112"></a>

#### `dynamic_proxy.http_proxy.more_option.idle_timeout` property

Type: `"number"`. Computed.

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.
The stream is terminated with an HTTP 504 (Gateway Timeout) error code if no upstream response
header has been received, otherwise the stream is reset.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 3600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "3600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "3600000"
  }
}
```

<a id="canonical-0331032133232121-2233003012131223-0032012131212123-1200312020112211-3303310202321202-1331132102323103-0120020122310320-3030223202221113"></a>

<a id="canonical-2032002002131023-3321323322102231-2130121212312001-0123202232030322-1303312000033322-0000310121002321-1312222121132232-1330133113032122"></a>

#### `dynamic_proxy.http_proxy.more_option.max_request_header_size` property

Type: `"number"`. Computed.

The maximum request header size for downstream connections, in KiB. An HTTP 431 (Request Header
Fields Too Large) error code is sent for requests that exceed this size.

If multiple load balancers share the same advertise\_policy, the highest value configured across all
such load balancers is used for all the load balancers in question.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 96,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "96"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "96"
  }
}
```

<a id="canonical-2301223333323001-0022112210023021-3222012003101100-3300223111212312-0033032103001232-2011200212020211-3211111301210223-2302231002232221"></a>

<a id="canonical-2011302202202000-1110321222002302-1203311031320133-0113111123101010-0202000001323132-1123120000033100-0310110200030110-2230302020320303"></a>

#### `dynamic_proxy.http_proxy.more_option.max_requests_per_connection` property

Type: `"number"`. Computed.

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests a
downstream client can send over a single connection to Envoy. Enter a value &gt;=1 to define the
request limit per connection.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

- [no_request_limit_per_connection](data-sources--proxy--reference--group-001.md#canonical-1022212012133233-2230231101300000-2201203102103101-2233110103201202-3020303221310220-2210131311112023-3312113102322221-2313213012032201): complete subsection reference.

- [request_cookies_to_add](data-sources--proxy--reference--group-001.md#canonical-2303301033010202-3133022000003110-1031331200010303-2230302030022310-0100230303310021-0120221220002211-3121010321000033-1002103011232120): complete subsection reference.

<a id="canonical-2010112220311330-3221310033202300-3333321003312022-2203033030000021-2222201312101102-1023101233003131-3013231111030233-3210221101332331"></a>

<a id="canonical-2013213323230110-1312233130332130-0110320230031022-1230313020122020-1120020332210321-2321022111330311-2123302000211313-1200221313133213"></a>

#### `dynamic_proxy.http_proxy.more_option.request_cookies_to_remove` property

Type: `["list", "string"]`. Computed.

List of keys of Cookies to be removed from the HTTP request being sent towards upstream.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [request_headers_to_add](data-sources--proxy--reference--group-002.md#canonical-1003213031100030-3111200032003222-2231310222011322-3102330002003302-2310213002030322-0102220300201000-1311312001301103-1113302021310111): complete subsection reference.

<a id="canonical-1233012331102201-0203001131001223-1313011002111222-0113121222210203-2200102223231332-1002011031121333-0011211233313133-1301100330020101"></a>

<a id="canonical-3332322211031331-1011130213020112-0021301011201233-1330321000032130-2320130303310321-3200032232030222-3033302213233111-3330110230112001"></a>

#### `dynamic_proxy.http_proxy.more_option.request_headers_to_remove` property

Type: `["list", "string"]`. Computed.

List of keys of Headers to be removed from the HTTP request being sent towards upstream.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-1001121231112113-2321013313123233-3121011120032112-2230103211203211-2220002211200301-0323022031220331-2232320202211030-3233021220031030): complete subsection reference.

<a id="canonical-1333123000020021-2122322132331113-3302010131303300-0320032313303003-2012201131301132-1200200203003010-1103222201312003-0133312102013233"></a>

<a id="canonical-2232313232322131-3222131013121032-1023323013120202-2201130231102013-1111131003303011-0202122020220200-1033110102300331-0013312213131212"></a>

#### `dynamic_proxy.http_proxy.more_option.response_cookies_to_remove` property

Type: `["list", "string"]`. Computed.

List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire
set-cookie header will be removed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_headers_to_add](data-sources--proxy--reference--group-002.md#canonical-1203020210021221-3100021212030223-0112223000011011-3030123133121320-1131332121102312-0101112333012222-1201023112032320-3300013210320123): complete subsection reference.

<a id="canonical-1312001010100030-2312110320210232-3023322332332132-1220131202220111-3230330222213212-0011031032133123-3320310331220020-2221232102023022"></a>

<a id="canonical-2221102330023023-3331103022023223-3110100223111002-3023000011232022-1012210100020200-3113011022301330-2220211012122023-0003221200020303"></a>

#### `dynamic_proxy.http_proxy.more_option.response_headers_to_remove` property

Type: `["list", "string"]`. Computed.

List of keys of Headers to be removed from the HTTP response being sent towards downstream.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2112113001020132-0110332202131113-0021200011312123-0010031002212123-0223022001100023-2010132023033232-1232330330122203-3002210012211100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.buffer_policy` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- dynamic_proxy.http_proxy.more_option.buffer_policy

<a id="canonical-0222223111112122-0311010333221231-1330112033312323-1110032213002322-1233012130211221-3202122112303132-0233112330220011-2232221003320031"></a>

Type: `"single"`. Computed.

Some upstream applications are not capable of handling streamed data. This config enables buffering
the entire request before sending to upstream application. We can specify the maximum buffer size
and buffer interval with this config.

Buffering can be enabled and disabled at VirtualHost and Route levels Route level buffer
configuration takes precedence.

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

<a id="canonical-0330322321332233-3133230123122212-0012211101322332-3212231122000303-1203001200023230-2011332302132122-1202222002232033-2123031202002122"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option.buffer_policy`

<a id="canonical-2302003030211222-3302133203311203-2111030211113223-1220221232130320-1201202101211023-3121231322010003-1332011030330010-1130102122300220"></a>

#### `dynamic_proxy.http_proxy.more_option.buffer_policy.disabled` property

Type: `"bool"`. Computed.

Disable buffering for a particular route. This is useful when virtual-host has buffering, but we
need to disable it on a specific route. The value of this field is ignored for virtual-host.

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

<a id="canonical-1322201233103223-3212321023013011-0201020010201022-3010300320111222-0301332113200312-0320011022031213-1331131110320202-1122130021321003"></a>

<a id="canonical-0222313000102100-0100110113102110-1010333110110102-0303120023212210-3322023113012123-3202030122000130-1011132231300101-1310032201311323"></a>

#### `dynamic_proxy.http_proxy.more_option.buffer_policy.max_request_bytes` property

Type: `"number"`. Computed.

The maximum request size that the filter will buffer before the connection manager will stop
buffering and return a RequestEntityTooLarge (413) response.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10485760,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```

<a id="canonical-1220222323113313-2333213122220121-2121213311112002-3123211111201221-3221111323303203-3233001210121203-0310322131301123-1112223000211322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.compression_params` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- dynamic_proxy.http_proxy.more_option.compression_params

<a id="canonical-3131231122122120-1311012232200231-0013223201213033-3132203112113013-2121102012312220-3122200200021102-1011311212113201-3221302133332203"></a>

Type: `"single"`. Computed.

Enables loadbalancer to compress dispatched data from an upstream service upon client request. The
content is compressed and then sent to the client with the appropriate headers if either response
and request allow. Only GZIP compression is supported.

By default compression will be skipped when:

A request does NOT contain accept-encoding header. A request includes accept-encoding header, but it
does not contain “gzip” or “\*”. A request includes accept-encoding with “gzip” or “\*” with the
weight “q=0”. Note that the “gzip” will have a higher weight then “\*”. For example, if
accept-encoding is “gzip;q=0,\*;q=1”, the filter will not compress. But if the header is set to
“\*;q=0,gzip;q=1”, the filter will compress. A request whose accept-encoding header includes
“identity”. A response contains a content-encoding header. A response contains a cache-control
header whose value includes “no-transform”. A response contains a transfer-encoding header whose
value includes “gzip”. A response does not contain a content-type value that matches one of the
selected mime-types, which default to application/JavaScript, application/JSON,
application/xhtml+XML, image/svg+XML, text/CSS, text/HTML, text/plain, text/XML. Neither
content-length nor transfer-encoding headers are present in the response. Response size is smaller
than 30 bytes (only applicable when transfer-encoding is not chunked).

When compression is applied:

The content-length is removed from response headers. Response headers contain “transfer-encoding:
chunked” and do not contain “content-encoding” header. The “vary: accept-encoding” header is
inserted on every response.

GZIP Compression Level:

A value which is optimal balance between speed of compression and amount of compression is chosen.

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

<a id="canonical-1333312310013201-1121123223233313-1230232132202032-2012021020302122-0133331231203131-3010333232030130-2220102331002323-2213203030323321"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option.compression_params`

<a id="canonical-0120222323323200-0002230030123330-3102212131121033-0303221031221031-1010312321220102-1212121320312021-3022222131323313-1100031021021230"></a>

#### `dynamic_proxy.http_proxy.more_option.compression_params.content_length` property

Type: `"number"`. Computed.

Minimum response length, in bytes, which will trigger compression. The. Defaults to \`30\`.

Additional upstream details:

The default value is 30.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 30
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "30"
  }
}
```

<a id="canonical-1203112102021003-1103012013331100-0212122112121301-2220302130233030-1122120203332323-3013311121202100-2300213011230321-1010301221303231"></a>

<a id="canonical-3302020012003321-3002231032033222-0103323101323302-0331110311021223-0133332321222332-1001031133210100-1130113110120330-2210313122323112"></a>

#### `dynamic_proxy.http_proxy.more_option.compression_params.content_type` property

Type: `["list", "string"]`. Computed.

Set of strings that allows specifying which mime-types yield compression When this field is not
defined, compression will be applied to the following mime-types: 'application/JavaScript'
'application/JSON', 'application/xhtml+XML' 'image/svg+XML' 'text/CSS' 'text/HTML' 'text/plain'
'text/XML'.

Additional upstream details:

Set of strings that allows specifying which mime-types yield compression When this field is not
defined, compression will be applied to the following mime-types: "application/JavaScript"
"application/JSON", "application/xhtml+XML" "image/svg+XML" "text/CSS" "text/HTML" "text/plain"
"text/XML"

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 50,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 50,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1302031213300313-1013113230021212-3331032202031310-2001023212300000-2310210233111123-0312010313110220-0331322010220030-3122202102103033"></a>

<a id="canonical-3033031323121211-3120332320032200-3301330013023013-3020323133331301-3232131323331122-1332212103313331-2100322321000122-0200120021103123"></a>

#### `dynamic_proxy.http_proxy.more_option.compression_params.disable_on_etag_header` property

Type: `"bool"`. Computed.

If true, disables compression when the response contains an etag header. When it is false, weak
etags will be preserved and the ones that require strong validation will be removed.

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

<a id="canonical-2111101123033202-0011333333221022-0030311000103111-3203303233133113-2331220323230102-1121303311210103-1201223120331231-2110002012303222"></a>

<a id="canonical-1230023013312221-2100203202313011-1133130321112210-3310130213121101-3323301132311120-2031133111113021-3100313201231312-1033202030102133"></a>

#### `dynamic_proxy.http_proxy.more_option.compression_params.remove_accept_encoding_header` property

Type: `"bool"`. Computed.

If true, removes accept-encoding from the request headers before dispatching it to the upstream so
that responses do not GET compressed before reaching the filter.

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

<a id="canonical-3203323212312201-1231032110031320-2020103200220313-3212233112202321-1030133022313212-3212101120003222-0012012010330012-2000103003213133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.disable_path_normalize` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- dynamic_proxy.http_proxy.more_option.disable_path_normalize

<a id="canonical-1113233131312122-3130102222100323-3231320311313003-1310221021330023-3030222122111001-2101022103200131-0013322303312130-3131013210211301"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3200333303133021-3231310312203010-2233202233001200-0113213233012032-0020231311133133-1130212020030322-2121010020313133-2111331200210323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.enable_path_normalize` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- dynamic_proxy.http_proxy.more_option.enable_path_normalize

<a id="canonical-1013200221322220-2303203122231100-2220211203312111-3131012311313011-3032010320121222-3032020223110100-1022222233312332-1212000331030330"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1022212012133233-2230231101300000-2201203102103101-2233110103201202-3020303221310220-2210131311112023-3312113102322221-2313213012032201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.no_request_limit_per_connection` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- dynamic_proxy.http_proxy.more_option.no_request_limit_per_connection

<a id="canonical-2000323201012320-3100103022020310-3123303320001201-1030021200210023-1030001230021201-2030100323302131-3122122130100233-0310212333023031"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no request limit per connection.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2303301033010202-3133022000003110-1031331200010303-2230302030022310-0100230303310021-0120221220002211-3121010321000033-1002103011232120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.request_cookies_to_add` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- dynamic_proxy.http_proxy.more_option.request_cookies_to_add

<a id="canonical-2130320311301030-1232110232001001-2303031323322313-3101213313002222-2320230031302332-1233131000130212-3312300322001333-2131311311101133"></a>

Type: `"list"`. Computed.

Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies
specified at this level are applied after cookies from matched Route are applied.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0331011021232230-1130000230210100-3003131013233301-2133202213100111-1103310220013332-1202322030201033-2030330032221131-3231132320023220"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option.request_cookies_to_add`

<a id="canonical-1232112232231320-2103131202200113-1020211021103210-0201230001230332-1032322021233103-2300222032120032-3230322313032102-0001123012122012"></a>

#### `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.name` property

Type: `"string"`. Computed.

Name. Name of the cookie in Cookie header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3311101103220112-1002132233203333-1203021102013101-0202222313221000-0232120033101211-3030330310102031-3233121212313222-3333023122011032"></a>

<a id="canonical-1211113003301010-1032202003323222-1122020010012323-2200233203232013-3112202301013313-0320110122012202-2001121011113221-3301322212131012"></a>

#### `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.overwrite` property

Type: `"bool"`. Computed.

Should the value be overwritten? If true, the value is overwritten to existing values. not
overwrite. Defaults to \`do\`.

Additional upstream details:

If true, the value is overwritten to existing values. Default value is do not overwrite.

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

- [secret_value](data-sources--proxy--reference--group-001.md#canonical-3121223110303320-0020230331202132-1033113120102020-3221031201133132-0230123011323313-2310131201213313-2311321311102301-2223333233002032): complete subsection reference.

<a id="canonical-3023030233021220-0202311202300332-1033012333012113-1103322301333322-3320000023031230-0202323203200020-3323023101023000-2112333212010202"></a>

<a id="canonical-0132331123333102-0212232321021001-3320313103011331-1311012303330211-0320010032113031-3222112113201003-1102121321223001-1031103211120122"></a>

#### `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.value` property

Type: `"string"`. Computed.

Exclusive with \[secret\_value\] Value of the Cookie header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-3121223110303320-0020230331202132-1033113120102020-3221031201133132-0230123011323313-2310131201213313-2311321311102301-2223333233002032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add](data-sources--proxy--reference--group-001.md#canonical-2303301033010202-3133022000003110-1031331200010303-2230302030022310-0100230303310021-0120221220002211-3121010321000033-1002103011232120)
- dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value

<a id="canonical-0333230303230001-1233300210233331-3020210231321132-1031311122212001-1220232330130111-2000032233002001-1213012030001221-1033331012303011"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-1203122313300323-0113302203013313-0223131303102103-1213313031133332-0100031300112320-1112332022101012-3030211112102131-2101311111022202"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value`

- [blindfold_secret_info](data-sources--proxy--reference--group-002.md#canonical-0313013020021030-2002113111010212-2210121310232331-1100011000010003-3223202131130123-0021001311223120-3102322223312111-0100120033123322): complete subsection reference.

- [clear_secret_info](data-sources--proxy--reference--group-002.md#canonical-0232103032301203-0201202110301122-2213321000022112-2323020312201033-0023333220302310-3003330233110230-0003031003202111-3032121231102013): complete subsection reference.
