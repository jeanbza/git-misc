from PIL import Image
def check(palette, copy):
    palette = sorted(Image.open(palette).convert('RGB').getdata())
    copy = sorted(Image.open(copy).convert('RGB').getdata())
    print 'Images are compatible' if copy == palette else 'Images are not compatible'

check('src1.png', 'src2.png')