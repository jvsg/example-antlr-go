grammar Expr;

prog:   stat+ ;

stat:   command EOF                   # commandStat
    |   NEWLINE                       # blank
    ;


command: sysconf ;// | hsm | service ;

sysconf: 'sysconf' (appliance) ;

appliance: 'appliance' (factory) ;

factory: 'factory' (force_flag)? ;

force_flag: '-f'
          | '--force'
          ;

value:  INT
     |  STRING
     ;


SYSCONF :   'sysconf' ;
HSM     :   'hsm' ;
SERVICE :   'service' ;
ID  :   [a-zA-Z]+ ; // match identifiers
INT :   [0-9]+ ;    // match integers
STRING : '"' .*? '"' ; // match strings
NEWLINE:'\r'? '\n' ; // match newlines
WS  :   [ \t]+ -> skip ; // skip spaces, tabs

