<template>
  <v-app id="inspire" class="my-yellow">
    <div class="App" style="height: 100vh;">

      <v-navigation-drawer
        v-model="drawer"
        absolute
        temporary
      >
        <v-list dense>
          <v-list-item link @click="drawer = false">
            <v-list-item-action>
              <v-icon> mdi-close-box-outline</v-icon>
            </v-list-item-action>
            <v-list-item-content>
              <v-list-item-title>Close</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
          <v-list-item link>
            <v-list-item-action>
              <v-icon>mdi-cog</v-icon>
            </v-list-item-action>
            <v-list-item-content>
              <v-list-item-title v-if="flag_keking" @click="stopKek">Перестать Кекать</v-list-item-title>
              <v-list-item-title v-else @click="startKek">Кекнуть</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list>
      </v-navigation-drawer>

      <div class="btn-close">
        <v-btn class="mx-2"
               fab dark small color="primary"
               @click="drawer = true">
          <v-icon>mdi-menu</v-icon>
        </v-btn>
      </div>

      <div class="" style="height: 100%; margin: 0; padding-top: 15px;">
        <div class="my-container">
          <div class="one-user my-blue" v-for="countUser in 4">
            <div class="img-container">
              <img :src="users[countUser - 1].img"
                   alt="lorem">
            </div>
            <div class="one-user__info">
              <div class="one-user__info-row">
                <div class="my-blue cell" :title="users[countUser - 1].p_1">
                  <i :class="users[countUser - 1].p_1_icon + ' icon'"></i>
                  {{users[countUser - 1].p_1}}
                </div>
                <div class="my-blue cell" :title="users[countUser - 1].p_2">
                  <i :class="users[countUser - 1].p_2_icon + ' icon'"></i>
                  {{users[countUser - 1].p_2}}
                </div>
                <div class="my-blue cell" :title="users[countUser - 1].p_3">
                  <i :class="users[countUser - 1].p_3_icon + ' icon'"></i>
                  {{users[countUser - 1].p_3}}
                </div>
              </div>
              <div class="one-user__info-row">
                <div class="my-blue cell" :title="users[countUser - 1].p_4">
                  <i :class="users[countUser - 1].p_4_icon + ' icon'"></i>
                  {{users[countUser - 1].p_4}}

                </div>
                <div class="my-blue cell" :title="users[countUser - 1].p_5">
                  <i :class="users[countUser - 1].p_5_icon + ' icon'"></i>
                  {{users[countUser - 1].p_5}}
                </div>
                <div class="my-blue cell" :title="users[countUser - 1].p_6">
                  <i :class="users[countUser - 1].p_6_icon + ' icon'"></i>
                  {{users[countUser - 1].p_6}}
                </div>
              </div>
              <div class="one-user__info-row">
                <div class="my-blue cell" :title="users[countUser - 1].p_7">
                  <i :class="users[countUser - 1].p_7_icon + ' icon'"></i>
                  {{users[countUser - 1].p_7}}

                </div>
                <div class="my-blue cell" :title="users[countUser - 1].p_8">
                  <i :class="users[countUser - 1].p_8_icon + ' icon'"></i>
                  {{users[countUser - 1].p_8}}
                </div>
                <div class="my-blue cell" :title="users[countUser - 1].p_9">
                  <i :class="users[countUser - 1].p_9_icon + ' icon'"></i>
                  {{users[countUser - 1].p_9}}
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="my-container">
          <div class="one-user my-blue">
            <div class="img-container">
              <img :src="users[4].img"
                   alt="lorem">
            </div>
            <div class="one-user__info">
              <div class="one-user__info-row ">
                <div class="my-blue cell" :title="users[4].p_2">
                  <i :class="users[4].p_1_icon + ' icon'"></i>
                  {{users[4].p_1}}
                </div>
                <div class="my-blue cell  " title="м.(полный привод)">
                  <i :class="users[4].p_2_icon + ' icon'"></i>
                  {{users[4].p_2}}
                </div>
                <div class="my-blue cell" :title="users[4].p_2">
                  <i :class="users[4].p_3_icon + ' icon'"></i>
                  {{users[4].p_3}}
                </div>
                <!--        <div class="my-yellow cell">
                          <div class="icon">
                            <i class="fas fa-anchor"></i>
                          </div>
                          <div class="one-user__info-title">
                            Еда
                          </div>
                        </div>-->
              </div>
              <div class="one-user__info-row">
                <div class="my-blue cell" :title="users[4].p_2">
                  <i :class="users[4].p_4_icon + ' icon'"></i>
                  {{users[4].p_4}}

                </div>
                <div class="my-blue cell  " title="м.(полный привод)">
                  <i :class="users[4].p_5_icon + ' icon'"></i>
                  {{users[4].p_5}}
                </div>
                <div class="my-blue cell" :title="users[4].p_2">
                  <i :class="users[4].p_6_icon + ' icon'"></i>
                  {{users[4].p_6}}

                </div>
                <!--        <div class="my-yellow cell">
                          <div class="icon">
                            <i class="fas fa-anchor"></i>
                          </div>
                          <div class="one-user__info-title">
                            Еда
                          </div>
                        </div>-->
              </div>
              <div class="one-user__info-row">
                <div class="my-blue cell" :title="users[4].p_2">
                  <i :class="users[4].p_7_icon + ' icon'"></i>
                  {{users[4].p_7}}

                </div>
                <div class="my-blue cell  " title="м.(полный привод)">
                  <i :class="users[4].p_8_icon + ' icon'"></i>
                  {{users[4].p_8}}

                </div>
                <div class="my-blue cell" :title="users[4].p_2">
                  <i :class="users[4].p_9_icon + ' icon'"></i>
                  {{users[4].p_9}}
                </div>
                <!--        <div class="my-yellow cell">
                          <div class="icon">
                            <i class="fas fa-anchor"></i>
                          </div>
                          <div class="one-user__info-title">
                            Еда
                          </div>
                        </div>-->
              </div>
            </div>
          </div>
          <div class="control-pane my-blue">1</div>
          <div class="one-user my-blue">
            <div class="img-container">
              <img :src="users[5].img"

                   alt="lorem">
            </div>
            <div class="one-user__info">
              <div class="one-user__info-row">
                <div class="my-blue cell" :title="users[5].p_2">
                  <i :class="users[5].p_1_icon + ' icon'"></i>
                  {{users[5].p_1}}
                </div>
                <div class="my-blue cell  " title="м.(полный привод)">
                  <i :class="users[5].p_2_icon + ' icon'"></i>
                  {{users[5].p_2}}
                </div>
                <div class="my-blue cell" :title="users[5].p_2">
                  <i :class="users[5].p_3_icon + ' icon'"></i>
                  {{users[5].p_3}}
                </div>
                <!--        <div class="my-yellow cell">
                          <div class="icon">
                            <i class="fas fa-anchor"></i>
                          </div>
                          <div class="one-user__info-title">
                            Еда
                          </div>
                        </div>-->
              </div>
              <div class="one-user__info-row">
                <div class="my-blue cell" :title="users[5].p_2">
                  <i :class="users[5].p_4_icon + ' icon'"></i>
                  {{users[5].p_4}}

                </div>
                <div class="my-blue cell  " title="м.(полный привод)">
                  <i :class="users[5].p_5_icon + ' icon'"></i>
                  {{users[5].p_5}}
                </div>
                <div class="my-blue cell" :title="users[5].p_2">
                  <i :class="users[5].p_6_icon + ' icon'"></i>
                  {{users[5].p_6}}

                </div>
                <!--        <div class="my-yellow cell">
                          <div class="icon">
                            <i class="fas fa-anchor"></i>
                          </div>
                          <div class="one-user__info-title">
                            Еда
                          </div>
                        </div>-->
              </div>
              <div class="one-user__info-row">
                <div class="my-blue cell" :title="users[5].p_2">
                  <i :class="users[5].p_7_icon + ' icon'"></i>
                  {{users[5].p_7}}

                </div>
                <div class="my-blue cell  " title="м.(полный привод)">
                  <i :class="users[5].p_8_icon + ' icon'"></i>
                  {{users[5].p_8}}

                </div>
                <div class="my-blue cell" :title="users[5].p_2">
                  <i :class="users[5].p_9_icon + ' icon'"></i>
                  {{users[5].p_9}}
                </div>
                <!--        <div class="my-yellow cell">
                          <div class="icon">
                            <i class="fas fa-anchor"></i>
                          </div>
                          <div class="one-user__info-title">
                            Еда
                          </div>
                        </div>-->
              </div>
            </div>
          </div>
        </div>

        <div class="my-container">
          <div class="one-user my-blue" v-for="countUser in 4">
            <div class="img-container">
              <img :src="users[countUser + 5].img"
                   alt="lorem">
            </div>
            <div class="one-user__info">
              <div class="one-user__info-row">
                <div class="my-blue cell" :title="users[countUser + 5].p_1">
                  <i :class="users[countUser + 5].p_1_icon + ' icon'"></i>
                  {{users[countUser + 5].p_1}}
                </div>
                <div class="my-blue cell" :title="users[countUser + 5].p_2">
                  <i :class="users[countUser + 5].p_2_icon + ' icon'"></i>
                  {{users[countUser + 5].p_2}}
                </div>
                <div class="my-blue cell" :title="users[countUser + 5].p_3">
                  <i :class="users[countUser + 5].p_3_icon + ' icon'"></i>
                  {{users[countUser + 5].p_3}}
                </div>
              </div>
              <div class="one-user__info-row">
                <div class="my-blue cell" :title="users[countUser + 5].p_4">
                  <i :class="users[countUser + 5].p_4_icon + ' icon'"></i>
                  {{users[countUser + 5].p_4}}

                </div>
                <div class="my-blue cell" :title="users[countUser + 5].p_5">
                  <i :class="users[countUser + 5].p_5_icon + ' icon'"></i>
                  {{users[countUser + 5].p_5}}
                </div>
                <div class="my-blue cell" :title="users[countUser + 5].p_6">
                  <i :class="users[countUser + 5].p_6_icon + ' icon'"></i>
                  {{users[countUser + 5].p_6}}
                </div>
              </div>
              <div class="one-user__info-row">
                <div class="my-blue cell" :title="users[countUser + 5].p_7">
                  <i :class="users[countUser + 5].p_7_icon + ' icon'"></i>
                  {{users[countUser + 5].p_7}}

                </div>
                <div class="my-blue cell" :title="users[countUser + 5].p_8">
                  <i :class="users[countUser + 5].p_8_icon + ' icon'"></i>
                  {{users[countUser + 5].p_8}}
                </div>
                <div class="my-blue cell" :title="users[countUser + 5].p_9">
                  <i :class="users[countUser + 5].p_9_icon + ' icon'"></i>
                  {{users[countUser + 5].p_9}}
                </div>
              </div>
            </div>
          </div>
        </div>

      </div>


      <!--    <v-footer app>
            <span>&copy; {{ new Date().getFullYear() }}</span>
          </v-footer>-->
    </div>
  </v-app>
</template>

<script>
  import UserElem from "../class/UserElem";

  export default {
    props: {
      source: String,
    },
    data: () => ({
      timerKek: null,
      flag_keking: false,
      drawer: false,
      users: [],
    }),
    created() {
      // this.$vuetify.theme.dark = true
      for (let i = 0; i < 10; i++) {
        this.users.push(new UserElem());

      }
    },
    methods: {
      getRandomInt(max) {
        return Math.floor(Math.random() * Math.floor(max));
      },
      startKek() {
        this.flag_keking = true;
        this.timerKek = setInterval(() => {
          let user = this.users[this.getRandomInt(10)];
          if(user.img === UserElem.shrekStandart){
            user.img = UserElem.shrekKek;
          }else{
            user.img = UserElem.shrekStandart;
          }
        }, 300);
      },
      stopKek(){
        this.flag_keking = false;
        clearInterval(this.timerKek);
        this.users.forEach(value => {
          value.img = UserElem.shrekStandart;
        })
      }
    },
  }
</script>
