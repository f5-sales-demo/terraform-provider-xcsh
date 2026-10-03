---
page_title: "xcsh_network_interface reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_interface reference."
---

# xcsh_network_interface reference

<a id="canonical-1021212310123203-1103103233320023-1232033231212231-1323032030311131-2312131013102210-3031303101112101-3303300230201333-0133202213101333"></a>

## Direct properties — static_ip / 001303133230 / 3

- [cluster_static_ip](data-sources--network_interface--reference--group-002.md#canonical-0133120302310230-3112011022020110-2212301300100010-3200021301000220-0202020000132300-0023233133013112-2103031120130301-2233202332102013): complete subsection reference.

- [node_static_ip](data-sources--network_interface--reference--group-002.md#canonical-3230031122333310-1130013023032301-3303011331212201-2231332320203231-3323013210221212-3302000020222320-0210111121010031-2312311202031102): complete subsection reference.

<a id="canonical-0110323120031230-2330213111111320-3302231231202103-0231102030012332-2213223313330120-0213220301221032-1011101000012230-0223113220333300"></a>

## Next pages — static_ip / 001303133230 / 4

- [ethernet_interface.static_ip.cluster_static_ip](data-sources--network_interface--reference--group-002.md#canonical-0133120302310230-3112011022020110-2212301300100010-3200021301000220-0202020000132300-0023233133013112-2103031120130301-2233202332102013)
- [ethernet_interface.static_ip.node_static_ip](data-sources--network_interface--reference--group-002.md#canonical-3230031122333310-1130013023032301-3303011331212201-2231332320203231-3323013210221212-3302000020222320-0210111121010031-2312311202031102)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-0133120302310230-3112011022020110-2212301300100010-3200021301000220-0202020000132300-0023233133013112-2103031120130301-2233202332102013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3023220223012300-0223330010333123-2213221001321310-2321031301302033-0330122032223122-1112333130103130-3121011013112201-3111001220211000"></a>

## ethernet_interface.static_ip.cluster_static_ip — cluster_static_ip / 301030221222 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.static_ip](data-sources--network_interface--reference--group-001.md#canonical-3312221322021303-3120113113303322-0130201201103221-3232302303112021-3223032010200211-1122330011012312-2200002231303102-1310333113123231)
- ethernet_interface.static_ip.cluster_static_ip

<a id="canonical-1031130201223312-1002010011231021-1023131120012231-3312313020123003-1013130302310212-2111103023022103-3302133032320000-1212322300221320"></a>

Type: `"single"`. Computed.

Configure Static IP parameters for cluster.

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

<a id="canonical-1323333112301320-0330320332331200-2020110101102013-2130311332210310-0120120203313120-3212300320303310-3002133330101312-1223013120112202"></a>

## Direct properties — cluster_static_ip / 301030221222 / 3

<a id="canonical-3010023201022000-0002103200211213-1131101320013211-0230300231230300-1111111333312033-0002202002313212-1232103100112123-0320213033110323"></a>

<a id="canonical-3121323320000311-0233023002321130-3130033032000022-2203211200201102-0011132303310131-1103122323203220-3232211101320000-2123330312330023"></a>

## interface_ip_map property — cluster_static_ip / 301030221222 / 4

Type: `["map", "string"]`. Computed.

Map of Node to Static IP configuration value, Key:Node, Value:IP Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "object",
    "maxProperties": 128,
    "metadata": {
      "confidence": 0.75,
      "source": "inferred",
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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  }
}
```

<a id="canonical-0112220330133033-3320320323020033-1232203233003033-0000233023032203-1102122200100133-2011010121000013-2320121133000311-1310302021223310"></a>

## Next pages — cluster_static_ip / 301030221222 / 5

- [ethernet_interface.static_ip](data-sources--network_interface--reference--group-001.md#canonical-3312221322021303-3120113113303322-0130201201103221-3232302303112021-3223032010200211-1122330011012312-2200002231303102-1310333113123231)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-3230031122333310-1130013023032301-3303011331212201-2231332320203231-3323013210221212-3302000020222320-0210111121010031-2312311202031102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203130001133021-0002120131333200-2111101013113010-1013212323210032-2321201221102213-2130133203020311-3212312003231332-1320100133310133"></a>

## ethernet_interface.static_ip.node_static_ip — node_static_ip / 323030320023 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.static_ip](data-sources--network_interface--reference--group-001.md#canonical-3312221322021303-3120113113303322-0130201201103221-3232302303112021-3223032010200211-1122330011012312-2200002231303102-1310333113123231)
- ethernet_interface.static_ip.node_static_ip

<a id="canonical-3220110030331230-1032002202032023-2333010201201130-3203233103311303-1132321012031032-0130311200010112-3311001102201303-0202012112021300"></a>

Type: `"single"`. Computed.

Configure Static IP parameters for a node.

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

<a id="canonical-0112323311130033-0012031313203133-1210333131130211-2332111320202021-3111210330032300-3222333011221322-2310021210301221-0200120221321023"></a>

## Direct properties — node_static_ip / 323030320023 / 3

<a id="canonical-3110022012211100-1000330221221332-2010203022002301-1112121212213000-1132203221022123-3303102022113022-0330020230001310-3230312302113303"></a>

<a id="canonical-3302220010003331-1011220101320022-0012031101300200-2331302321201032-3210320023103010-1231010332022331-2333131121200111-0011200033010111"></a>

## default_gw property — node_static_ip / 323030320023 / 4

Type: `"string"`. Computed.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-1130111221303111-0002310333300023-1331131011020111-3233233323100022-3310131131300331-2030130313023332-3320222132110323-3110030012310013"></a>

<a id="canonical-0012101210321003-2030121133220001-2203110221030203-3202232233012131-2213122020212101-1100300022022312-0021303320203133-3313123111133231"></a>

## dns_server property — node_static_ip / 323030320023 / 5

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-2103113210201222-0011033303212021-2102120302002230-2131021130330311-2302320233322121-2203212323322031-1203033120323103-2300131000300201"></a>

<a id="canonical-3322023311113301-2200022132113221-2002331313001220-3120202233101203-2010301323331321-3011121020211301-0023000133022001-0233003031200322"></a>

## ip_address property — node_static_ip / 323030320023 / 6

Type: `"string"`. Computed.

IP address of the interface and prefix length.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

<a id="canonical-2032230320311033-3223103303011232-1011230322020302-3213221033010201-1021020101120303-1232312010211131-1102302220123001-2300202301012332"></a>

## Next pages — node_static_ip / 323030320023 / 7

- [ethernet_interface.static_ip](data-sources--network_interface--reference--group-001.md#canonical-3312221322021303-3120113113303322-0130201201103221-3232302303112021-3223032010200211-1122330011012312-2200002231303102-1310333113123231)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-2122131222310110-1332200023023011-0111031123111230-3110000330210323-2132201123233021-2222111300011033-3001000030212323-0221213333013310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001220331203023-3101220111311011-2021101121331213-3102113100302030-2221030300112032-2323133001220012-2302131113021312-0201021221311332"></a>

## ethernet_interface.static_ipv6_address — static_ipv6_address / 012330001123 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- ethernet_interface.static_ipv6_address

<a id="canonical-2012012100233103-0020302021020220-1021110301020301-0211320203011110-3033201122200101-2033021031201011-3311223230020021-2323030221003131"></a>

Type: `"single"`. Computed.

Static IP Parameters. Configure Static IP parameters.

Upstream description:

Configure Static IP parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_prefix_choice": "[\"cluster_static_ip\",\"node_static_ip\"]"
}
```

<a id="canonical-3320231033203230-1210033203003033-3200210223130310-0232121122032003-0321103311232202-3102213123212221-1113111210002100-0213303212000130"></a>

## Direct properties — static_ipv6_address / 012330001123 / 3

- [cluster_static_ip](data-sources--network_interface--reference--group-002.md#canonical-2000333032022001-3322313322223203-1013133100320320-3033310202132103-2203113332121223-2322020003032112-1131120000310220-2132120112013312): complete subsection reference.

- [node_static_ip](data-sources--network_interface--reference--group-002.md#canonical-2302202323030311-3302100320112003-2033232230321102-1101111301001110-1220330230220000-3231210201001121-0323103312331311-0111020120231012): complete subsection reference.

<a id="canonical-1302010203200132-1220112002111121-2132103111301132-2212131101323303-0001010133301332-3310131010010302-2312013233322021-2220210232100130"></a>

## Next pages — static_ipv6_address / 012330001123 / 4

- [ethernet_interface.static_ipv6_address.cluster_static_ip](data-sources--network_interface--reference--group-002.md#canonical-2000333032022001-3322313322223203-1013133100320320-3033310202132103-2203113332121223-2322020003032112-1131120000310220-2132120112013312)
- [ethernet_interface.static_ipv6_address.node_static_ip](data-sources--network_interface--reference--group-002.md#canonical-2302202323030311-3302100320112003-2033232230321102-1101111301001110-1220330230220000-3231210201001121-0323103312331311-0111020120231012)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-2000333032022001-3322313322223203-1013133100320320-3033310202132103-2203113332121223-2322020003032112-1131120000310220-2132120112013312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031202203211103-1222122112320030-1100102210233210-0330313233020022-3011202230213002-2103200301302310-3100323032121000-1232332130102102"></a>

## ethernet_interface.static_ipv6_address.cluster_static_ip — cluster_static_ip / 222131223013 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.static_ipv6_address](data-sources--network_interface--reference--group-002.md#canonical-2122131222310110-1332200023023011-0111031123111230-3110000330210323-2132201123233021-2222111300011033-3001000030212323-0221213333013310)
- ethernet_interface.static_ipv6_address.cluster_static_ip

<a id="canonical-2131010032210103-0301310100010222-3331312032212000-3120323223031120-0001221020222220-0122032103310201-3113330222323322-2203011002101131"></a>

Type: `"single"`. Computed.

Configure Static IP parameters for cluster.

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

<a id="canonical-3030201023112331-2322100112030021-2121232023112202-1013132333213013-0130231033003303-0320011102120002-3331330311111231-2230202201213202"></a>

## Direct properties — cluster_static_ip / 222131223013 / 3

<a id="canonical-2211030020312330-1133122213031302-0120020203213133-3223132310330312-0223122211320132-3130022030310231-3330103113121130-0203223322321220"></a>

<a id="canonical-1302301113303033-2113121333211203-1030111022212032-0203001330032000-1030231010200301-1011331222331120-2111200113301122-2003011220231133"></a>

## interface_ip_map property — cluster_static_ip / 222131223013 / 4

Type: `["map", "string"]`. Computed.

Map of Node to Static IP configuration value, Key:Node, Value:IP Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "object",
    "maxProperties": 128,
    "metadata": {
      "confidence": 0.75,
      "source": "inferred",
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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  }
}
```

<a id="canonical-3000011211230033-3103103101210002-1301322230233021-0231230023120201-3133320203300033-3031110003212033-3101333113213301-3111101301223220"></a>

## Next pages — cluster_static_ip / 222131223013 / 5

- [ethernet_interface.static_ipv6_address](data-sources--network_interface--reference--group-002.md#canonical-2122131222310110-1332200023023011-0111031123111230-3110000330210323-2132201123233021-2222111300011033-3001000030212323-0221213333013310)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-2302202323030311-3302100320112003-2033232230321102-1101111301001110-1220330230220000-3231210201001121-0323103312331311-0111020120231012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0322101110121222-0221023200221211-0013010210121332-3320200030311311-1333310300001233-0333000123213032-3312001203310012-0122211133303010"></a>

## ethernet_interface.static_ipv6_address.node_static_ip — node_static_ip / 123220310121 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.static_ipv6_address](data-sources--network_interface--reference--group-002.md#canonical-2122131222310110-1332200023023011-0111031123111230-3110000330210323-2132201123233021-2222111300011033-3001000030212323-0221213333013310)
- ethernet_interface.static_ipv6_address.node_static_ip

<a id="canonical-1032002110320220-0320210300113000-2123213232011231-3102030233010020-2100200122333303-0212032303210331-1212023010232210-0111211101203020"></a>

Type: `"single"`. Computed.

Configure Static IP parameters for a node.

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

<a id="canonical-2201201331121231-2103012101301010-0223132333020223-1211213320123110-1123131123333210-0213022203102022-1133232233111011-1232223323112131"></a>

## Direct properties — node_static_ip / 123220310121 / 3

<a id="canonical-2000211102002202-2121133233233201-2131201331212102-1120013231111300-2332320211210120-2320130111311010-2102313123202221-3202002211323302"></a>

<a id="canonical-1012012133100321-1312001233000021-1103011003222021-3030110322003023-0320330102332013-0121313100031112-2210333213333332-2220310130020122"></a>

## default_gw property — node_static_ip / 123220310121 / 4

Type: `"string"`. Computed.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-2003223203303030-0220312301313001-1310212102112201-0302012020010100-3100231221010200-2312133032331302-3133130011321220-3213220212301133"></a>

<a id="canonical-0221323321012221-3123301031133202-2221100230220021-2132210303101320-0012333230222300-0133010223130012-1322102010032130-0221133301021013"></a>

## dns_server property — node_static_ip / 123220310121 / 5

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-0203013002223113-0201100122031301-0120300221012000-0202031322203110-0302200300032231-2003330001110313-3113332222230321-2321221020013201"></a>

<a id="canonical-0311330212000022-3301311010321220-2203013202221332-2112102211321200-1121301021122323-2110213122311003-2122202201023201-0223332000001023"></a>

## ip_address property — node_static_ip / 123220310121 / 6

Type: `"string"`. Computed.

IP address of the interface and prefix length.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

<a id="canonical-2100221300132113-0032102103101230-1312302200103100-1201302012300322-0313133121213022-1220313313200221-2321230120233300-0203300013311003"></a>

## Next pages — node_static_ip / 123220310121 / 7

- [ethernet_interface.static_ipv6_address](data-sources--network_interface--reference--group-002.md#canonical-2122131222310110-1332200023023011-0111031123111230-3110000330210323-2132201123233021-2222111300011033-3001000030212323-0221213333013310)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-0333033210010321-2132211112223013-1230003210301302-3330133302221210-0113332030013030-1331320233211220-3001013200001121-2203312010010122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001121233331102-2021032023000023-1303030032003002-0302023313321321-3213210023222333-1211022121302001-0113323013320002-3301000302330332"></a>

## ethernet_interface.storage_network — storage_network / 202210323121 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- ethernet_interface.storage_network

<a id="canonical-2222221302222022-1230033212101101-1222323331022000-0130003210020222-0102133221030111-3331300022113330-0010133020123323-3310123231323310"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for storage network.

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

<a id="canonical-0032112230033122-1302330013002233-0133320210120301-0113312233130233-1000103213311012-0033013322122231-1302100223232123-0331333212331003"></a>

## Direct properties — storage_network / 202210323121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0212322322321123-3330013100012213-3223100232010300-3220020100110322-1322221333310122-3001232233131133-2212220332023133-0303233003112112"></a>

## Next pages — storage_network / 202210323121 / 4

- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-2303100002223300-2211310120020011-3133100223301321-2323322300103112-1230330031231133-3230103021330133-1023322211311012-0001221012231312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131321330111200-3202120333032002-3221330000103112-1131111031021303-0202110011322222-0030301221000003-1211212032213301-1332103232032232"></a>

## ethernet_interface.untagged — untagged / 222120121303 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- ethernet_interface.untagged

<a id="canonical-3332202232033323-3023222020030011-3230200311123331-3021012122032332-0131010220030001-1332031123233013-3120333302302030-1101110130010102"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1011021321311322-2203232201003312-3221211223011210-1000021333230003-1221010002202013-2220032030131330-1101013023302103-1130312311022120"></a>

## Direct properties — untagged / 222120121303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2023103223023223-1032330130221121-2022001133220033-2201321120231330-2331332000323201-1023023121222022-0022112032032011-2200200223300332"></a>

## Next pages — untagged / 222120121303 / 4

- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-0203332310121320-0320201223332130-3003222223202313-0112031223031223-1020213311213033-3313000032021132-3232210333100223-0230211123130030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123323132031212-0313033111031330-2112010121121202-0023313012211111-2110321122102210-0112133131103213-3322121111132021-1003223121323132"></a>

## layer2_interface — layer2_interface / 030302302221 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- layer2_interface

<a id="canonical-3202113010230203-3202320132300102-1102221210032110-3313003010222230-0232130003102200-0100010011021322-1003230323121233-0033100200303223"></a>

Type: `"single"`. Computed.

Configuration parameter for layer2 interface.

Upstream description:

Layer2 Interface Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-layer2_interface_choice": "[\"l2sriov_interface\",\"l2vlan_interface\",\"l2vlan_slo_interface\"]"
}
```

<a id="canonical-3023203021301023-1202130301213311-2321312222101121-1231112032021303-2032132132332100-3223020013000130-0120011332033332-1032300110331003"></a>

## Direct properties — layer2_interface / 030302302221 / 3

- [l2sriov_interface](data-sources--network_interface--reference--group-002.md#canonical-2321233030331021-2213203223212012-2300223322311023-3130133333303222-0211302321321120-2232321231200220-2302011030310300-3221333030301021): complete subsection reference.

- [l2vlan_interface](data-sources--network_interface--reference--group-002.md#canonical-0330330120020103-1222300012103030-1232232331201001-2303012223312000-1223302103300000-1221113313122123-1013010302300132-0102113300031232): complete subsection reference.

- [l2vlan_slo_interface](data-sources--network_interface--reference--group-002.md#canonical-1310020032101313-3031201121133011-1230010330213320-2222133123313220-0222120030331301-1221220002202101-1300220222313213-3101332120332022): complete subsection reference.

<a id="canonical-3220331012331113-0010102232330001-1123330323020322-3010102010222022-3200122223120220-0330021312212021-1232221330202130-2033313302130323"></a>

## Next pages — layer2_interface / 030302302221 / 4

- [layer2_interface.l2sriov_interface](data-sources--network_interface--reference--group-002.md#canonical-2321233030331021-2213203223212012-2300223322311023-3130133333303222-0211302321321120-2232321231200220-2302011030310300-3221333030301021)
- [layer2_interface.l2vlan_interface](data-sources--network_interface--reference--group-002.md#canonical-0330330120020103-1222300012103030-1232232331201001-2303012223312000-1223302103300000-1221113313122123-1013010302300132-0102113300031232)
- [layer2_interface.l2vlan_slo_interface](data-sources--network_interface--reference--group-002.md#canonical-1310020032101313-3031201121133011-1230010330213320-2222133123313220-0222120030331301-1221220002202101-1300220222313213-3101332120332022)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-2321233030331021-2213203223212012-2300223322311023-3130133333303222-0211302321321120-2232321231200220-2302011030310300-3221333030301021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1132122011203323-0012120102021113-1222133010230021-2232110322321232-0003023013020323-3203122301202303-0231121202231321-3003320322122133"></a>

## layer2_interface.l2sriov_interface — l2sriov_interface / 131310201021 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [layer2_interface](data-sources--network_interface--reference--group-002.md#canonical-0203332310121320-0320201223332130-3003222223202313-0112031223031223-1020213311213033-3313000032021132-3232210333100223-0230211123130030)
- layer2_interface.l2sriov_interface

<a id="canonical-3132200021010003-1031232001333112-3313022332200113-1303021311010002-2023210112130201-2020211011321002-2333213110330102-2203023111202012"></a>

Type: `"single"`. Computed.

Configuration parameter for l2sriov interface.

Upstream description:

Layer2 SR-IOV Interface Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-vlan_choice": "[\"untagged\",\"vlan_id\"]"
}
```

<a id="canonical-1210123232200330-1032010301312332-3032203032201032-2010110230313121-3300010020013031-3000323202103100-2001331311312230-0002313031111102"></a>

## Direct properties — l2sriov_interface / 131310201021 / 3

<a id="canonical-1112033311201310-1131020021213012-0100012132231000-0130010112313332-0031102113231231-3321011302331130-3330230122113322-3320322303010203"></a>

<a id="canonical-1033222022021210-3103133122133311-3030301301231323-2330112123011021-3210022110321120-1223132123000100-3301013202101313-1002030330323023"></a>

## device property — l2sriov_interface / 131310201021 / 4

Type: `"string"`. Computed.

Ethernet Device. Physical ethernet interface.

Upstream description:

Physical ethernet interface.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [untagged](data-sources--network_interface--reference--group-002.md#canonical-2103020311113321-3031302111201022-1221222221011223-2111222311100331-0012102101211202-1212301213000001-3301323103311010-0210321110232213): complete subsection reference.

<a id="canonical-1330333313200301-1210020213320221-2332100321130123-3101002311233003-2121222133120011-2221121223301102-1323010113313302-0213333111232100"></a>

<a id="canonical-2003110012212121-1301132131021132-0310031231011301-0232230220001012-3030000012203111-0002101222022313-1221311323203201-0301232330102011"></a>

## vlan_id property — l2sriov_interface / 131310201021 / 5

Type: `"number"`. Computed.

Exclusive with \[untagged\] Configure a VLAN tagged interface.

Upstream description:

Exclusive with \[untagged\] Configure a VLAN tagged interface.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 4095,
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
    "ves.io.schema.rules.uint32.lte": "4095"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  }
}
```

<a id="canonical-0100321322220130-3210033212310323-0012122312120310-2232333023302232-0212330221221121-0232112333012222-0333231233113020-3302000211032200"></a>

## Next pages — l2sriov_interface / 131310201021 / 6

- [layer2_interface.l2sriov_interface.untagged](data-sources--network_interface--reference--group-002.md#canonical-2103020311113321-3031302111201022-1221222221011223-2111222311100331-0012102101211202-1212301213000001-3301323103311010-0210321110232213)
- [layer2_interface](data-sources--network_interface--reference--group-002.md#canonical-0203332310121320-0320201223332130-3003222223202313-0112031223031223-1020213311213033-3313000032021132-3232210333100223-0230211123130030)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-2103020311113321-3031302111201022-1221222221011223-2111222311100331-0012102101211202-1212301213000001-3301323103311010-0210321110232213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011300233121133-3012323313001103-0113131321022131-2310103031331322-0010332131232300-3122010002310213-1222213012311322-0110002330011222"></a>

## layer2_interface.l2sriov_interface.untagged — untagged / 213021113211 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [layer2_interface](data-sources--network_interface--reference--group-002.md#canonical-0203332310121320-0320201223332130-3003222223202313-0112031223031223-1020213311213033-3313000032021132-3232210333100223-0230211123130030)
- [layer2_interface.l2sriov_interface](data-sources--network_interface--reference--group-002.md#canonical-2321233030331021-2213203223212012-2300223322311023-3130133333303222-0211302321321120-2232321231200220-2302011030310300-3221333030301021)
- layer2_interface.l2sriov_interface.untagged

<a id="canonical-2011230012310221-1233300113013111-1133230331120100-2233131231222101-0121332213122112-3213032120301030-1102122303112302-0002001231033222"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3333120323330210-2010210002322130-3003300300210332-1223130223112213-3231231032132221-0312103100331313-2012233120331013-1201310311202312"></a>

## Direct properties — untagged / 213021113211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2110023310100322-2130021002030202-1113222220321131-0222031233210223-0321030231020112-3112333231023211-3131101133102323-1233202032020201"></a>

## Next pages — untagged / 213021113211 / 4

- [layer2_interface.l2sriov_interface](data-sources--network_interface--reference--group-002.md#canonical-2321233030331021-2213203223212012-2300223322311023-3130133333303222-0211302321321120-2232321231200220-2302011030310300-3221333030301021)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-0330330120020103-1222300012103030-1232232331201001-2303012223312000-1223302103300000-1221113313122123-1013010302300132-0102113300031232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3102333122330213-2333233332200000-3233232000231103-3001211330120203-3300123001111220-3122202121022312-0112132320232320-3300001103110203"></a>

## layer2_interface.l2vlan_interface — l2vlan_interface / 201332310000 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [layer2_interface](data-sources--network_interface--reference--group-002.md#canonical-0203332310121320-0320201223332130-3003222223202313-0112031223031223-1020213311213033-3313000032021132-3232210333100223-0230211123130030)
- layer2_interface.l2vlan_interface

<a id="canonical-0232122010331133-3302132312333320-1102001130210213-1032031132220333-3011020220212311-3331102000223133-1021122021012133-1332033031100210"></a>

Type: `"single"`. Computed.

Configuration parameter for l2vlan interface.

Upstream description:

Layer2 VLAN Interface Configuration.

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

<a id="canonical-0121200122133323-3111301101011021-3223113301301211-0022201332023220-3332232200113311-0323012301102100-0103223020303102-1013102003030021"></a>

## Direct properties — l2vlan_interface / 201332310000 / 3

<a id="canonical-0113023300301001-0130030320211111-0003311210221212-3211123101010302-2210122010030313-3203111331033322-3002200201331222-1230123323211000"></a>

<a id="canonical-3120023121203202-2001010120231321-0103222300303121-0233002331112230-1220203220213021-1332011011233112-1310210131011002-3231203221303132"></a>

## device property — l2vlan_interface / 201332310000 / 4

Type: `"string"`. Computed.

Ethernet Device. Physical ethernet interface.

Upstream description:

Physical ethernet interface.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0022022123232202-2000212030132302-0200123333331320-3212123321320020-0230211320333102-2321210321033112-1332212100230022-3101123223010331"></a>

<a id="canonical-1022303233202201-1301132313020022-2333200221233233-0301323112100131-1322012121203002-3001021213232030-0020222211022133-2300000031230111"></a>

## vlan_id property — l2vlan_interface / 201332310000 / 5

Type: `"number"`. Computed.

VLAN ID. VLAN ID

Upstream description:

VLAN ID

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 4095,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  }
}
```

<a id="canonical-3102303011300323-1333320213033220-3132221323132000-1010322021030330-0101120323300023-0231233331110321-0020123220130103-2011003111303310"></a>

## Next pages — l2vlan_interface / 201332310000 / 6

- [layer2_interface](data-sources--network_interface--reference--group-002.md#canonical-0203332310121320-0320201223332130-3003222223202313-0112031223031223-1020213311213033-3313000032021132-3232210333100223-0230211123130030)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-1310020032101313-3031201121133011-1230010330213320-2222133123313220-0222120030331301-1221220002202101-1300220222313213-3101332120332022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0222231203300001-1212103332023002-2103200210231310-1101330332323300-1312300113030121-2021122122113303-2011333100103111-1201323133110031"></a>

## layer2_interface.l2vlan_slo_interface — l2vlan_slo_interface / 330323210310 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [layer2_interface](data-sources--network_interface--reference--group-002.md#canonical-0203332310121320-0320201223332130-3003222223202313-0112031223031223-1020213311213033-3313000032021132-3232210333100223-0230211123130030)
- layer2_interface.l2vlan_slo_interface

<a id="canonical-2213122321301030-2303330210333230-3330100000120121-2302332212012100-0303020220331233-2032310200101322-3000320320133331-0303110002021132"></a>

Type: `"single"`. Computed.

Layer2 Site Local Outside VLAN Interface Configuration.

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

<a id="canonical-3222021302003033-2332203200203200-1103132323320330-1211332002333212-0210120211311321-3011030331310111-0323133101301120-3131102033131001"></a>

## Direct properties — l2vlan_slo_interface / 330323210310 / 3

<a id="canonical-2113113000333321-2320012032112303-3133231003003030-0331223223212110-1022012222203300-2220123202300201-0222320222200022-0332302011311111"></a>

<a id="canonical-2321330121122211-1221232033013320-0232033131132230-0212132020003212-1132301010032112-0101030231330003-2123230333321112-1213300110232020"></a>

## vlan_id property — l2vlan_slo_interface / 330323210310 / 4

Type: `"number"`. Computed.

VLAN ID. VLAN ID

Upstream description:

VLAN ID

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 4095,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  }
}
```

<a id="canonical-0233332120023203-3323331303012133-3333122311222231-3222130313322311-1030111131320111-2111011330221102-1233010010301231-0300110200321112"></a>

## Next pages — l2vlan_slo_interface / 330323210310 / 5

- [layer2_interface](data-sources--network_interface--reference--group-002.md#canonical-0203332310121320-0320201223332130-3003222223202313-0112031223031223-1020213311213033-3313000032021132-3232210333100223-0230211123130030)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-1030102220310031-2101220021010303-1132023020221330-1013302003330133-2331301123030223-3233302202221123-3320020030202302-1320303030232313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3301313023312122-3111110023101331-0210212021021002-3311231200132130-0121030320000033-0331001303133002-2003303101210330-0333111223213213"></a>

## tunnel_interface — tunnel_interface / 233321002102 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- tunnel_interface

<a id="canonical-0202122131102011-3232213112222323-2223112333132222-1101010011000212-3201033332312313-1112311332113213-3310013212203003-0200303300303233"></a>

Type: `"single"`. Computed.

Configuration parameter for tunnel interface.

Upstream description:

Tunnel Interface Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"site_local_inside_network\",\"site_local_network\"]",
  "x-ves-oneof-field-node_choice": "[\"node\"]"
}
```

<a id="canonical-2002203202222023-1120211022001032-2113303311121231-2303030213030031-3120213110122223-0233211332103210-3120102033112331-2123202123310131"></a>

## Direct properties — tunnel_interface / 233321002102 / 3

<a id="canonical-2301122233103120-3112112022331222-0231123032211032-3110312111020213-2203203103133233-0200001103313123-2101303013211020-2000122113310232"></a>

<a id="canonical-2223312000122111-1003311331023011-0302221312223011-3321310320103012-0223213031031022-1201011303010131-2123111110233311-1033012211033113"></a>

## mtu property — tunnel_interface / 233321002102 / 4

Type: `"number"`. Computed.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 9000,
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
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  }
}
```

<a id="canonical-2233332222021020-0212203121222212-1201333202121302-3322330033213101-2223000322030020-2113221132000223-2321300020302312-1312133313133201"></a>

<a id="canonical-3220001211330030-2132123130203320-2031203231133220-3033211223302111-0331102212130022-1122122303212121-1312322121021133-2002333203301131"></a>

## node property — tunnel_interface / 233321002102 / 5

Type: `"string"`. Computed.

Exclusive with \[\] Configuration will apply to a given device on the given node.

Upstream description:

Exclusive with \[\] Configuration will apply to a given device on the given node.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2122101203312001-2323121303032301-1322312322302221-3322311331232333-0230320212003212-0201120000322321-2120111221233331-0313003222333233"></a>

<a id="canonical-1023230130133113-1313210102203130-3321333100200121-0222100101323303-3321332103233201-1030113320322310-2213013301100021-0112122120203112"></a>

## priority property — tunnel_interface / 233321002102 / 6

Type: `"number"`. Computed.

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Upstream description:

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

- [site_local_inside_network](data-sources--network_interface--reference--group-002.md#canonical-3200332131130022-3120200023013333-2002211203321130-1322310032233303-1003303131302120-0101303201220213-2222001110031020-2221303132033100): complete subsection reference.

- [site_local_network](data-sources--network_interface--reference--group-002.md#canonical-2221222222220313-0103032102101022-3232123232013030-1030202202321210-1300333113001022-0000302010201210-1310302133012213-1002100222303130): complete subsection reference.

- [static_ip](data-sources--network_interface--reference--group-002.md#canonical-3121212033100033-0132200113233111-1130113303120101-0331313032310301-2312200021332010-0102333211323101-0333123111113031-2303203102211032): complete subsection reference.

- [tunnel](data-sources--network_interface--reference--group-002.md#canonical-0102301312001330-2110301120232212-0101133301010002-3123313010311003-0221110020121013-1330123002102212-1022122002001323-1202110330310231): complete subsection reference.

<a id="canonical-2030212232130023-1103200203002322-1323220011002321-2113133210020002-0200310332231220-0212211102303323-1330032303032313-3211023113222202"></a>

## Next pages — tunnel_interface / 233321002102 / 7

- [tunnel_interface.site_local_inside_network](data-sources--network_interface--reference--group-002.md#canonical-3200332131130022-3120200023013333-2002211203321130-1322310032233303-1003303131302120-0101303201220213-2222001110031020-2221303132033100)
- [tunnel_interface.site_local_network](data-sources--network_interface--reference--group-002.md#canonical-2221222222220313-0103032102101022-3232123232013030-1030202202321210-1300333113001022-0000302010201210-1310302133012213-1002100222303130)
- [tunnel_interface.static_ip](data-sources--network_interface--reference--group-002.md#canonical-3121212033100033-0132200113233111-1130113303120101-0331313032310301-2312200021332010-0102333211323101-0333123111113031-2303203102211032)
- [tunnel_interface.tunnel](data-sources--network_interface--reference--group-002.md#canonical-0102301312001330-2110301120232212-0101133301010002-3123313010311003-0221110020121013-1330123002102212-1022122002001323-1202110330310231)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-3200332131130022-3120200023013333-2002211203321130-1322310032233303-1003303131302120-0101303201220213-2222001110031020-2221303132033100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1230123113001223-3100230221120010-1220312130131020-1213313020110301-3112003320011203-2121332311200230-0222322221311033-0130123312033131"></a>

## tunnel_interface.site_local_inside_network — site_local_inside_network / 230103130122 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [tunnel_interface](data-sources--network_interface--reference--group-002.md#canonical-1030102220310031-2101220021010303-1132023020221330-1013302003330133-2331301123030223-3233302202221123-3320020030202302-1320303030232313)
- tunnel_interface.site_local_inside_network

<a id="canonical-3131333032013223-1021201131223103-3121100210012223-3003312331101032-3123131112323102-3222200300101301-2233222321233032-0223102001110123"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2130113233000331-2130010100213211-1321101201220320-3013013201010223-0122111311301223-0221311102120012-1332111303023103-3132011110030123"></a>

## Direct properties — site_local_inside_network / 230103130122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1130020002030110-1021123331122132-3211221201221333-0221113332012012-1203200130013203-3002132130011232-3133022200311112-2020103211200020"></a>

## Next pages — site_local_inside_network / 230103130122 / 4

- [tunnel_interface](data-sources--network_interface--reference--group-002.md#canonical-1030102220310031-2101220021010303-1132023020221330-1013302003330133-2331301123030223-3233302202221123-3320020030202302-1320303030232313)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-2221222222220313-0103032102101022-3232123232013030-1030202202321210-1300333113001022-0000302010201210-1310302133012213-1002100222303130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320313201310211-1002111313313212-1231011021002221-1302233003100300-1313230010231332-3123101112222110-3220221200313002-1232203202133102"></a>

## tunnel_interface.site_local_network — site_local_network / 112202323203 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [tunnel_interface](data-sources--network_interface--reference--group-002.md#canonical-1030102220310031-2101220021010303-1132023020221330-1013302003330133-2331301123030223-3233302202221123-3320020030202302-1320303030232313)
- tunnel_interface.site_local_network

<a id="canonical-1000312200201102-3330320200321211-2212132111202100-3030330320110133-1111232132221101-0133103232233200-0112330201110013-1031001312110121"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0213133121201010-0203220112331101-2100313310331121-0100222110333110-0001103201323110-1313111131232213-3012133321222011-2301033130300002"></a>

## Direct properties — site_local_network / 112202323203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2300102323011002-0110301222000101-2213200121230230-1232101230103100-2333202001231212-1111023020331013-0313030303111131-0230121230130002"></a>

## Next pages — site_local_network / 112202323203 / 4

- [tunnel_interface](data-sources--network_interface--reference--group-002.md#canonical-1030102220310031-2101220021010303-1132023020221330-1013302003330133-2331301123030223-3233302202221123-3320020030202302-1320303030232313)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-3121212033100033-0132200113233111-1130113303120101-0331313032310301-2312200021332010-0102333211323101-0333123111113031-2303203102211032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0322122112122131-0332122312231210-3221013322031003-3213001000132020-2102002310003110-0021313213122010-2320212010020032-2210222122123310"></a>

## tunnel_interface.static_ip — static_ip / 202122032212 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [tunnel_interface](data-sources--network_interface--reference--group-002.md#canonical-1030102220310031-2101220021010303-1132023020221330-1013302003330133-2331301123030223-3233302202221123-3320020030202302-1320303030232313)
- tunnel_interface.static_ip

<a id="canonical-3113112132103132-3203032102300021-0103202132212110-3322112313010211-0122300231201223-3310202132113311-2301231320023232-2332012333313130"></a>

Type: `"single"`. Computed.

Static IP Parameters. Configure Static IP parameters.

Upstream description:

Configure Static IP parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_prefix_choice": "[\"cluster_static_ip\",\"node_static_ip\"]"
}
```

<a id="canonical-0212100131132301-0121211321131212-1211123120313233-1322022220111120-2131202212111030-1300322123021230-3120323323120222-0321311031003212"></a>

## Direct properties — static_ip / 202122032212 / 3

- [cluster_static_ip](data-sources--network_interface--reference--group-002.md#canonical-0000300202223022-3302100312022303-0202020302222110-3320022230230331-3210013203103021-1231321223111312-1003233220110010-3332031202310001): complete subsection reference.

- [node_static_ip](data-sources--network_interface--reference--group-002.md#canonical-1113231101321112-3303000123310131-2000230003120133-1313010210333012-2131222233000220-3000230022123310-1032020033112020-0223203121322031): complete subsection reference.

<a id="canonical-3030333003212111-2110322213212201-1323311132100222-0332023203120231-2323033133212313-3030200231120111-2320013331220100-2231113022211322"></a>

## Next pages — static_ip / 202122032212 / 4

- [tunnel_interface.static_ip.cluster_static_ip](data-sources--network_interface--reference--group-002.md#canonical-0000300202223022-3302100312022303-0202020302222110-3320022230230331-3210013203103021-1231321223111312-1003233220110010-3332031202310001)
- [tunnel_interface.static_ip.node_static_ip](data-sources--network_interface--reference--group-002.md#canonical-1113231101321112-3303000123310131-2000230003120133-1313010210333012-2131222233000220-3000230022123310-1032020033112020-0223203121322031)
- [tunnel_interface](data-sources--network_interface--reference--group-002.md#canonical-1030102220310031-2101220021010303-1132023020221330-1013302003330133-2331301123030223-3233302202221123-3320020030202302-1320303030232313)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-0000300202223022-3302100312022303-0202020302222110-3320022230230331-3210013203103021-1231321223111312-1003233220110010-3332031202310001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023200011202230-1101330111122313-0311330221312311-0012312131301311-0130230201023033-3023123110213311-3211312231002010-1331203021310132"></a>

## tunnel_interface.static_ip.cluster_static_ip — cluster_static_ip / 120033111221 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [tunnel_interface](data-sources--network_interface--reference--group-002.md#canonical-1030102220310031-2101220021010303-1132023020221330-1013302003330133-2331301123030223-3233302202221123-3320020030202302-1320303030232313)
- [tunnel_interface.static_ip](data-sources--network_interface--reference--group-002.md#canonical-3121212033100033-0132200113233111-1130113303120101-0331313032310301-2312200021332010-0102333211323101-0333123111113031-2303203102211032)
- tunnel_interface.static_ip.cluster_static_ip

<a id="canonical-0321110100210030-1113001033321023-0210230320101121-1321230122332123-3011201332031200-1012330232211212-0022111022033113-3021201120332210"></a>

Type: `"single"`. Computed.

Configure Static IP parameters for cluster.

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

<a id="canonical-1123231021203331-0012211031031232-0331223332002233-2303202012033003-0103212301313122-3120123103232011-0133201320202033-0110131301212102"></a>

## Direct properties — cluster_static_ip / 120033111221 / 3

<a id="canonical-3101130211022311-1023301213212323-2130213331001301-0033121110020201-3100330113310312-2123012113023310-0011222131002331-1120131102032133"></a>

<a id="canonical-2213310021100022-0110001023310030-2133003302032313-2313331120200030-2001133220010210-2232003022110332-2303112322021103-1103203201002032"></a>

## interface_ip_map property — cluster_static_ip / 120033111221 / 4

Type: `["map", "string"]`. Computed.

Map of Node to Static IP configuration value, Key:Node, Value:IP Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "object",
    "maxProperties": 128,
    "metadata": {
      "confidence": 0.75,
      "source": "inferred",
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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  }
}
```

<a id="canonical-3203121031032102-2121221321001133-2223021132020223-3332123330000033-0322213011232230-0122023303220113-1320310012212213-0221200023302112"></a>

## Next pages — cluster_static_ip / 120033111221 / 5

- [tunnel_interface.static_ip](data-sources--network_interface--reference--group-002.md#canonical-3121212033100033-0132200113233111-1130113303120101-0331313032310301-2312200021332010-0102333211323101-0333123111113031-2303203102211032)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-1113231101321112-3303000123310131-2000230003120133-1313010210333012-2131222233000220-3000230022123310-1032020033112020-0223203121322031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1332323133230320-1223122223213203-1130012021020322-2230111010000220-3013223013221132-3002321302000022-3120210303131131-1213222122021200"></a>

## tunnel_interface.static_ip.node_static_ip — node_static_ip / 303131101311 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [tunnel_interface](data-sources--network_interface--reference--group-002.md#canonical-1030102220310031-2101220021010303-1132023020221330-1013302003330133-2331301123030223-3233302202221123-3320020030202302-1320303030232313)
- [tunnel_interface.static_ip](data-sources--network_interface--reference--group-002.md#canonical-3121212033100033-0132200113233111-1130113303120101-0331313032310301-2312200021332010-0102333211323101-0333123111113031-2303203102211032)
- tunnel_interface.static_ip.node_static_ip

<a id="canonical-1311332322200012-2123002113110032-1012212100031111-2220101010130003-1001110031223110-1022310030023002-3103222230003121-2202101023300133"></a>

Type: `"single"`. Computed.

Configure Static IP parameters for a node.

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

<a id="canonical-1202212223233111-0303303102303302-0203220222021312-1312130210312013-0331032102200331-0120202221121210-1132131011002233-1331233222110303"></a>

## Direct properties — node_static_ip / 303131101311 / 3

<a id="canonical-0030132022001213-1330101023110212-3111001212032032-0312030032331330-3232323222212132-3122012331303112-3131301211130202-0021330312203021"></a>

<a id="canonical-2320201101012210-3010001030023321-1110323213002211-1033100321103331-3322313113003231-0210130011230203-2101000100131032-0133103232303333"></a>

## default_gw property — node_static_ip / 303131101311 / 4

Type: `"string"`. Computed.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-2303121121210330-1303230133303032-1230031333121233-3023220211332103-0310001121230032-3303030210323320-0303001222030121-2031313201031323"></a>

<a id="canonical-0020302023012112-2201313320130003-0221032103210103-0101203023111333-0030232302121202-3223112220222000-0211213122023122-1220113032011321"></a>

## dns_server property — node_static_ip / 303131101311 / 5

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-1223020222311233-0203332000321022-1122103323030102-2020211220110101-3213000310321132-3103311321100021-0111220100031330-3112033020100120"></a>

<a id="canonical-0101300313203231-1230200333101123-1030123111311111-0032220020323021-0210031030202303-0223202203010302-0203010212313233-3311001331300221"></a>

## ip_address property — node_static_ip / 303131101311 / 6

Type: `"string"`. Computed.

IP address of the interface and prefix length.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

<a id="canonical-3230201132012010-3012120310220331-3313032212103312-2121021223002311-3303000322010233-1331212333102200-2221202332100021-1123102232102122"></a>

## Next pages — node_static_ip / 303131101311 / 7

- [tunnel_interface.static_ip](data-sources--network_interface--reference--group-002.md#canonical-3121212033100033-0132200113233111-1130113303120101-0331313032310301-2312200021332010-0102333211323101-0333123111113031-2303203102211032)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-0102301312001330-2110301120232212-0101133301010002-3123313010311003-0221110020121013-1330123002102212-1022122002001323-1202110330310231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122331031330020-3301110221220300-0313020100000131-0213330120110321-3132013312203132-1100223301001211-1122220010012120-0010002310212113"></a>

## tunnel_interface.tunnel — tunnel / 301203220120 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [tunnel_interface](data-sources--network_interface--reference--group-002.md#canonical-1030102220310031-2101220021010303-1132023020221330-1013302003330133-2331301123030223-3233302202221123-3320020030202302-1320303030232313)
- tunnel_interface.tunnel

<a id="canonical-0313310202220321-3100221300010201-0322011131122303-1332000011200001-0312032200123323-0111102223233001-3011312020223322-0022010333122233"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-2110103201001012-1213201220232220-2201333331100010-3331320313010103-3220233330132302-2021010101321200-2302220011312030-1123210001303030"></a>

## Direct properties — tunnel / 301203220120 / 3

<a id="canonical-2231123023210121-3020002101303113-1032320312300323-0312120110111231-3213200002003223-2122202203203130-2033222030312013-2010033111001111"></a>

<a id="canonical-3300022100322312-1032102332131011-3312212013233002-2200212131321310-3213300113121102-2131130133302023-0020031300230013-0031121002003233"></a>

## name property — tunnel / 301203220120 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-2130220312101112-0223322103302131-1212020212132213-0031212331230103-1300202103021033-3113231311210031-2120332213101003-0033303310203330"></a>

<a id="canonical-1202020331220120-1203022232000310-2001320303031211-0212333023203332-2020331013133110-3203231311000210-1332323313003320-2012213303313031"></a>

## namespace property — tunnel / 301203220120 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-3011221000031031-2331131120023320-0002311032110300-3023121111001333-3120310103122023-3102131011323302-0032032203210313-2121012013321331"></a>

<a id="canonical-1101202230310313-1022323133022301-2333122203103203-0222101211001323-1201200132212321-0012200132132230-2211002030213230-1033320210212122"></a>

## tenant property — tunnel / 301203220120 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-2030301000202333-3021302101011212-2020330333211222-2123232131212211-2223113302210003-1311230103203001-0111213101101230-3223200333103202"></a>

## Next pages — tunnel / 301203220120 / 7

- [tunnel_interface](data-sources--network_interface--reference--group-002.md#canonical-1030102220310031-2101220021010303-1132023020221330-1013302003330133-2331301123030223-3233302202221123-3320020030202302-1320303030232313)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
