#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define QUERY_SIZE 2048
#define ARGS_SIZE 3072
void main()
{

    char query[QUERY_SIZE];
    int status;
    FILE *filePointer; 
    char *reqArgs;
    char dataToBeRead[50];
    char songSelect;
    int songSelectInt;
    char *playerArgs[ARGS_SIZE];
    char ids[5][128];
    char titles[5][128];
    char i = 0;
   

    while(1){
        for(int i = 0; i < QUERY_SIZE; i++)
            query[i] = NULL;
        status = 0;
        songSelect = NULL;
        songSelectInt = 0;
        i = 0;

        printf("Enter your query\n");
        fgets(query, QUERY_SIZE, stdin);
        
        reqArgs = malloc(ARGS_SIZE * sizeof(char));
        strcat(reqArgs, "./req ");
        strcat(reqArgs, query);
        status = system(reqArgs);
        
        
        filePointer = fopen("/tmp/vidOut", "r");

        while( fgets ( dataToBeRead, 50, filePointer ) != NULL ) 
        {       
            // Print the dataToBeRead  
            strcpy(ids[i], dataToBeRead);
            //printf("%s", dataToBeRead);
            if(fgets ( dataToBeRead, 50, filePointer ) != NULL){
                strcpy(titles[i], dataToBeRead);
                //printf("%s", dataToBeRead);
            }

            i++;
        } 

        fclose(filePointer);

        for(int j = 1; j <= 5; j++){
            printf("%d: %s %s", j, ids[j-1], titles[j-1]);
        }
        printf("Choose the song: (1-5)\n");
        songSelect = fgetc(stdin);
        songSelectInt = (int)(songSelect - '0');
        if(songSelect == 'q' || songSelect == 'Q')
            continue;
        char theURL[512] = {NULL};
        strcat(theURL, "https://www.youtube.com/watch?v=");

        strcat(theURL, ids[songSelectInt - 1]);
        strtok(theURL, "\n");
        for (int k = 0; k < ARGS_SIZE; k++)
            playerArgs[k] = NULL;
        strcat(playerArgs, "mpv ");
        strcat(playerArgs, theURL);
        strcat(playerArgs, " --no-video");
        //printf("\n%s\n", playerArgs);
        status = system(playerArgs);
    }
}
