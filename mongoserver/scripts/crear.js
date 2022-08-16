db = connect( 'mongodb://localhost/assessments' );

db.apps.drop();
db.wf.drop();
db.roles.drop();
db.usuarios.drop();
db.organizaciones.drop();
db.procesos.drop();
db.assessments.drop();

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

 db.buenaspracticas.insertMany( [
    {
       _id : 1,
       practica: 'CMMI-DEV',
       categories:[
        {cat:"doing",
        capabilitiareas:[
            {ca:"ensuring quality",
             practiceareas:[
                {pa:"Ensuring quality",
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
                ]}
             ]
            }
        ]}

       ]

    }
]);

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
      nombre: 'Andres',
      apellido: 'Kouvach',
      roles: [1,2]
   }
] );

