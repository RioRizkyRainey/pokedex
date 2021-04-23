package main

import (
	"fmt"

	"github.com/RioRizkyRainey/pokedex/internal/gateway/service"
)

func main() {
	fmt.Print(
		`
                                    -/osyyyhhhyyyso/-\'                                             
                                 .+yhso+///////////+oshyo-                                          
                              \'/hho///////++oooooooo++//oyh+\'         \'/s:                          
                             /ds+/////oosyyyyyyyyyyyyyyyso+hm-     -+ys:\'                           
                           .hy+////osyyyyyyyyyyyyyyyyyyhdds/\'   :syo-   .+yy                        
                          -do////oyyyyyyyyyyyyyyyyyyhdho-    :. -\'  \':ohysm:                        
                         -mo///oyyyyyyyyhddysssyhddy+.   -+\' \'   ./yhs+//yy                         
                        \'ds///syyyyyyyhdo.       .\'  \'/sy+.   -oyho/////+m-                         
                        /d///syyyyyyydy\' \'+yhdhy+\'  :s/\'  \':sddy////////hs                          
                        yy//oyyyyyyydd  .ddyyyyyyd-    .+ydhyyys///////om.                          
                        ds/+yyyyyyyhm+  omyyys+//yy  -ddhyyyyyy+///////do                           
                        hy/oyyyyhdy+.   -my+////+m/  sdyyyyyyys///////om\'                           
                        od/shdds:\'  \'/:  .shyyyhy-  /myyyyyyyy+///////d+                            
                        .mhy+-   -+ys:\'     \'..   -ydyyyyyyyhys+/////sd\'                            
                         .\'     o+-   .+hdyo+/+osddhyyyyyyddsooyd+///m/                             
                          \' \'o    \':oddhyyyyhhhyyyyyyyyyyyN+////yh//yh                              
                         \'+    ./yddyyyyyyyyyyyyyyyyyyys+/hho++smo/+m:                              
                             .hdsssyyyyyyyyyyyyyyyyss+/////oyyys+//yy                               
                              .shs+/+ooossssssooo+////////////////+m-                               
                                \'/yhyo///////////////oshd+////////hs                                
                                    -+syyhyysssyyyyys+-\':do//////+m.                                
                                         \'\'.---..        -ds/////ho                                 
                                                          \'hy///om\'                                 
                                                           \'sh+/d+                                  
                                                             +dyd\'                                  
                                                              :m/                                                                       
`)

	fmt.Println()
	fmt.Println(`Pokédex.
	The Pokédex s a digital encyclopedia created by Professor Oak as an invaluable tool to Trainers in the Pokémon world.
	It gives information about all Pokémon in the world that are contained in its database, although it differs in how it acquires and presents information over the different media.`)

	service.Start()
}
