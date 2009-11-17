package curcon
/*
<------------------------------------------Cursor Control------------------------------->
Originally written by Daniel Polnoff
Date 11/14/09
Description: Created as a wrapper for xterm escape codes for controlling the xterm cursor. 

Ansi escape functions refrenced here: http://www.xfree86.org/current/ctlseqs.html
<---------------------------------------------------------------------------------------->
*/
import . "fmt"
import "os" // needed for stdout
import "reflect"
//import "strings"
const esc byte = '\033';
const csic byte = '[';

/* origional prototype:
  var buf [4]byte;
  buf[0] = esc;
  buf[1] = csic;
  buf[2] = tmp[0];
  buf[3] = 'B';
  os.Stdout.Write(buf[0:3]);
  */

func BlankC(num byte)
{
  tmp := Sprintf("%x",num);
  buf := []byte{esc,csic,tmp[0],'@'};
   //Printf("\n<<DEBUG>> 0x%x\n",buf[0:4]);
  os.Stdout.Write(buf[0:4]);
}
func Cup(lines byte)
{
  tmp := Sprintf("%x",lines);
  buf := []byte{esc,csic,tmp[0],'A'};
   //Printf("\n<<DEBUG>> 0x%x\n",buf[0:4]);
  os.Stdout.Write(buf[0:4]);
}
func Cdown(lines byte)
{
  tmp := Sprintf("%x",lines);
  buf := []byte{esc,csic,tmp[0],'B'};
   //Printf("\n<<DEBUG>> 0x%x\n",buf[0:4]);
  os.Stdout.Write(buf[0:4]);
}
func Cforward(lines byte)
{
  tmp := Sprintf("%x",lines);
  buf := []byte{esc,csic,tmp[0],'C'};
   //Printf("\n<<DEBUG>> 0x%x\n",buf[0:4]);
  os.Stdout.Write(buf[0:4]);
}
func Cback(lines byte)
{
  tmp := Sprintf("%x",lines);
  buf := []byte{esc,csic,tmp[0],'D'};
   //Printf("\n<<DEBUG>> 0x%x\n",buf[0:4]);
  os.Stdout.Write(buf[0:4]);
}
func Cnextl(lines byte)
{
  tmp := Sprintf("%x",lines);
  buf := []byte{esc,csic,tmp[0],'E'};
   //Printf("\n<<DEBUG>> 0x%x\n",buf[0:4]);
  os.Stdout.Write(buf[0:4]);
}
func CprevL(lines byte)
{
  tmp := Sprintf("%x",lines);
  buf := []byte{esc,csic,tmp[0],'F'};
   //Printf("\n<<DEBUG>> 0x%x\n",buf[0:4]);
  os.Stdout.Write(buf[0:4]);
}
func Cabs(column byte)
{
  tmp := Sprintf("%x",column);
  buf := []byte{esc,csic,tmp[0],'G'};
   //Printf("\n<<DEBUG>> 0x%x\n",buf[0:4]);
  os.Stdout.Write(buf[0:4]);
}
func Cpos(row byte,col byte)
{
  tmp := Sprintf("%x",row);
  tmp2 := Sprintf("%x",col);
  buf := []byte{esc,csic,tmp[0],';',tmp2[0],'H'};
  //Printf("\n<<DEBUG>> 0x%x\n",buf[0:6]);
  os.Stdout.Write(buf[0:6]);
}
func Ctab(tabs byte)
{
  tmp := Sprintf("%x",tabs);
  buf := []byte{esc,csic,tmp[0],'I'};
   //Printf("\n<<DEBUG>> 0x%x\n",buf[0:4]);
  os.Stdout.Write(buf[0:4]);
}

// ED functions
func EraseBelow()
{
   ed('0'); // call ed with mode of 0
}
func EraseAbove()
{
  ed('1');
}
func EraseAll()
{
  ed('2');
}
func EreaseSavedLines()
{
  ed('3');
}
func ed(mode byte)
{
  buf := []byte{esc,csic,mode,'J'};
   //Printf("\n<<DEBUG>> 0x%x\n",buf[0:4]);
  os.Stdout.Write(buf[0:4]);
}

// EL functions
func LineEraseRight()
{
  el('0');
}
func LineEraseLeft()
{
  el('1');
}
func LineEraseAll()
{
  el('2');
}
func el (mode byte)
{
  buf := []byte{esc,csic,mode,'K'};
  //Printf("\n<<DEBUG>> 0x%x\n",buf[0:4]);
  os.Stdout.Write(buf[0:4]);
}

func InsertLines (lines int)
{
  tmp := Sprintf("%x",lines);
  buf := []byte{esc,csic,tmp[0],'L'};
  //Printf("\n<<DEBUG>> 0x%x\n",buf[0:4]);
  os.Stdout.Write(buf[0:4]);
}
func DeleteLines (lines int)
{
  tmp := Sprintf("%x",lines);
  buf := []byte{esc,csic,tmp[0],'M'};
  //Printf("\n<<DEBUG>> 0x%x\n",buf[0:4]);
  os.Stdout.Write(buf[0:4]);
}
func DeleteChars (chars int)
{
  tmp := Sprintf("%x",chars);
  buf := []byte{esc,csic,tmp[0],'P'};
  //Printf("\n<<DEBUG>> 0x%x\n",buf[0:4]);
  os.Stdout.Write(buf[0:4]);
}
func CScrollUp (lines int)
{
  tmp := Sprintf("%x",lines);
  buf := []byte{esc,csic,tmp[0],'S'};
  //Printf("\n<<DEBUG>> 0x%x\n",buf[0:4]);
  os.Stdout.Write(buf[0:4]);
}
func CScrollDown (lines int)
{
  tmp := Sprintf("%x",lines);
  buf := []byte{esc,csic,tmp[0],'T'};
  //Printf("\n<<DEBUG>> 0x%x\n",buf[0:4]);
  os.Stdout.Write(buf[0:4]);
}
func EreaseChars (chars int)
{
  tmp := Sprintf("%x",chars);
  buf := []byte{esc,csic,tmp[0],'X'};
  //Printf("\n<<DEBUG>> 0x%x\n",buf[0:4]);
  os.Stdout.Write(buf[0:4]);
}
func CBackTab (tabs int)
{
  tmp := Sprintf("%x",tabs);
  buf := []byte{esc,csic,tmp[0],'Z'};
  //Printf("\n<<DEBUG>> 0x%x\n",buf[0:4]);
  os.Stdout.Write(buf[0:4]);  
}

// how to implement cursor attributes, need to choose form 25 settings?
const (
  CattribNormal = "0";
  CattribBold = "1";
  CattribUnderline ="4";
  CattribBlink = "5";
  CattribInverse = "7";
  CattribInvisible = "8";
  CattribNomal = "22";
  CattribNUnderline = "24";
  CattribSteady = "25";
  CattribPositive = "27";
  CattribVisible = "28";
  
  CattribFBlack = "30";
  CattribFRed = "31";
  CattribFGreen = "32";
  CattribFYellow = "33";
  CattribFBlue = "34";
  CattribFMagenta = "35";
  CattribFCyan = "36";
  CattribFWhite = "37";
  CattribFDefault = "39";
  
  CattribBBlack = "40";
  CattribBRed = "41";
  CattribBGreen = "42";
  CattribBYellow = "43";
  CattribBBlue = "44";
  CattribBMagenta = "45";
  CattribBCyan = "46";
  CattribBWhite = "47";
  CattribBDefault = "49";
  
  CattribFBBlack = "90";
  CattribFBRed = "91";
  CattribFBGreen = "92";
  CattribFBYellow = "93";
  CattribFBBlue = "94";
  CattribFBMagenta = "95";
  CattribFBCyan = "96";
  CattribFBWhite = "97";

  CattribBBBlack = "100";
  CattribBBRed = "101";
  CattribBBGreen = "102";
  CattribBBYellow = "103";
  CattribBBBlue = "104";
  CattribBBMagenta = "105";
  CattribBBCyan = "106";
  CattribBBWhite = "107";

)
// Get the i'th arg of the struct value.
// If the arg itself is an interface, return a value for
// the thing inside the interface, not the interface itself.
// from fmt pkg
func getField(v *reflect.StructValue, i int) reflect.Value {
    val := v.Field(i);
    if i, ok := val.(*reflect.InterfaceValue); ok {
        if inter := i.Interface(); inter != nil {
            return reflect.NewValue(inter)
        }
    }
    return val;
}
func Cattrib(a ...)
{
  v := reflect.NewValue(a).(*reflect.StructValue);
  var chars string;
  for curVal :=0; curVal < v.NumField(); curVal++ 
  {
    //field:= getField(v,curVal);
    curChar := v.Field(curVal).(*reflect.StringValue).Get();    
    if curVal != 0
    {
      // dont add on first loop
      chars = Sprintf("%s%c",chars,';');
    }
    chars = Sprintf("%s%s",chars,curChar);
  }
  // how to bild buffer string?
  buf := make([]byte, len(chars) + 3);
  buf[0] = esc;
  buf[1] = csic;
  Printf("\n");
  for i := 2; i < len(chars) + 2; i++
  {
    buf[i] = chars[i-2];
  }
  buf[len(chars) +2] ='m';
  //Printf("<<Debug>>%x\n",buf);
  os.Stdout.Write(buf[0:len(buf)]);

}