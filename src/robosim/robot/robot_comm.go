package robot

import (
    "fmt"
    "log"
    "net"

    "robosim/common"
)

type robotMsg interface{}

type RobotComm struct {
    host string
    id int
    quededMsgs []robotMsg
    // consoleAddr interface{}
    sock net.Listener
}

func (rc *RobotComm) open(host string, id int) {
    rc.host = host
    rc.id = id
    rc.quededMsgs = nil

    addr := fmt.Sprintf("%s:%d", host, common.DefConsolePort + id)
    sock, err := net.Listen("tcp", addr)
    if err != nil {
        log.Fatal(err)
    }
    rc.sock = sock
}

func (rc *RobotComm) checkMsgs(waitFor int) *robotMsg {
        // make sure we the socket has been opened
        if rc.sock == nil {
            log.Fatal(fmt.Sprintf("need to open robot (%d) comm", rc.id))
        }

        // handle any queued messages
        if (waitFor <= 0) && (len(rc.quededMsgs) > 0) {
            for _, msg := range rc.quededMsgs {
                rc.handleMsg(msg)
            }
            rc.quededMsgs = nil
        }

        // # keep going while there's messages waiting
        // while ( 1 ):
        //     # see if there's any messages waiting
        //     ( i, o, e ) = select.select( [self.sock], [], [], 0.01 )
        //     if ( not i ):
        //         break

        //     # read and handle a message
        //     ( msg, (host, port) ) = self.sock.recvfrom( 65536 )
        //     msg = marshal.loads( msg )

        //     if ( msg[0] == wait_for ):
        //         return msg
        //     elif ( wait_for ):
        //         self.queued_msgs.append( msg )
        //     else:
        //         self.handle_msg( msg )

    return nil
}

    // def wait_for_msg( self, wait_for ):
    //     msg = None
    //     while ( not msg ):
    //         msg = self.check_msgs( wait_for=wait_for )
    //     return msg

func (rc *RobotComm) handleMsg(msg robotMsg) {
//     if ( msg[0] in RobotComm.handlers.keys() ):
//         RobotComm.handlers[ msg[0] ]( self, msg )
//     else:
//         print 'Error: unregistered message number:', msg[0]
}
