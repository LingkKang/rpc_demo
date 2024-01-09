var searchIndex = JSON.parse('{\
"client":{"doc":"","t":"RRFFFFAFFFAFAADLLMLLLLLLLLLLFAAAFFGNRRDENNLLLLLLLLLLLLMMMLLLLLLLLLLLFF","n":["MAX_TASKS","URL","bytes_to_hex_str","collect_a_task","get_random_f64","get_sides","logger","main","process_message","process_response","protocol","receive_message","logger","time","Logger","borrow","borrow_mut","default_level","enabled","flush","from","init","into","log","try_from","try_into","type_id","vzip","get_formatted_time","checksum","message","protocol","generate_checksum","verify_checksum","Byte","ERROR","MAX_PAYLOAD_SIZE","MAX_PROTOCOL_SIZE","Message","MessageType","REQUEST","RESPONSE","borrow","borrow","borrow_mut","borrow_mut","from","from","from_byte","get_head","get_payload","get_type","into","into","msg_length","msg_payload","msg_type","new","new_request","to_byte","try_from","try_from","try_into","try_into","type_id","type_id","vzip","vzip","deserialize","serialize"],"q":[[0,"client"],[12,"client::logger"],[14,"client::logger::logger"],[28,"client::logger::time"],[29,"client::protocol"],[32,"client::protocol::checksum"],[34,"client::protocol::message"],[68,"client::protocol::protocol"],[70,"alloc::string"],[71,"alloc::vec"],[72,"std::net::tcp"],[73,"log"],[74,"log"],[75,"log"],[76,"core::any"]],"d":["","","Converts a <code>Vec</code> of <code>Byte</code>s to a <code>String</code> of hexadecimal numbers.","Basically generates a task of calculating the hypotenuse …","Generates a random <code>f64</code> number.","Generates a pair of random <code>f64</code> numbers, which will be used …","","","Process a message, basically check the type of the message …","Process a message of type <code>MessageType::RESPONSE</code>.","","Receives a message from the server and parse it.","Custom logger module.","Time utility for the logger.","A custom logger struct that uses stdout to print logs.","","","The default level of the logger.","Check if current message level is enabled for logging.","Flush the logger. As stdout is used, no need to flush, so …","Returns the argument unchanged.","Static method to initialize the logger with an optional …","Calls <code>U::from(self)</code>.","Log the message.","","","","","Returns the current time in a formatted string. The format …","The checksum submodule provides functions for generating …","A protocol message is defined as follows:","","Use XOR to generate checksum.","Verify checksum by comparing it with the generated checksum","A byte is logically equivalent to an 8-bit unsigned …","Message type for error, maps to <code>0b00</code> in the first two bits.","Maximum size of payload comes from <code>MAX_PROTOCOL_SIZE - 2</code>.","Maximum size of a protocol message. It is counted in bytes.","A protocol message. See the module-level documentation for …","The type (action) of a message which is the first two bits …","Message type for request, maps to <code>0b01</code> in the first two …","Message type for response, maps to <code>0b10</code> in the first two …","","","","","Returns the argument unchanged.","Returns the argument unchanged.","Convert a <code>Byte</code> to a <code>MessageType</code> as defined in the protocol.","Get the first byte of the message, which is the first two …","Get the payload of the message in <code>Vec&lt;Byte&gt;</code>.","","Calls <code>U::from(self)</code>.","Calls <code>U::from(self)</code>.","The length of the entire message, which will be encoded in …","The payload of the message, should always be a multiple of …","The type (action) of the message, which will be encoded in …","Create a new message.","Create a new message in <code>MessageType::REQUEST</code> type.","Convert a <code>MessageType</code> to a <code>Byte</code> as defined in the protocol.","","","","","","","","","Deserializes a binary format into a message. Some checks …","Serializes all message data into a binary format, added …"],"i":[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,9,9,9,9,9,9,9,9,9,9,9,9,9,0,0,0,0,0,0,0,17,0,0,0,0,17,17,17,7,17,7,17,7,17,7,7,7,17,7,7,7,7,7,7,17,17,7,17,7,17,7,17,7,0,0],"f":[0,0,[[[2,[1]]],3],[[],[[4,[1]]]],[[],5],[[],[[6,[5,5]]]],0,[[],6],[7,6],[7,6],0,[8,7],0,0,0,[-1,-2,[],[]],[-1,-2,[],[]],0,[[9,10],11],[9,6],[-1,-1,[]],[[[13,[12]]],6],[-1,-2,[],[]],[[9,14],6],[-1,[[15,[-2]]],[],[]],[-1,[[15,[-2]]],[],[]],[-1,16,[]],[-1,-2,[],[]],[[],3],0,0,0,[[[2,[1]]],1],[[[2,[1]],1],11],0,0,0,0,0,0,0,0,[-1,-2,[],[]],[-1,-2,[],[]],[-1,-2,[],[]],[-1,-2,[],[]],[-1,-1,[]],[-1,-1,[]],[1,[[15,[17,18]]]],[7,1],[7,[[4,[1]]]],[7,17],[-1,-2,[],[]],[-1,-2,[],[]],0,0,0,[[17,1,[4,[1]]],7],[[[4,[1]]],7],[17,1],[-1,[[15,[-2]]],[],[]],[-1,[[15,[-2]]],[],[]],[-1,[[15,[-2]]],[],[]],[-1,[[15,[-2]]],[],[]],[-1,16,[]],[-1,16,[]],[-1,-2,[],[]],[-1,-2,[],[]],[[[2,[1]]],[[15,[7,18]]]],[7,[[4,[1]]]]],"c":[],"p":[[15,"u8"],[15,"slice"],[3,"String",70],[3,"Vec",71],[15,"f64"],[15,"tuple"],[3,"Message",34],[3,"TcpStream",72],[3,"Logger",14],[3,"Metadata",73],[15,"bool"],[4,"LevelFilter",73],[4,"Option",74],[3,"Record",73],[4,"Result",75],[3,"TypeId",76],[4,"MessageType",34],[15,"str"]],"b":[]}\
}');
if (typeof window !== 'undefined' && window.initSearch) {window.initSearch(searchIndex)};
if (typeof exports !== 'undefined') {exports.searchIndex = searchIndex};