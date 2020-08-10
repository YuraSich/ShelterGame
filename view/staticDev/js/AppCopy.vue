<template>
  <v-app id="inspire">
    <!--        <v-navigation-drawer
              v-model="drawer"
              app
              clipped
            >
              <v-list dense>
                <v-list-item link>
                  <v-list-item-action>
                    <v-icon>mdi-view-dashboard</v-icon>
                  </v-list-item-action>
                  <v-list-item-content>
                    <v-list-item-title>Dashboard</v-list-item-title>
                  </v-list-item-content>
                </v-list-item>
                <v-list-item link>
                  <v-list-item-action>
                    <v-icon>mdi-cog</v-icon>
                  </v-list-item-action>
                  <v-list-item-content>
                    <v-list-item-title>Settings</v-list-item-title>
                  </v-list-item-content>
                </v-list-item>
              </v-list>
            </v-navigation-drawer>-->

    <!--    <v-app-bar
          app
          clipped-left
        >
          <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
          <v-toolbar-title>Application</v-toolbar-title>
        </v-app-bar>-->


    <v-container class="grey lighten-5" fluid>
      <v-row
        class=""
        no-gutters
      >
        <v-col
          v-for="n in 4"
          class="mx-0 px-0"
        >
          <v-card
            class="ma-0 pa-0 rounded-sm"
            outlined
            tile

            style="border: 2px solid black"
          >
            <v-row>
              <v-col>
                <div>
                  <div style="text-align: center;  margin:0; padding: 0;  ">
                    <img
                      src="https://pic.rutube.ru/video/78/c8/78c81f80d20b55a8c9c5b23c4eab42b1.jpg?size=m"
                      style="max-width: 100%;  border-radius: 0 0 0 0;  display: block; margin: 0 auto;"
                      alt="lorem">
                  </div>
                  <div style="text-align: center; background-color: #C3C3C3; margin:0; padding: 0;">
                    <span style="color: white">{{user.name}}</span>
                    <v-container>
                      <v-row>
                        <table style="display: table" >
                          <tr>
                            <td>
                              <div style="padding-left: 5px; padding-right: 5px"><i class="fas fa-anchor"></i></div>
                              <div>{{user.p_4}}</div>
                            </td>
                            <td>
                              <div style="padding-left: 5px; padding-right: 5px"><i class="fas fa-anchor"></i></div>
                              <div>{{user.p_4}}</div>
                            </td>
                            <td>
                              <div style="padding-left: 5px; padding-right: 5px"><i class="fas fa-anchor"></i></div>
                              <div>{{user.p_4}}</div>
                            </td>
                          </tr>
                        </table>
                      </v-row>
                      <v-row>
                        <v-col style="height:60px; border: 1px solid black; margin:0; padding: 2px;" cols="4">
                          <div style="height:100%; display: flex; justify-content: space-around; align-items: center">
                            <div style="padding-left: 5px; padding-right: 5px"><i class="fas fa-anchor"></i></div>
                            <div>{{user.p_4}}</div>
                          </div>
                        </v-col>
                        <v-col style="height:70px; border: 1px solid black; margin:0; padding: 2px;" cols="4">
                          <div style="height:100%; display: flex; justify-content: space-between; align-items: center; ">
                            <div style="padding-left: 2px; padding-right: 2px"><i class="fas fa-anchor"></i></div>
                            <div style="width: 60px; padding-left: 2px; padding-right: 2px; overflow-wrap:  break-word ; ">{{user.p_5}}</div>
                          </div>
                        </v-col>
                        <v-col style="height:60px; border: 1px solid black; margin:0; padding: 2px;" cols="4">
                          <div style="height:100%; display: flex; justify-content: space-around; align-items: center">
                            <div style="padding-left: 5px; padding-right: 5px"><i class="fas fa-anchor"></i></div>
                            <div>{{user.p_6}}</div>
                          </div>
                        </v-col>
                      </v-row>
                    </v-container>
                  </div>
                </div>
                <div style=" font-size: 12px">
                  <!--                  <div>Профессия: {{user.job}}</div>-->
                  <!--                  <div>Пол: {{user.sex}}</div>-->
                  <!--                  <div>Возраст: {{user.age}}</div>-->
                  <!--                  <div>Здоровье: {{user.health}}</div>-->
                  <!--                  <div>Ориентация: {{user.orientation}}</div>-->
                  <!--                  <div>Фобия: {{user.phobias}}</div>-->
                </div>
              </v-col>
            </v-row>
          </v-card>

        </v-col>
      </v-row>

    </v-container>


    <!--    <v-footer app>
          <span>&copy; {{ new Date().getFullYear() }}</span>
        </v-footer>-->
  </v-app>
</template>

<script>
  export default {
    props: {
      source: String,
    },
    data: () => ({
      drawer: false,
      user: {
        p_1: "Садовник",
        p_2: "Полный привод",
        p_3: "Еда",
        p_4: "45",
        p_5: "Агрессивный1123",
        p_6: "Глухой",
        p_7: "Алкоголик",
        p_8: "Лгун",
        p_9: "Ослофобия",
      },
      // users: {
      //   rows: [
      //     [
      //       {
      //         name: "Петя",
      //         age: 123,
      //         orientation: "натурал",
      //         health: "--",
      //         job: "1c",
      //         sex: "муж",
      //         phobias: "Боровцов",
      //       }, {
      //       name: "Петя",
      //       age: 123,
      //       orientation: "натурал",
      //       health: "--",
      //       job: "1c",
      //       sex: "муж",
      //       phobias: "Боровцов",
      //     }, {
      //       name: "Петя",
      //       age: 123,
      //       orientation: "натурал",
      //       health: "--",
      //       job: "1c",
      //       sex: "муж",
      //       phobias: "Боровцов",
      //     },
      //       {
      //         name: "Петя",
      //         age: 123,
      //         orientation: "натурал",
      //         health: "100",
      //         job: "1c",
      //         sex: "муж",
      //         phobias: "Боровцов",
      //       },
      //       {
      //         name: "Петя",
      //         age: 123,
      //         orientation: "натурал",
      //         health: "100",
      //         job: "1c",
      //         sex: "муж",
      //         phobias: "Боровцов",
      //       },
      //       {
      //         name: "Петя",
      //         age: 123,
      //         orientation: "натурал",
      //         health: "100",
      //         job: "1c",
      //         sex: "муж",
      //         phobias: "Боровцов",
      //       },
      //       {
      //         name: "Петя",
      //         age: 123,
      //         orientation: "натурал",
      //         health: "100",
      //         job: "1c",
      //         sex: "муж",
      //         phobias: "Боровцов",
      //       },
      //       {
      //         name: "Петя",
      //         age: 123,
      //         orientation: "натурал",
      //         health: "100",
      //         job: "1c",
      //         sex: "муж",
      //         phobias: "Боровцов",
      //       },
      //
      //     ],
      //     [
      //       {
      //         name: "Петя",
      //         age: 123,
      //         orientation: "натурал",
      //         health: "--",
      //         job: "1c",
      //         sex: "муж",
      //         phobias: "Боровцов",
      //       },
      //       {
      //       name: "Петя",
      //       age: 123,
      //       orientation: "натурал",
      //       health: "--",
      //       job: "1c",
      //       sex: "муж",
      //       phobias: "Боровцов",
      //     },
      //       {
      //       name: "Петя",
      //       age: 123,
      //       orientation: "натурал",
      //       health: "--",
      //       job: "1c",
      //       sex: "муж",
      //       phobias: "Боровцов",
      //     },
      //       {
      //         name: "Петя",
      //         age: 123,
      //         orientation: "натурал",
      //         health: "100",
      //         job: "1c",
      //         sex: "муж",
      //         phobias: "Боровцов",
      //       }
      //
      //     ],
      //   ],
      // }
    }),
    created() {
      // this.$vuetify.theme.dark = true
    },
  }
</script>
