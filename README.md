# tileaddr
Universal Z-ordered address space for arbitrarily nested geospatial data.

## Layout

  +---------------+--------------------+
  | zoom (8 bits) | address (variable) |
  +---------------+--------------------+

## Ordering
Pixel addresses are Z-Ordered. For any bounding box of `tileaddr.TileSize`
(default 255), all sub-pixels at an arbitrary even resolution can be queried
with a single range query.

## Projection
The internal representation of all points will be
[EPSG:3857](http://spatialreference.org/ref/sr-org/7483/), or Web Mercator. The
reasoning is that summaries can be quickly computed for a [Slippy
map](http://wiki.openstreetmap.org/wiki/Slippy_Map) at arbitrary zoom levels and
pixel sizes.
