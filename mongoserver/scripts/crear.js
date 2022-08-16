db = connect( 'mongodb://localhost/assessments' );

db.apps.drop();
db.wf.drop();
db.roles.drop();
db.usuarios.drop();
db.organizaciones.drop();
db.procesos.drop();
db.assessments.drop();
db.buenaspracticas.drop();
db.proyectos.drop();
db.positions.drop();

db.positions.insertMany([
    {_id:1,position:"Portfolio Manager"},
    {_id:2,position:"Program Manager"},
    {_id:3,position:"Project Manager"},
    {_id:4,position:"Product Manager"},
    {_id:5,position:"Product Owner"},
    {_id:6,position:"Business Analyst"},
    {_id:7,position:"Team lead"},
    {_id:8,position:"Scrum master"},
    {_id:9,position:"Solution architect"},
    {_id:10,position:"Developer"},
    {_id:11,position:"QA"},
    {_id:12,position:"Data engineer"},
    {_id:13,position:"QA Manager"},
]);


db.proyectos.insertMany([
    {
        _id:1,
        proy:"TDF",
        teams: [
            {   
                id:1,
                team:"TDF-Collection",
                from:"2022-03-01",
                upto:null
            },
            {   
                id:2,
                team:"TDF-Collection",
                from:"2022-03-01",
                upto:null
            }            
        ]
    }
]);

db.buenaspracticas.insertMany( [
    {
        _id : 1,
       practica: 'CMMI-DEV',
       categories:[
        {
            cat:"doing",
            capabilitiareas:[
            {
                ca:"ensuring quality",
                practiceareas:[
                {
                    pa:"Peer Review",
                    practices:[
                    {
                        id:1,
                        practice:"Perform reviews of work products and record issues",
                        ml:1
                    },
                    {
                        id:2,
                        practice:"Develop and keep updated procedures and supporting materials used to preparre for and perform peer review",
                        ml:2
                    },
                    {
                        id:3,
                        practice:"Select work products to be peer reviewed",
                        ml:2
                    },
                    {
                        id:4,
                        practice:"Prepare and perform peer review on selected work products using stablished procedures",
                        ml:2
                    },
                    {
                        id:5,
                        practice:"Resolve issues identified in peer reviews",
                        ml:2
                    },
                    {
                        id:6,
                        practice:"Analyze the results and data from peer reviews",
                        ml:3
                    }                                        
                ]
                },              
                {
                    pa:"Verification & Validation",
                    practices:[
                        {
                            id:20,
                            practice:"Perform verification to ensure the requirements are implemented and record and communicate results",
                            ml:1
                        },
                        {
                            id:21,
                            practice:"Perform validation to ensure the solution will function as intendeed and its target environmente and record and communicate results",
                            ml:1
                        },
                        {
                            id:22,
                            practice:"Select components and methods for verification and validation",
                            ml:2
                        },
                        {
                            id:23,
                            practice:"Develop, keep updated and use the environment needed to support verifcation and validation",
                            ml:2
                        },
                        {
                            id:24,
                            practice:"Develop, keep update and follow procedures for verification and validation",
                            ml:2
                        },
                        {
                            id:25,
                            practice:"Develop, keep updated, and use criteria for verification and validation",
                            ml:3
                        },
                        {
                            id:26,
                            practice:"Analyze and communicate verifcation and validation results",
                            ml:3
                        }                                         
                    ]
                } 
             ]
            }
            ]
        }
       ],
       genericGoals:[
        {
            gg:"1.- Achieve specific goals (of each process area)",
            gps: [
                {gp:"1.1.- Perform specific practices"}
            ]
        },
        {
            gg:"2.- Institutionalize a managed process",
            gps: [
                {gp:"2.1.- Establish an organizational Policy"},
                {gp:"2.2.- Plan the process"},
                {gp:"2.3.- Provide resources"},
                {gp:"2.4.- Assign responsabilities"},
                {gp:"2.5.- Train people"},
                {gp:"2.6.- Control work products"},
                {gp:"2.7.- Identify and involve relevant stakeholders"},
                {gp:"2.8.- Monitor and control the process"},
                {gp:"2.9.- Objetively evaluate adherance"},
                {gp:"2.10.- Review status with higher management"}
            ]
        },
        {
            gg:"3.- Institutionalize a defined process",
            gps: [
                {gp:"3.1.- Establish an defined process"},
                {gp:"3.2.- Collect Process related experiences"}
            ]
        }                 
       ]

    }
]
);

db.organizaciones.insertMany( [
    {
       _id: 1,
       nombre: 'ey',
       padre: null
    },
    {
        _id: 2,
        nombre: 'ey-tax',
        padre: 1
     },
     {
        _id: 3,
        nombre: 'ey-CT',
        padre: 1
     },
     {
        _id: 4,
        nombre: 'ey-CT-GDS',
        padre: 3
     }
 ] );

 db.procesos.insertMany( [
    {
       _id : 1,
       proceso: 'Desarrollo de sistemas',
       organizacionid: 4,
       procesopadreid : null,
       buenaPracticaAsoc :[{bpid:1, practices:[1,2,3]}]
    },
    {
        _id : 2,
        proceso: 'Análisis de requerimientos',
        organizacionid: 4,
        procesopadreid : 1
     },
     {
         _id : 3,
         proceso: 'Creación de prototipos',
         organizacionid: 4,
         procesopadreid : 1
    }
 ] );
 
 
db.apps.insertMany( [
    {
       _id:1,
       nombre:'blog',
       wfs:[1,2]
    },
    {
        _id:2,
        nombre:'Calidad',
        wfs:[20,21,22]
     }
 ] );

 db.wf.insertMany( [
    {
       _id:1,
       nombre:'Agregar mensaje blog',
       appid:1
    },
    {
        _id:2,
        nombre:'Agregar tema blog',
        appid:1
     },
     {
        _id:20,
        nombre:'Agregar Proceso',
        appid:2
     },
     {
        _id:21,
        nombre:'Agregar Indicador',
        appid:2
     },
     {
        _id:22,
        nombre:'Agregar Buena Prática',
        appid:2
     },
 ] );

db.roles.insertMany( [
    {
       _id: 1,
       rol: 'Administrador',
       permisos:[
        {app:1, wf:[1,2,3]},
        {app:2, wf:[1,2]}
       ]
    },
    {
        _id: 2,
        rol: 'Adm Grupo',
        permisos:[
         {app:1, wf:[1]},
         {app:2, wf:[1]}
        ]
     }
 ] );
 
db.usuarios.insertMany( [
   {
      name: 'Andres',
      surename: 'Kouvach',
      roles: [1,2],
      assignments:[
        {project:1, team:1,since:"2022-01-01", upto:"2022-04-30", hsxweek:40, position:3},
        {project:2, team:1,since:"2022-05-01", upto:null, hsxweek:20, position:3}
      ]
   }
] );

