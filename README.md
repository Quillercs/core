# core

## Running with Docker
  
  *Clone this repository*
  
    $ git clone git@github.com:quillercs/core.git
  
  *Run docker-compose to run the container (and build at first time the image)*
    
    $ docker-compose up -d
  
  *See the log*
  
    $ docker-compose logs core
  
  *In other terminal*
  
    $ curl -i 'http://{HOST}:3000/' -X POST -H "Content-Type: application/json"  -d '{"platform": "freeswitch", "version": "1.6.2"}}'
    $ curl -i 'http://{HOST}:3000/' -X POST -H "Content-Type: application/json"  -d '{"platform": "freeswitch", "version": "1.6.4"}}'
    $ curl -i 'http://{HOST}:3000/' -X POST -H "Content-Type: application/json"  -d '{"platform": "asterisk", "version": "1.11"}}'
    $ curl -i 'http://{HOST}:3000/' -X POST -H "Content-Type: application/json"  -d '{"platform": "yate", "version": "5.5"}}'
