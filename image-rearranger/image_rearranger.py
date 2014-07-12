# Original idea from flawr at http://codegolf.stackexchange.com/questions/33172/american-gothic-in-the-palette-of-mona-lisa-rearrange-the-pixels

from PIL import Image

n = 5 #number of partitions per channel.

src_img = "src5.png"
dst_img = "src1.png"

def sortAndDivide(coordlist, pixelimage, channel):
    global src,dst,n
    retlist = []
    #sort
    coordlist.sort(key = lambda t: pixelimage[t][channel])
    #divide
    partitionLength = int(len(coordlist) / n)
    if partitionLength <= 0:
        partitionLength = 1
    if channel < 2:
        for i in range(0, len(coordlist),partitionLength):
            retlist += sortAndDivide(coordlist[i: i + partitionLength], pixelimage,channel + 1)
    else:
        retlist += coordlist
    return retlist

def makePixelList(img):
    l = []

    for x in range(img.size[0]):
        for y in range(img.size[1]):
            l.append((x,y))

    return l

def go():
    print "Starting!"

    src_handle = Image.open(src_img)
    dst_handle = Image.open(dst_img)

    src = src_handle.load()
    dst = dst_handle.load()

    assert src_handle.size[0] * src_handle.size[1] == dst_handle.size[0] * dst_handle.size[1], "images must be same size"

    lsrc = makePixelList(src_handle)
    ldst = makePixelList(dst_handle)

    lsrc = sortAndDivide(lsrc, src, 0)
    ldst = sortAndDivide(ldst, dst, 0)

    for i in range(len(ldst)):
        dst[ldst[i]] = src[lsrc[i]]

    dst_handle.save("exchange.png")

    print "Done!"

go()