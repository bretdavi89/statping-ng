<template>
  <div>
    <div class="card contain-card mb-5">
      <div class="card-header">
        {{ message.id ? `${$t('update')} ${message.title}` : $t('message_create') }}
        <transition name="slide-fade">
          <button
            v-if="message.id"
            class="btn btn-sm float-right btn-danger btn-sm"
            @click="removeEdit"
          >
            {{ $t('close') }}
          </button>
        </transition>
      </div>
      <div class="card-body">
        <form @submit="saveMessage">
          <div class="form-group row">
            <label class="col-sm-4 col-form-label">{{ $t('title') }}</label>
            <div class="col-sm-8">
              <input
                id="title"
                v-model="message.title"
                type="text"
                name="title"
                class="form-control"
                placeholder="Announcement Title"
                required
              >
            </div>
          </div>

          <div class="form-group row">
            <label class="col-sm-4 col-form-label">{{ $t('description') }}</label>
            <div class="col-sm-8">
              <textarea
                id="description"
                v-model="message.description"
                rows="5"
                name="description"
                class="form-control"
                required
              />
              <font-awesome-icon
                class="mdicon"
                icon="fab fa-markdown"
              />
            </div>
          </div>

          <div class="form-group row">
            <label class="col-sm-4 col-form-label">{{ $t('service') }}</label>
            <div class="col-sm-8">
              <select
                v-model.number="message.service"
                name="service_id"
                class="form-control"
              >
                <option :value="0">
                  {{ $t('global_announcement') }}
                </option>
                <option
                  v-for="service in $store.getters.services"
                  :key="service.id"
                  :value="service.id"
                >
                  {{ service.name }}
                </option>
              </select>
            </div>
          </div>

          <div class="form-group row">
            <label class="col-sm-4 col-form-label">{{ $t('announcement_date') }}</label>
            <div class="col-sm-4">
              <flatPickr
                id="start_on"
                v-model="message.start_on"
                :config="config"
                type="text"
                name="start_on"
                class="form-control form-control-plaintext"
                required
                value="0001-01-01T00:00:00Z"
                @on-change="startChange"
              />
            </div>
            <div class="col-sm-4 mt-3 mt-md-0">
              <flatPickr
                id="end_on"
                v-model="message.end_on"
                :config="config"
                type="text"
                name="end_on"
                class="form-control form-control-plaintext"
                value="0001-01-01T00:00:00Z"
                required
                @on-change="endChange"
              />
            </div>
          </div>

          <div class="form-group row">
            <label
              for="notify_method"
              class="col-sm-4 col-form-label"
            >{{ $t('notify_users') }}</label>
            <div class="col-sm-8">
              <span
                class="switch"
                @click="message.notify = !!message.notify"
              >
                <input
                  id="switch-normal"
                  v-model="message.notify"
                  type="checkbox"
                  class="switch"
                >
                <label for="switch-normal">{{ $t('notify_desc') }}</label>
              </span>
            </div>
          </div>

          <div
            v-if="message.service !== 0"
            class="form-group row"
          >
            <label
              for="notify_method"
              class="col-sm-4 col-form-label"
            >{{ $t('maintenance_mode') }}</label>
            <div class="col-sm-8">
              <span
                class="switch"
                @click="message.maintenance_mode = !!message.maintenance_mode"
              >
                <input
                  id="switch-maint"
                  v-model="message.maintenance_mode"
                  type="checkbox"
                  class="switch"
                >
                <label for="switch-maint">{{ $t('maintenance_desc') }}</label>
              </span>
            </div>
          </div>

          <div
            v-if="message.notify"
            class="form-group row"
          >
            <label
              for="notify_method"
              class="col-sm-4 col-form-label"
            >{{ $t('notify_method') }}</label>
            <div class="col-sm-8">
              <input
                id="notify_method"
                v-model="message.notify_method"
                type="text"
                name="notify_method"
                class="form-control"
                value=""
                placeholder="email"
              >
            </div>
          </div>

          <div
            v-if="message.notify"
            class="form-group row"
          >
            <label
              for="notify_before"
              class="col-sm-4 col-form-label"
            >{{ $t('notify_before') }}</label>
            <div class="col-sm-8">
              <div class="form-inline">
                <input
                  id="notify_before"
                  v-model.number="message.notify_before"
                  type="number"
                  name="notify_before"
                  class="col-4 form-control"
                >
                <select
                  id="notify_before_scale"
                  v-model="message.notify_before_scale"
                  class="ml-2 col-7 form-control"
                  name="notify_before_scale"
                >
                  <option value="minute">
                    {{ $t('minutes') }}
                  </option>
                  <option value="hour">
                    {{ $t('hours') }}
                  </option>
                  <option value="day">
                    {{ $t('days') }}
                  </option>
                </select>
              </div>
            </div>
          </div>

          <div class="form-group row">
            <div class="col-sm-12">
              <button
                :disabled="!message.title || !message.description"
                type="submit"
                class="btn btn-block"
                :class="{'btn-primary': !message.id, 'btn-secondary': message.id}"
                @click="saveMessage"
              >
                {{ message.id ? $t('message_edit') : $t('message_create') }}
              </button>
            </div>
          </div>
          <div
            id="alerter"
            class="alert alert-danger d-none"
            role="alert"
          />
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import Api from '../API';
import flatPickr from 'vue-flatpickr-component';
import 'flatpickr/dist/flatpickr.css';

export default {
    name: 'FormMessage',
    components: {
        flatPickr
    },
    props: {
        in_message: {
            type: Object
        },
        service: {
            type: Object
        },
        edit: {
            type: Function
        }
    },
    data () {
        return {
            message: {
                title: '',
                description: '',
                start_on: new Date(),
                end_on: new Date(),
                service_id: 0,
                service: 0,
                notify_method: '',
                notify: false,
                notify_before: 0,
                notify_before_scale: 'minute',
            },
            config: {
                altFormat: 'l M J, \\at h:iK',
                altInput: true,
                enableTime: true,
                dateFormat: 'Z',
            },
            temp: {}
        };
    },
    watch: {
        in_message () {
            this.message = this.in_message;
        }
    },
    mounted () {
        if (this.service) {
            this.service_id = this.service.id;
        }
    },
    methods: {
        startChange (e) {
            window.console.log(e);
        },
        endChange (e) {
            window.console.log(e);
        },
        removeEdit () {
            this.message = {};
            this.edit(false);
        },
        async saveMessage (e) {
            e.preventDefault();
            if (this.message.id) {
                await this.updateMessage();
            } else {
                await this.createMessage();
            }
        },
        async createMessage () {
            await Api.message_create(this.message);
            const messages = await Api.messages();
            this.$store.commit('setMessages', messages);
            this.message = {};
        },
        async updateMessage () {
            await Api.message_update(this.message);
            const messages = await Api.messages();
            this.$store.commit('setMessages', messages);
            this.edit(false);
        }
    }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
